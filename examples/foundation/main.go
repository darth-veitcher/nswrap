package main

import (
	"fmt"

	"gitlab.wow.st/gmp/nswrap/examples/foundation/ns"
)

func main() {
	n1 := ns.StringWithUTF8String(ns.CharFromString("hi there"))
	c1 := n1.CapitalizedString()
	gs := c1.UTF8String().String()
	fmt.Println(gs)
	a := ns.ArrayWithObjects(n1)
	_ = a
}
