package main
//go:generate nswrap

import (
	"fmt"

	"git.wow.st/gmp/nswrap/examples/foundation/ns"
)

func main() {
	fmt.Printf("Creating some strings:\n")
	n1 := ns.NSStringWithUTF8String(ns.CharWithGoString("hi there"))
	c1 := n1.CapitalizedString()
	gs := c1.UTF8String().String()
	fmt.Println(gs)
	n2 := ns.NSStringAlloc()
	n2 = n2.InitWithGoString("hi world")
	n3 := ns.NSStringWithGoString("ok bye")
	fmt.Printf("%s\n",n3)

	fmt.Printf("\nCreating an array:\n")
	a := ns.NSMutableArrayWithObjects(n1,n2,n3)
	fmt.Println("Length(a) = ",a.Count())
	fmt.Println("is n2 in a?",a.ContainsObject(n2))
	fmt.Println("is c1 in a?",a.ContainsObject(c1))
	fmt.Printf("\nCreating substrings and adding to the array:\n")
	n4 := n2.SubstringFromIndex(3)
	n5 := n3.SubstringToIndex(4)
	a.AddObject(n4)
	a.AddObject(n5)
	fmt.Println("Length(a) = ",a.Count())
	fmt.Printf("\nSubarray with range:\n")
	a2 := a.SubarrayWithRange(ns.NSMakeRange(1,3))
	fmt.Println("Length(a2) = ",a2.Count())
	fmt.Printf("\nObjectAtIndex(1):\n")
	i1 := a.ObjectAtIndex(1).NSString()
	fmt.Printf("i1 = %@\n",i1)
	fmt.Printf("i1.Ptr() = %p\n",i1.Ptr())
	fmt.Printf("\nNSArray.ObjectEnumerator().ForIn():\n")
	a.ObjectEnumerator().ForIn(func(o ns.Id) bool {
		fmt.Println(o.NSString())
		return true
	})
	fmt.Printf("\nNSSetWithObjectsCount():\n")
	s1 := ns.NSSetWithObjectsCount(&[]ns.Id{n1.Id,n2.Id},2)
	fmt.Printf("\nNSSet.ObjectEnumerator().ForIn():\n")
	s1.ObjectEnumerator().ForIn(func(o ns.Id) bool {
		fmt.Println(o.NSString())
		return true
	})
	fmt.Printf("\nNSMutableArrayWithObjects()\n")
	a = ns.NSMutableArrayWithObjects(n1,s1)
	fmt.Printf("\nNSArray.ObjectEnumerator().ForIn():\n")
	a.ObjectEnumerator().ForIn(func(o ns.Id) bool {
		fmt.Printf("%s -- ",o.ClassName().UTF8String())
		switch {
		case o.IsKindOfClass(ns.NSStringClass()):
			fmt.Printf("  It's a string\n")
		case o.IsKindOfClass(ns.NSSetClass()):
			fmt.Printf("  It's a set\n")
		default:
			fmt.Printf("  I don't know what it is!\n")
		}
		return true
	})
	fmt.Printf("\nNSArrayWithObjects()\n")
	a2 = ns.NSArrayWithObjects(n1,n2,n3,s1)
	fmt.Printf("\nNSArray.ObjectEnumerator().ForIn():\n")
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

	nst := ns.NSStringWithGoString
	fmt.Printf("\nNSDictionaryWithObjectsForKeys()\n")
	d := ns.NSDictionaryWithObjectsForKeys(
		ns.NSArrayWithObjects(nst("obj1"),nst("obj2")),
		ns.NSArrayWithObjects(nst("key1"),nst("key2")),
	)
	os := make([]ns.Id,0,5)
	fmt.Printf("Length of os is %d\n",len(os))
	ks := make([]ns.Id,0,5)
	fmt.Printf("\nGetObjects()\n")
	d.GetObjects(&os,&ks,4)
	fmt.Printf("Length of os is now %d\n",len(os))
	for i,k := range ks {
		fmt.Printf("-- %s -> %s\n",k.NSString(),os[i].NSString())
	}
	fmt.Printf("\nNSStringWithContentsOfURLEncoding()\n")
	err := make([]ns.NSError,1)
	n1 = ns.NSStringWithContentsOfURLEncoding(ns.NSURLWithGoString("htttypo://example.com"),0,&err)
	fmt.Printf("err: %s\n",err[0].LocalizedDescription())

	fmt.Printf("\nNSStringWithFormat()\n")
	str := ns.NSStringWithFormat(nst("(%@) (%@)\n(%@)\n"),n2,n3,s1)
	fmt.Printf("%s\n",str)
}
