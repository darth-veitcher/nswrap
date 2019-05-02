package main

import (
	"fmt"
	"time"
	"gitlab.wow.st/gmp/nswrap/examples/bluetooth/ble"
)

func main() {
	cd := ble.NewBle_delegate().Init()
	fmt.Println("LE capable:",cd.IsLECapableHardware())
	time.Sleep(time.Second * 1)
	fmt.Println("LE capable:",cd.IsLECapableHardware())
	uuid := ble.CBUUIDWithString(ble.NSStringWithUTF8String(ble.CharFromString("180d")))
	cd.ScanFor(uuid)
	time.Sleep(time.Second * 15)
}
