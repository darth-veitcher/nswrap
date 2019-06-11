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
		arr := make([]ns.MyClass,1000)
		for i := 0; i < 1000; i++ {
			// Alloc methods set a finalizer that causes the Go GC to
			// Release these objects.
			arr[i] = ns.MyClassAlloc()
			arr[i].ReleaseCallback(releaseX(i))

			// You can still manually retain objects, but that will cause
			// them to stick around after their Go pointers are collected.
			// This may be necessary if you are adding objects to an
			// Objective-C collection?
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
	for {
		o1 := ns.NSStringAlloc().InitWithGoString("one string")

		// NSWrap runs object constructors inside an @autoreleasepool block,
		// and then calls "retain" on them before returning to Go. A Go
		// finalizer is set allowing the Go GC to call Release().

		o2 := ns.NSStringWithGoString("two string") // does not leak

		arr := ns.NSArrayWithObjects(o1,o2)
		_ = arr

		runtime.GC()
		time.Sleep(time.Second/50)
	}
}

func addStr(arr ns.NSMutableArray) {
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
		time.Sleep(time.Second/10)
	}
}

func main() {
	go memtest1()
	go memtest2()
	go memtest3()
	go memtest4()
	select {}
}
