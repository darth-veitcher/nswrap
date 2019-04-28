package main

import "C"

import (
	"fmt"
	"gitlab.wow.st/gmp/nswrap/examples/foundation/ns"
)

func main() {
	fmt.Println("Started")
	n1 := ns.StringWithUTF8String(ns.CharFromString("hi there"))
	c := n1.CapitalizedString()
	gstring := c.UTF8String().String()
	fmt.Println(gstring)
	n2 := ns.StringWithUTF8String(ns.CharFromString("ok bye"))
	n3 := ns.StringWithUTF8String(ns.CharFromString("one two three"))
	a := ns.ArrayWithObjects(n1)
	fmt.Println("ok")
}

