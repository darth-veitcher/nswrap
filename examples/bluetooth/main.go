package main

import "C"

import (
	"encoding/binary"
	"fmt"
	"git.wow.st/gmp/nswrap/examples/bluetooth/ns"
)

func updateState(c ns.CBCentralManager) {
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
		cm.ScanForPeripheralsWithServices(ns.NSArrayWithObjects(hrm_uuid),ns.NSDictionary{})
	}
}

func discoverPeripheral(c ns.CBCentralManager, p ns.CBPeripheral, d ns.NSDictionary, rssi ns.NSNumber) {
	fmt.Printf("Did discover peripheral\n")
	c.StopScan()
	if peripheral.Ptr() != nil {
		peripheral.Release()
	}
	peripheral = p
	peripheral.Retain()
	c.ConnectPeripheral(peripheral,ns.NSDictionary{})
}

func connectPeripheral(c ns.CBCentralManager, p ns.CBPeripheral) {
	fmt.Printf("Did connect peripheral\n")
	p.SetDelegate(cd)
	p.DiscoverServices(ns.NSArray{})
}

func discoverServices(p ns.CBPeripheral, e ns.NSError) {
	fmt.Printf("Did discover services\n")
	p.Services().ObjectEnumerator().ForIn(func(o ns.Id) bool {
		serv := o.CBService()
		switch {
		case serv.UUID().IsEqualTo(hrm_uuid):
			fmt.Printf("--heart rate monitor service\n")
			p.DiscoverCharacteristics(ns.NSArray{},serv)
		case serv.UUID().IsEqualTo(ns.CBUUIDWithGoString("180A")):
			fmt.Printf("--device information service\n")
			p.DiscoverCharacteristics(ns.NSArray{},serv)
		}
		return true
	})
}

func hr(d ns.NSData) int {
	if l := int(d.Length()); l < 4 {
		return 0
	}
	x := C.GoBytes(d.Bytes(),4)
	flags := x[0]
	if flags & 0x80 != 0 { // uint16 format
		return int(binary.BigEndian.Uint16(x[1:2]))
	} else {
		return int(x[1])
	}
}

func discoverCharacteristics(p ns.CBPeripheral, s ns.CBService, e ns.NSError) {
	fmt.Printf("Did discover characteristics\n")
	fmt.Printf("----%s\n",s.UUID().UUIDString().UTF8String())
	if s.UUID().IsEqualTo(hrm_uuid) {
		s.Characteristics().ObjectEnumerator().ForIn(func(o ns.Id) bool {
			chr := o.CBCharacteristic()
			fmt.Printf("------%s\n",chr.UUID().UUIDString().UTF8String())
			if chr.UUID().IsEqualTo(hrv_uuid) {
				p.SetNotifyValue(1,chr)
				v := chr.Value()
				fmt.Println(hr(v))
			}
			return true
		})
	}
}

func updateValue(p ns.CBPeripheral, chr ns.CBCharacteristic, e ns.NSError) {
	if chr.UUID().IsEqualTo(hrv_uuid) {
		v := chr.Value()
		fmt.Printf("Heart rate: %d\n",hr(v))
	}
}

var (
	hrm_uuid ns.CBUUID
	hrv_uuid ns.CBUUID
	cd ns.CBDelegate
	cm ns.CBCentralManager
	peripheral ns.CBPeripheral
)

func main() {
	queue := ns.DispatchQueueCreate(ns.CharWithGoString("my_new_queue"),nil)

	cd = ns.CBDelegateAlloc()

	cd.CentralManagerDidUpdateStateCallback(updateState)
	cd.CentralManagerDidDiscoverPeripheralCallback(discoverPeripheral)
	cd.CentralManagerDidConnectPeripheralCallback(connectPeripheral)
	cd.PeripheralDidDiscoverServicesCallback(discoverServices)
	cd.PeripheralDidDiscoverCharacteristicsForServiceCallback(discoverCharacteristics)
	cd.PeripheralDidUpdateValueForCharacteristicCallback(updateValue)

	hrm_uuid = ns.CBUUIDWithGoString("180D")
	hrv_uuid = ns.CBUUIDWithGoString("2A37")

	cm = ns.CBCentralManagerAlloc().InitWithDelegateQueue(cd,queue)

	select { }
}
