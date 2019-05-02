#import "ble_delegate.h"

@implementation ble_delegate

- (ble_delegate*) init
{
    [super init];
    NSLog(@"Initializing manager");
    q = dispatch_queue_create("st.wow.gitlab.ble",NULL);
    [q retain];
    manager = [[CBCentralManager alloc] initWithDelegate:self queue:q];
    return self;
}

- (void) dealloc
{
    [self stopScan];

    [peripheral setDelegate:nil];
    [peripheral release];

    [manager release];

    [super dealloc];
}

#pragma mark - Start/Stop Scan methods

/*
 Request CBCentralManager to scan for heart rate peripherals using service UUID 0x180D
 */
- (void) scanFor:(CBUUID*)uuid
{
    autoConnect = TRUE;
    [self _scanFor:uuid];
}

- (void) _scanFor:(CBUUID*)uuid 
{
    if (uuid != nil) {
	if (looking_for != nil) {
		[looking_for release];
	}
        looking_for = uuid;
    }
    if (looking_for == nil) {
        NSLog(@"No scan target specified");
        return;
    }
    if (![self isLECapableHardware]) {
        NSLog(@"Not LE capable hardware");
        NSLog(@"Setting wantScan to true");
        wantScan = true;
        return;
    } else {
        NSLog(@"Is LE capable hardware");
    }
    NSLog(@"Scanning");
    [manager scanForPeripheralsWithServices:[NSArray arrayWithObject:looking_for] options:nil];
}

/*
 Request CBCentralManager to stop scanning for heart rate peripherals
 */
- (void) stopScan 
{
    [manager stopScan];
    scanCallback(nil);
}

#pragma mark - CBCentralManager delegate methods
/*
 Invoked whenever the central manager's state is updated.
 */
- (void) centralManagerDidUpdateState:(CBCentralManager *)central 
{
    NSLog(@"State changed");
    if ([self isLECapableHardware] && wantScan) {
        NSLog(@"Starting scan?");
        [self _scanFor:nil];
    }
}

- (BOOL) isLECapableHardware
{
    NSString * state = nil;

    switch ([manager state])
    {
        case CBManagerStateUnsupported:
            NSLog(@"--unsupported");
            state = @"The platform/hardware doesn't support Bluetooth Low Energy.";
            break;
        case CBManagerStateUnauthorized:
            NSLog(@"--unauthorized");
            state = @"The app is not authorized to use Bluetooth Low Energy.";
            break;
        case CBManagerStatePoweredOff:
            NSLog(@"--powered off");
            state = @"Bluetooth is currently powered off.";
            break;
        case CBManagerStatePoweredOn:
            NSLog(@"--powered on");
            return TRUE;
        case CBManagerStateResetting:
            NSLog(@"--resetting");
            return FALSE;
        case CBManagerStateUnknown:
            NSLog(@"--unknown");
            state = @"Bluetooth state is unknown.";
    }

    NSLog(@"Central manager state: %@", state);
    return FALSE;
}
    
/*
 Invoked when the central discovers a peripheral while scanning.
 */
- (void) centralManager:(CBCentralManager *)central didDiscoverPeripheral:(CBPeripheral *)aPeripheral advertisementData:(NSDictionary *)advertisementData RSSI:(NSNumber *)RSSI 
{
    NSLog(@"Found peripheral");
    peripheral = aPeripheral;
    [peripheral retain];
    if (autoConnect) {
        [self stopScan];
    	printf("Connecting\n");
    	[manager connectPeripheral:aPeripheral options:[NSDictionary dictionaryWithObject:[NSNumber numberWithBool:YES] forKey:CBConnectPeripheralOptionNotifyOnDisconnectionKey]];
    } else {
        scanCallback(aPeripheral);
    }
}

/*
 Invoked when the central manager retrieves the list of known peripherals.
 Automatically connect to first known peripheral
 */
- (void)centralManager:(CBCentralManager *)central didRetrievePeripherals:(NSArray *)peripherals
{
    NSLog(@"Retrieved peripheral: %lu - %@", [peripherals count], peripherals);
    
    [self stopScan];
    
    /* If there are any known devices, automatically connect to it.*/
    if([peripherals count] >=1)
    {
        peripheral = [peripherals objectAtIndex:0];
        [peripheral retain];
        printf("Connecting\n");
        [manager connectPeripheral:peripheral options:[NSDictionary dictionaryWithObject:[NSNumber numberWithBool:YES] forKey:CBConnectPeripheralOptionNotifyOnDisconnectionKey]];
    }
}

/*
 Invoked whenever a connection is succesfully created with the peripheral. 
 Discover available services on the peripheral
 */
- (void) centralManager:(CBCentralManager *)central didConnectPeripheral:(CBPeripheral *)aPeripheral 
{ 
    NSLog(@"Connected");
    [aPeripheral setDelegate:self];
    [aPeripheral discoverServices:nil];
}

/*
 Invoked whenever an existing connection with the peripheral is torn down. 
 Reset local variables
 */
- (void)centralManager:(CBCentralManager *)central didDisconnectPeripheral:(CBPeripheral *)aPeripheral error:(NSError *)error
{
    if( peripheral )
    {
        [peripheral setDelegate:nil];
        [peripheral release];
        peripheral = nil;
    }
}

/*
 Invoked whenever the central manager fails to create a connection with the peripheral.
 */
- (void)centralManager:(CBCentralManager *)central didFailToConnectPeripheral:(CBPeripheral *)aPeripheral error:(NSError *)error
{
    NSLog(@"Fail to connect to peripheral: %@ with error = %@", aPeripheral, [error localizedDescription]);
    if( peripheral )
    {
        [peripheral setDelegate:nil];
        [peripheral release];
        peripheral = nil;
    }
}

#pragma mark - CBPeripheral delegate methods
/*
 Invoked upon completion of a -[discoverServices:] request.
 Discover available characteristics on interested services
 */
- (void) peripheral:(CBPeripheral *)aPeripheral didDiscoverServices:(NSError *)error 
{
    NSLog(@"Discovered services");
    for (CBService *aService in aPeripheral.services) 
    {
        NSLog(@"Service found with UUID: %@", aService.UUID);
        
        /* Heart Rate Service */
        if ([aService.UUID isEqual:[CBUUID UUIDWithString:@"180D"]]) 
        {
            NSLog(@"--heart rate service");
            [aPeripheral discoverCharacteristics:nil forService:aService];
        }
        
        /* Device Information Service */
        if ([aService.UUID isEqual:[CBUUID UUIDWithString:@"180A"]]) 
        {
            NSLog(@"--device information service");
            [aPeripheral discoverCharacteristics:nil forService:aService];
        }
        
        /* GAP (Generic Access Profile) for Device Name */
        if ( [aService.UUID isEqual:[CBUUID UUIDWithString:@"1800"]] )
        {
            NSLog(@"--generic access profile");
            [aPeripheral discoverCharacteristics:nil forService:aService];
        }
    }
}

/*
 Invoked upon completion of a -[discoverCharacteristics:forService:] request.
 Perform appropriate operations on interested characteristics
 */
- (void) peripheral:(CBPeripheral *)aPeripheral didDiscoverCharacteristicsForService:(CBService *)service error:(NSError *)error 
{
    NSLog(@"Discovered characteristics");
    if ([service.UUID isEqual:[CBUUID UUIDWithString:@"180D"]]) 
    {
        for (CBCharacteristic *aChar in service.characteristics) 
        {
            /* Set notification on heart rate measurement */
            if ([aChar.UUID isEqual:[CBUUID UUIDWithString:@"2A37"]]) 
            {
                [peripheral setNotifyValue:YES forCharacteristic:aChar];
                NSLog(@"Found a Heart Rate Measurement Characteristic");
            }
            /* Read body sensor location */
            if ([aChar.UUID isEqual:[CBUUID UUIDWithString:@"2A38"]]) 
            {
                [aPeripheral readValueForCharacteristic:aChar];
                NSLog(@"Found a Body Sensor Location Characteristic");
            } 
            
            /* Write heart rate control point */
            if ([aChar.UUID isEqual:[CBUUID UUIDWithString:@"2A39"]])
            {
                uint8_t val = 1;
                NSData* valData = [NSData dataWithBytes:(void*)&val length:sizeof(val)];
                [aPeripheral writeValue:valData forCharacteristic:aChar type:CBCharacteristicWriteWithResponse];
            }
        }
    }
    
    if ( [service.UUID isEqual:[CBUUID UUIDWithString:@"1800"]] )
    {
        for (CBCharacteristic *aChar in service.characteristics) 
        {
            /* Read device name */
            if ([aChar.UUID isEqual:[CBUUID UUIDWithString:@"2A00"]])
            {
                [aPeripheral readValueForCharacteristic:aChar];
                NSLog(@"Found a Device Name Characteristic");
            }
        }
    }
    
    if ([service.UUID isEqual:[CBUUID UUIDWithString:@"180A"]]) 
    {
        for (CBCharacteristic *aChar in service.characteristics) 
        {
            /* Read manufacturer name */
            if ([aChar.UUID isEqual:[CBUUID UUIDWithString:@"2A29"]]) 
            {
                [aPeripheral readValueForCharacteristic:aChar];
                NSLog(@"Found a Device Manufacturer Name Characteristic");
            }
        }
    }
}

/*
 Invoked upon completion of a -[readValueForCharacteristic:] request or on the reception of a notification/indication.
 */
- (void) peripheral:(CBPeripheral *)aPeripheral didUpdateValueForCharacteristic:(CBCharacteristic *)characteristic error:(NSError *)error 
{
    NSLog(@"didUpdateValueForCharacteristic");
    /* Updated value for heart rate measurement received */
    if ([characteristic.UUID isEqual:[CBUUID UUIDWithString:@"2A37"]]) 
    {
        if( (characteristic.value)  || !error )
        {
            /* Update UI with heart rate data */
            //[self updateWithHRMData:characteristic.value];
        }
    } 
    /* Value for body sensor location received */
    else  if ([characteristic.UUID isEqual:[CBUUID UUIDWithString:@"2A38"]]) 
    {
        NSData * updatedValue = characteristic.value;        
        uint8_t* dataPointer = (uint8_t*)[updatedValue bytes];
        if(dataPointer)
        {
            uint8_t location = dataPointer[0];
            NSString*  locationString;
            switch (location)
            {
                case 0:
                    locationString = @"Other";
                    break;
                case 1:
                    locationString = @"Chest";
                    break;
                case 2:
                    locationString = @"Wrist";
                    break;
                case 3:
                    locationString = @"Finger";
                    break;
                case 4:
                    locationString = @"Hand";
                    break;
                case 5:
                    locationString = @"Ear Lobe";
                    break;
                case 6: 
                    locationString = @"Foot";
                    break;
                default:
                    locationString = @"Reserved";
                    break;
            }
            NSLog(@"Body Sensor Location = %@ (%d)", locationString, location);
        }
    }
    /* Value for device Name received */
    else if ([characteristic.UUID isEqual:[CBUUID UUIDWithString:@"2A00"]])
    {
        NSString * deviceName = [[[NSString alloc] initWithData:characteristic.value encoding:NSUTF8StringEncoding] autorelease];
        NSLog(@"Device Name = %@", deviceName);
    } 
    /* Value for manufacturer name received */
    else if ([characteristic.UUID isEqual:[CBUUID UUIDWithString:@"2A29"]]) 
    {
        NSString * manufacturer = [[[NSString alloc] initWithData:characteristic.value encoding:NSUTF8StringEncoding] autorelease];
        NSLog(@"Manufacturer Name = %@", manufacturer);
    }
}

@end

