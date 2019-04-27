package main

import "C"

import (
	"fmt"
	"gitlab.wow.st/gmp/nswrap/examples/foundation/ns"
)

func main() {
	fmt.Println("Started")
	n := ns.StringWithUTF8String(ns.CharFromString("hi there"))
	c := n.CapitalizedString()
	gstring := c.UTF8String().String()
	fmt.Println(gstring)
	fmt.Println("ok")
}

