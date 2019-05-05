package main

import (
	"fmt"
	"time"
	"gitlab.wow.st/gmp/nswrap/examples/bluetooth/ble"
)

func main() {
	cd := ble.Ble_delegateAlloc().Init()
	fmt.Println("LE capable:",cd.IsLECapableHardware())
	time.Sleep(time.Second * 1)
	fmt.Println("LE capable:",cd.IsLECapableHardware())
	uuid := ble.CBUUIDWithGoString("180d")
	cd.ScanFor(uuid)

	select { }
}
