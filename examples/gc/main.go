package main

import "C"
import (
	"fmt"
	"runtime"
	"time"

	"git.wow.st/gmp/nswrap/examples/gc/ns"
)

func releaseX(x int) func(ns.MyClass, ns.MyClassSupermethods) {
	return func(self ns.MyClass, super ns.MyClassSupermethods) {
		//fmt.Printf("--release %d\n", x)
		super.Release() // comment out for leak
	}
}

func memtest1() {
	fmt.Println("memtest1 started")
	for {
		arr := make([]*ns.MyClass, 1000)
		for i := 0; i < 1000; i++ {
			// Alloc methods set a finalizer that causes the Go GC to
			// Release these objects.
			arr[i] = ns.MyClassAlloc()
			arr[i].ReleaseCallback(releaseX(i))

			// You can still manually retain objects, but that will cause
			// them to stick around after their Go pointers are collected.
			//arr[i].Retain() // uncomment for leak
		}
		// Manually run the Go GC at every loop iteration. May not be needed
		// in a real program.
		runtime.GC()
		time.Sleep(time.Second / 50)
		//fmt.Printf("Loop complete\n")
	}
}

func memtest2() {
	fmt.Println("memtest2 started")
	i := 0
	for {
		o1 := ns.NSStringAlloc().InitWithGoString(fmt.Sprintf("two string %d", i))
		o2 := ns.NSStringWithGoString(fmt.Sprintf("two string %d", i))

		// NSWrap runs object constructors inside an @autoreleasepool block,
		// and then calls "retain" on them before returning to Go. A Go
		// finalizer is set allowing the Go GC to call Release().

		o3 := ns.NSStringWithString(o1)
		o4 := ns.NSStringAlloc()
		_ = o4

		a1 := ns.NSArrayAlloc()

		// init methods in Objective-C always return a retained object.
		// init may or may not return the same object that was sent in.

		a1 = a1.InitWithObjects(o1, o2, o3, o4)

		a2 := ns.NSArrayWithObjects(o1, o2, o3, o4)

		// you can always nest alloc and init.

		a3 := ns.NSMutableArrayAlloc().Init()
		a3.AddObject(o1)
		a3.AddObject(o2)
		a3.AddObject(o3)
		a3.AddObject(o4)
		_ = a1
		_ = a2
		_ = a3

		runtime.GC()
		time.Sleep(time.Second / 50)
	}
}

func addStr(arr *ns.NSMutableArray) {
	// temporary strings made by the 'WithGoString' methods should be released
	// automatically by the GC.

	s1 := ns.NSStringAlloc().InitWithGoString("a string")
	arr.AddObject(s1)

	// s1 should be eligible for Go garbage collection here, but is still referenced
	// on the Objective-C side. By adding s1 to an array, the array automatically
	// calls 'retain' on the underlying Objective-C string.
}

func memtest3() {
	fmt.Println("memtest3 started")

	for {
		// arr will be garbage collected by Go
		arr := ns.NSMutableArrayAlloc().Init()
		addStr(arr)
		runtime.GC()
		time.Sleep(time.Second)

		// check that our string was retained.

		s1 := arr.ObjectAtIndex(0)
		gstr := s1.NSString().String()
		_ = gstr
	}
}

func memtest4() {
	fmt.Println("memtest4 started")
	for {
		o1 := ns.NSStringAlloc().InitWithGoString("red string")

		// conversions to UTF8String internally create autoreleased strings
		// in the Objective-C runtime. NSWrap runs these in a mini-
		// @autoreleasepool block.

		c1 := o1.UTF8String()
		_ = o1
		_ = c1
		runtime.GC()
		time.Sleep(time.Second / 50)
		c1.Free() // you need to manually free UTF8Strings
	}
}

func memtest5() {
	fmt.Println("memtest5 started")
	i := 0
	for {
		// by incrementing i we can ensure that Objective-C needs to create
		// a new NSString object at each loop iteration and cannot be reusing
		// the same string object.

		str := ns.NSStringWithGoString(fmt.Sprintf("blue string %d", i))

		// SubstringFromIndex should be returning a newly allocated NSString,
		// which is getting retained by NSWrap and released by a Go GC
		// finalizer.

		sub := str.SubstringFromIndex(5)
		sub2 := sub.Copy().NSString()
		sub3 := sub2.MutableCopy().NSString()
		u := sub.UTF8String()
		u2 := sub2.UTF8String()
		u3 := sub3.UTF8String()
		time.Sleep(time.Second / 50)
		runtime.GC()
		i++
		u.Free()
		u2.Free()
		u3.Free()
		_ = u
		_ = u2
		_ = u3
		//fmt.Printf("loop completed\n")
	}
}

func tmpdict(i int) *ns.NSString {
	o1 := ns.NSStringWithGoString(fmt.Sprintf("temp string 1-%d", i))
	o2 := ns.NSStringWithGoString(fmt.Sprintf("temp string 2-%d", i))
	k1 := ns.NSStringWithGoString(fmt.Sprintf("temp key 1-%d", i))
	k2 := ns.NSStringWithGoString(fmt.Sprintf("temp key 2-%d", i))
	dict := ns.NSDictionaryWithObjectsAndKeys(o1, k1, o2, k2)
	ret := dict.ValueForKey(k1)
	//fmt.Printf("tmpdict(): string = %s\n",ret.NSString())

	defer runtime.GC()    // o1, o2, k1, k2, and dict can be released after we return
	return ret.NSString() // should be retained by NSDictionary.ValueForKey()
}

func tmparr(i int) *ns.NSString {
	o1 := ns.NSStringWithGoString(fmt.Sprintf("temp string 3-%d", i))
	o2 := ns.NSStringWithGoString(fmt.Sprintf("temp string 4-%d", i))
	arr := ns.NSArrayWithObjects(o1, o2)
	os := make([]*ns.Id, 0, 2)
	arr.GetObjectsRange(&os, ns.NSMakeRange(0, 2))

	defer runtime.GC()      // collect o1, o2 and arr
	return os[1].NSString() // should have been retained by NSArray.GetObjects()
}

func memtest6() {
	fmt.Println("memtest6 started")
	i := 0
	for {
		s1 := tmpdict(i)
		s2 := tmparr(i)
		time.Sleep(time.Second / 5)
		u1 := s1.String() // make sure s1 and s2 are still available
		u2 := s2.String()
		e1 := fmt.Sprintf("temp string 1-%d", i)
		if u1 != e1 {
			fmt.Printf("tmpdict() error: %s != %s\n", u1, e1)
		}
		e2 := fmt.Sprintf("temp string 4-%d", i)
		if u2 != e2 {
			fmt.Printf("tmparr() error: %s != %s\n", u2, e2)

		}
		i++
	}
}

func main() {
	fmt.Printf("MultiThreaded? %t\n", ns.NSThreadIsMultiThreaded())
	th := ns.NSThreadNew()
	th.Start()
	fmt.Printf("MultiThreaded? %t\n", ns.NSThreadIsMultiThreaded())
	go memtest1()
	go memtest2()
	go memtest3()
	go memtest4()
	go memtest5()
	go memtest6()

	go func() {
		for {
			// print a progress indicator
			fmt.Printf("t = %s\n", time.Now())
			time.Sleep(time.Second * 10)
		}
	}()
	select {}
}
