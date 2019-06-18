package main

import "C"

import (
	"encoding/binary"
	"fmt"
	"runtime"
	"time"
	"git.wow.st/gmp/nswrap/examples/bluetooth/ns"
)

func updateState(c *ns.CBCentralManager) {
	fmt.Printf("Go: did update state\n")
	switch cm.CBManager.State() {
	case ns.CBManagerStateUnknown:
		fmt.Printf("  unknown\n")
	case ns.CBManagerStateResetting:
		fmt.Printf("  resetting\n")
	case ns.CBManagerStateUnsupported:
		fmt.Printf("  unsupported\n")
	case ns.CBManagerStateUnauthorized:
		fmt.Printf("  unauthorized\n")
	case ns.CBManagerStatePoweredOff:
		fmt.Printf("  powered off\n")
	case ns.CBManagerStatePoweredOn:
		fmt.Printf("  powered on\n")
		cm.ScanForPeripheralsWithServices(ns.NSArrayWithObjects(hrm_uuid), nil)
	}
	fmt.Printf("Go: updateState returning\n")
}

func discoverPeripheral(c *ns.CBCentralManager, p *ns.CBPeripheral, d *ns.NSDictionary, rssi *ns.NSNumber) {
	fmt.Printf("Did discover peripheral\n")
	c.StopScan()

	// if we already have a pointer to a peripheral, check that this one is different,
	// and if so release the old one. Be careful to check the Objective-C pointers
	// here as the Go pointers will differ.

	if peripheral != nil && p.Ptr() != peripheral.Ptr() {
		peripheral.Release()
	}

	peripheral = p

	// we need to take ownership of this peripheral so CoreBluetooth doesn't
	// dealloc it.

	peripheral.Retain()
	c.ConnectPeripheral(peripheral, nil)
}

func connectPeripheral(c *ns.CBCentralManager, p *ns.CBPeripheral) {
	fmt.Printf("Did connect peripheral\n")

	// set ourselves up as a peripheral delegate

	p.SetDelegate(cd)

	// discover all services on this device

	p.DiscoverServices(nil)
	fmt.Printf("Go: discoverPeripheral returning\n")
}

func discoverServices(p *ns.CBPeripheral, e *ns.NSError) {
	fmt.Printf("Did discover services\n")
	p.Services().ObjectEnumerator().ForIn(func(o *ns.Id) bool {
		serv := o.CBService()
		uuid := serv.UUID()
		switch {
		case uuid.IsEqualTo(hrm_uuid):
			fmt.Printf("--heart rate monitor service\n")
			p.DiscoverCharacteristics(nil, serv)
		case uuid.IsEqualTo(info_uuid):
			fmt.Printf("--device information service\n")
			p.DiscoverCharacteristics(nil, serv)
		default:
			fmt.Printf("--unknown service\n")
		}
		return true
	})
	fmt.Printf("Go: discoverServices returning\n")
}

func hr(d *ns.NSData) int {
	if l := int(d.Length()); l < 4 {
		return 0
	}
	x := C.GoBytes(d.Bytes(), 4)
	flags := x[0]
	if flags&0x80 != 0 { // uint16 format
		return int(binary.BigEndian.Uint16(x[1:2]))
	} else {
		return int(x[1])
	}
}

func discoverCharacteristics(p *ns.CBPeripheral, s *ns.CBService, e *ns.NSError) {
	fmt.Printf("Did discover characteristics\n")
	uuid := s.UUID()
	fmt.Printf("----%s\n", uuid.UUIDString().UTF8String())
	if uuid.IsEqualTo(hrm_uuid) {
		s.Characteristics().ObjectEnumerator().ForIn(func(o *ns.Id) bool {
			chr := o.CBCharacteristic()
			chuuid := chr.UUID()
			fmt.Printf("------%s\n", chuuid.UUIDString().UTF8String())
			if chuuid.IsEqualTo(hrv_uuid) {
				p.SetNotifyValue(1, chr)
				v := chr.Value()
				fmt.Println(hr(v))
			}
			return true
		})
	}
	fmt.Printf("Go: discoverCharacteristics returning\n")
}

func updateValue(p *ns.CBPeripheral, chr *ns.CBCharacteristic, e *ns.NSError) {
	if chr.UUID().IsEqualTo(hrv_uuid) {
		v := chr.Value()
		fmt.Printf("Heart rate: %d\n", hr(v))
	}
	fmt.Printf("Go: updateValue returning\n")
}

var (
	hrm_uuid   *ns.CBUUID
	hrv_uuid   *ns.CBUUID
	info_uuid  *ns.CBUUID
	cd         *ns.CBDelegate
	cm         *ns.CBCentralManager
	peripheral *ns.CBPeripheral
)

func main() {
	queue := ns.DispatchQueueCreate(ns.CharWithGoString("my_new_queue"), nil)

	cd = ns.CBDelegateAlloc()

	cd.CentralManagerDidUpdateStateCallback(updateState)
	cd.CentralManagerDidDiscoverPeripheralCallback(discoverPeripheral)
	cd.CentralManagerDidConnectPeripheralCallback(connectPeripheral)
	cd.PeripheralDidDiscoverServicesCallback(discoverServices)
	cd.PeripheralDidDiscoverCharacteristicsForServiceCallback(discoverCharacteristics)
	cd.PeripheralDidUpdateValueForCharacteristicCallback(updateValue)

	hrm_uuid = ns.CBUUIDWithGoString("180D")
	hrv_uuid = ns.CBUUIDWithGoString("2A37")
	info_uuid = ns.CBUUIDWithGoString("180A")

	// We defined our own queue because this won't work on the main queue.
	cm = ns.CBCentralManagerAlloc().InitWithDelegateQueue(cd, queue)

	// For debugging purposes, run GC every second to make sure things are not
	// over-released.
	go func() {
		for {
			runtime.GC()
			time.Sleep(time.Second)
		}
	}()
	select {}
}
