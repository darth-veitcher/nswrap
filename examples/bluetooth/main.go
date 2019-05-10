package main

import "C"

import (
	"fmt"
	"gitlab.wow.st/gmp/nswrap/examples/bluetooth/ns"
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
	}
}

var (
	cm *ns.CBCentralManager
)

func main() {
	queue := ns.DispatchQueueCreate(ns.CharWithGoString("st.wow.gitlab.ble"),nil)

	cd := ns.BleDelegateAlloc()
	cd.CentralManagerDidUpdateStateCallback(updateState)

	cm = ns.CBCentralManagerAlloc()
	cm.InitWithDelegate(cd,queue,nil)

	uuid := ns.CBUUIDWithGoString("180d")
	_ = uuid
	select { }
}
