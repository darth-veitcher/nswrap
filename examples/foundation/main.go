package main

//go:generate nswrap

import (
	"fmt"
	"os"

	"git.wow.st/gmp/nswrap/examples/foundation/ns"
)

var (
	gs1 *ns.NSString
	gs2 *ns.NSString
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
	fmt.Printf("%s\n", n3)

	fmt.Printf("\nCreating an array:\n")
	a := ns.NSMutableArrayWithObjects(n1, n2, n3)
	fmt.Println("Length(a) = ", a.Count())
	fmt.Println("is n2 in a?", a.ContainsObject(n2))
	fmt.Println("is c1 in a?", a.ContainsObject(c1))
	fmt.Printf("\nCreating substrings and adding to the array:\n")
	n4 := n2.SubstringFromIndex(3)
	n5 := n3.SubstringToIndex(4)
	a.AddObject(n4)
	a.AddObject(n5)
	fmt.Println("Length(a) = ", a.Count())
	fmt.Printf("\nSubarray with range:\n")
	a2 := a.SubarrayWithRange(ns.NSMakeRange(1, 3))
	fmt.Println("Length(a2) = ", a2.Count())
	fmt.Printf("\nObjectAtIndex(1):\n")
	i1 := a.ObjectAtIndex(1).NSString()
	fmt.Printf("i1 = %@\n", i1)
	fmt.Printf("i1.Ptr() = %p\n", i1.Ptr())
	fmt.Printf("\nNSArray.ObjectEnumerator().ForIn():\n")
	x := 0
	a.ObjectEnumerator().ForIn(func(o *ns.Id) bool {
		fmt.Printf("%d: %s\n", x, o.NSString())
		x++
		return true
	})
	fmt.Printf("\nNSSetWithObjectsCount():\n")
	s1 := ns.NSSetWithObjectsCount(&[]*ns.Id{&n1.Id, &n2.Id}, 2)
	fmt.Printf("\nNSSet.ObjectEnumerator().ForIn():\n")
	s1.ObjectEnumerator().ForIn(func(o *ns.Id) bool {
		fmt.Println(o.NSString())
		return true
	})
	fmt.Printf("\nNSMutableArrayWithObjects()\n")
	a = ns.NSMutableArrayWithObjects(n1, s1)
	fmt.Printf("\nNSArray.ObjectEnumerator().ForIn():\n")
	a.ObjectEnumerator().ForIn(func(o *ns.Id) bool {
		fmt.Printf("%s -- ", o.ClassName().UTF8String())
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
	fmt.Printf("a = %p a.NSArray = %p\n", a, &a.NSArray)
	fmt.Printf("\nNSArrayWithObjects() (length 1)\n")
	a2 = ns.NSArrayWithObjects(n1)
	a2.ObjectEnumerator().ForIn(func(o *ns.Id) bool {
		switch {
		case o.IsKindOfClass(ns.NSStringClass()):
			fmt.Println(o.NSString().UTF8String())
			return true // continue enumeration
		default:
			fmt.Println("Unknown class")
			return false // terminate enumeration
		}
	})

	fmt.Printf("\nNSArrayWithObjects()\n")
	a2 = ns.NSArrayWithObjects(n1, n2, n3, s1)
	fmt.Printf("\nNSArray.ObjectEnumerator().ForIn():\n")
	a2.ObjectEnumerator().ForIn(func(o *ns.Id) bool {
		switch {
		case o.IsKindOfClass(ns.NSStringClass()):
			fmt.Println(o.NSString().UTF8String())
			return true // continue enumeration
		default:
			fmt.Println("Unknown class")
			return false // terminate enumeration
		}
	})

	nst := ns.NSStringWithGoString
	fmt.Printf("\nNSDictionaryWithObjectsForKeys()\n")
	d := ns.NSDictionaryWithObjectsForKeys(
		ns.NSArrayWithObjects(nst("obj1"), nst("obj2")),
		ns.NSArrayWithObjects(nst("key1"), nst("key2")),
	)
	oarr := make([]*ns.Id, 0, 5)
	fmt.Printf("Length of oarr is %d\n", len(oarr))
	karr := make([]*ns.Id, 0, 5)
	fmt.Printf("\nGetObjectsAndKeysCount()\n")
	d.GetObjectsAndKeysCount(&oarr, &karr, 4)
	fmt.Printf("Length of oarr is now %d\n", len(oarr))
	for i, k := range karr {
		fmt.Printf("-- %s -> %s\n", k.NSString(), oarr[i].NSString())
	}
	fmt.Printf("\nNSStringWithContentsOfURLEncoding()\n")
	err := make([]*ns.NSError, 1)
	n1 = ns.NSStringWithContentsOfURLEncoding(ns.NSURLWithGoString("http://captive.apple.com"), ns.NSUTF8StringEncoding, &err)
	if len(err) == 0 {
		fmt.Printf("n1 = %s\n", n1)
	}
	n1 = ns.NSStringWithContentsOfURLEncoding(ns.NSURLWithGoString("htttypo://example.com"), ns.NSUTF8StringEncoding, &err)
	if len(err) > 0 {
		fmt.Printf("err[0] = %p -> %p\n", err[0], err[0].Ptr())
		fmt.Printf("err: %s\n", err[0].LocalizedDescription())
	}

	fmt.Printf("\nNSStringWithFormat()\n")
	str := ns.NSStringWithFormat(nst("(%@) (%@)\n(%@)\n"), n2, n3, s1)
	fmt.Printf("%s\n", str)

	fmt.Printf("\nGlobal strings\n")
	gs1 = ns.NSStringWithGoString("global string 1")
	gs2 = ns.NSStringWithGoString("global string 2")
	fmt.Printf("\nArrayWithObjects\n")
	a2 = ns.NSArrayWithObjects(gs1, gs2)
	a2.ObjectEnumerator().ForIn(func(o *ns.Id) bool {
		fmt.Printf("--%s\n", o.NSString())
		return true
	})
	dir, e := os.Getwd()
	if e != nil {
		fmt.Printf("Failed to get current working directory. %s\n", err)
		os.Exit(-1)
	}
	path := nst(dir)
	filter := ns.NSArrayWithObjects(nst("ast"), nst("yaml"))
	ost := make([]*ns.NSString, 0, 1)
	oar := make([]*ns.NSArray, 0, 1)
	fmt.Printf("\nCompletePathIntoString()\n")
	i := path.CompletePathIntoString(&ost, 0, &oar, filter)
	fmt.Printf("%d matches\n", i)
	if i > 0 {
		fmt.Printf("ost = %s\n", ost[0])
		fmt.Printf("oar =\n")
		oar[0].ObjectEnumerator().ForIn(func(o *ns.Id) bool {
			fmt.Printf("--%s\n", o.NSString())
			return true
		})
	}
}
