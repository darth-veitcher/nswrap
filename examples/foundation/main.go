package main
//go:generate nswrap

import (
	"fmt"

	"gitlab.wow.st/gmp/nswrap/examples/foundation/ns"
)

func main() {
	n1 := ns.StringWithUTF8String(ns.CharFromString("hi there"))
	c1 := n1.CapitalizedString()
	gs := c1.UTF8String().String()
	fmt.Println(gs)
	n2 := ns.StringWithUTF8String(ns.CharFromString("hi world"))
	n3 := ns.StringWithUTF8String(ns.CharFromString("ok bye"))
	a := ns.ArrayWithObjects(n1,n2,n3)
	fmt.Println("Length(a) = ",a.Count())
}
