package main

import "C"
import (
	"fmt"
	"runtime"
	"time"

	"git.wow.st/gmp/nswrap/examples/gc/ns"
)

func releaseX(x int) func (ns.MyClassSupermethods) {
	return func(super ns.MyClassSupermethods) {
		//fmt.Printf("--release %d\n", x)
		super.Release() // comment out for leak
	}
}
func memtest1() {
	fmt.Println("memtest1 started")
	for {
		arr := make([]*ns.MyClass,1000)
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
		time.Sleep(time.Second/50)
		fmt.Printf("Loop complete\n")
	}
}

func memtest2() {
	fmt.Println("memtest2 started")
	i := 0
	for {
		o1 := ns.NSStringAlloc().InitWithGoString(fmt.Sprintf("two string %d",i))
		o2 := ns.NSStringWithGoString(fmt.Sprintf("two string %d",i))

		// NSWrap runs object constructors inside an @autoreleasepool block,
		// and then calls "retain" on them before returning to Go. A Go
		// finalizer is set allowing the Go GC to call Release().

		o3 := ns.NSStringWithString(o1)
		o4 := ns.NSStringAlloc()
		_ = o4

		//arr := ns.NSArrayAlloc().InitWithObjects(o1,o1)
		arr := ns.NSArrayWithObjects(o1,o2,o3,o4)
		_ = arr

		//o1.Release()
		//o1.Release()
		runtime.GC()
		time.Sleep(time.Second/50)
	}
}

func addStr(arr *ns.NSMutableArray) {
	s1 := ns.NSStringAlloc().InitWithGoString("a string")
	arr.AddObject(s1)

	// s1 should be eligible for Go garbage collection here, but is still referenced
	// on the Objective-C side. By adding s1 to an array, the array automatically
	// calls 'retain' on the underlying Objective-C string.
}

func memtest3() {
	fmt.Println("memtest3 started")

	arr := ns.NSMutableArrayAlloc().Init() // arr will be garbage collected by Go
	addStr(arr)
	runtime.GC()
	time.Sleep(time.Second)
	s1 := arr.ObjectAtIndex(0)
	fmt.Println(s1.NSString().UTF8String())
	fmt.Println("memtest3 done")
}

func memtest4() {
	fmt.Println("memtest4 started")
	for {
		o1 := ns.NSStringAlloc().InitWithGoString("four string")
		c1 := o1.UTF8String()
		_ = o1
		_ = c1
		runtime.GC()
		time.Sleep(time.Second/10)
	}
}

func memtest5() {
	fmt.Println("memtest5 started")
	i := 0
	for {
		str := ns.NSStringWithGoString(fmt.Sprintf("five string %d",i))
		_ = str
		sub := str.SubstringFromIndex(5)
		_ = sub
		fmt.Printf("sub = %s\n",sub)
		time.Sleep(time.Second/10)
		runtime.GC()
		i++
	}
}

func main() {
	go memtest1()
	go memtest2()
	go memtest3()
	go memtest4()
	go memtest5()
	select {}
}
