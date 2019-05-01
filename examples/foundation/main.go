package main
//go:generate nswrap

import (
	"fmt"

	"gitlab.wow.st/gmp/nswrap/examples/foundation/ns"
)

func main() {
	n1 := ns.NSStringStringWithUTF8String(ns.CharFromString("hi there"))
	c1 := n1.CapitalizedString()
	gs := c1.UTF8String().String()
	fmt.Println(gs)
	n2 := ns.NSStringStringWithUTF8String(ns.CharFromString("hi world"))
	n3 := ns.NSStringStringWithUTF8String(ns.CharFromString("ok bye"))
	a := ns.NSMutableArrayArrayWithObjects(n1,n2,n3)
	fmt.Println("Length(a) = ",a.Count())
	fmt.Println("is n2 in a?",a.ContainsObject(n2))
	fmt.Println("is c1 in a?",a.ContainsObject(c1))
	n4 := n2.SubstringFromIndex(3)
	n5 := n3.SubstringToIndex(4)
	a.AddObject(n4)
	a.AddObject(n5)
	fmt.Println("Length(a) = ",a.Count())
}
