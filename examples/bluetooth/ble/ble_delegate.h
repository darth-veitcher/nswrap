#import <CoreBluetooth/CoreBluetooth.h>

@interface ble_delegate : NSObject <CBCentralManagerDelegate, CBPeripheralDelegate> 
{
    CBCentralManager *manager;
    CBPeripheral *peripheral;

    BOOL wantScan;
    BOOL autoConnect;
    void (*scanCallback)(void* p);
    dispatch_queue_t q;
    CBUUID *looking_for;
}


- (ble_delegate*) init;
- (void) scanFor:(CBUUID*)uuid;
- (void) stopScan;
- (BOOL) isLECapableHardware;


@end

