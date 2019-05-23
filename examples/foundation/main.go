package main
//go:generate nswrap

import (
	"fmt"

	"gitlab.wow.st/gmp/nswrap/examples/foundation/ns"
)

func main() {
	n1 := ns.NSStringWithUTF8String(ns.CharWithGoString("hi there"))
	c1 := n1.CapitalizedString()
	gs := c1.UTF8String().String()
	fmt.Println(gs)
	n2 := ns.NSStringAlloc()
	n2 = n2.InitWithGoString("hi world")
	n3 := ns.NSStringWithGoString("ok bye")
	a := ns.NSMutableArrayWithObjects(n1,n2,n3)
	fmt.Println("Length(a) = ",a.Count())
	fmt.Println("is n2 in a?",a.ContainsObject(n2))
	fmt.Println("is c1 in a?",a.ContainsObject(c1))
	n4 := n2.SubstringFromIndex(3)
	n5 := n3.SubstringToIndex(4)
	a.AddObject(n4)
	a.AddObject(n5)
	fmt.Println("Length(a) = ",a.Count())
	a2 := a.SubarrayWithRange(ns.NSMakeRange(1,3))
	fmt.Println("Length(a2) = ",a2.Count())
	i1 := a.ObjectAtIndex(1).NSString()
	fmt.Printf("i1 = %@\n",i1)
	fmt.Printf("i1.Ptr() = %p\n",i1.Ptr())
	a.ObjectEnumerator().ForIn(func(o ns.Id) bool {
		fmt.Println(o.NSString().UTF8String())
		return true
	})
	s1 := ns.NSSetWithObjects(n1,n2)
	a = ns.NSMutableArrayWithObjects(n1,s1)
	a.ObjectEnumerator().ForIn(func(o ns.Id) bool {
		fmt.Printf("%s -- ",o.ClassName().UTF8String())
		switch {
		case o.IsKindOfClass(ns.NSStringClass()):
			fmt.Printf("It's a string\n")
		case o.IsKindOfClass(ns.NSSetClass()):
			fmt.Printf("It's a set\n")
		default:
			fmt.Printf("I don't know what it is!\n")
		}
		return true
	})
	a2 = ns.NSArrayWithObjects(n1,n2,n3,s1)
	a2.ObjectEnumerator().ForIn(func (o ns.Id) bool {
		switch {
		case o.IsKindOfClass(ns.NSStringClass()):
			fmt.Println(o.NSString().UTF8String())
			return true  // continue enumeration
		default:
			fmt.Println("Unknown class")
			return false  // terminate enumeration
		}
	})
}
