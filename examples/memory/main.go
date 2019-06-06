package main

import "C"
import (
	"fmt"
	"runtime"
	"time"

	"git.wow.st/gmp/nswrap/examples/memory/ns"
)

func dealloc() {
	//[super dealloc] is called for you automatically, so no Supermethods
	//struct is provided here.
	fmt.Println("--dealloc called")
}

func release(super ns.MyClassSupermethods) {
	fmt.Println("--release called")

	super.Release() // comment out for leak
}

//Basic memory allocation test using a manual Autorelease pool. Also utilizes
//a custom object that overrides dealloc and release methods. Make sure you
//call [super release] or you will have a leak (see above). [super dealloc] is
//called for youautomatically as it is basically always required.
func memtest1() {

	//because time.Sleep() is called within each loop, it is necessary
	//to lock this goroutine to a thread. Otherwise, Sleep can return
	//and continue execution on a different thread, causing the risk that
	//the next call to NSAutoreleasePool.Drain() seg faults because it is
	//not in the same thread where the NSAutoreleasePool was allocated.

	runtime.LockOSThread()

	fmt.Println("memtest1: started")
	for {
		pool := ns.NSAutoreleasePoolAlloc().Init()
		o1 := ns.MyClassAlloc()
	//If autorelease: true is set in nswrap.yaml, the manual calls to
	//autorelease are not necessary.
		o1.Autorelease()
		o1.DeallocCallback(dealloc)
		o1.ReleaseCallback(release)
		o2 := ns.NSObjectAlloc()
		o2.Autorelease()
		o3 := ns.NSMutableArrayAlloc()
		o3.Autorelease()
		o4 := ns.NSStringAlloc()
		o4.Autorelease()
		_ = o1
		_ = o2
		_ = o3
		_ = o4
		pool.Drain()
		time.Sleep(time.Second/2)
	}
	fmt.Println("memtest1: done")
}

//Test the ns.Autoreleasepool() function. Also confirm that we do not need
//to release or manually autorelease objects returned by constructor methods
//(i.e. not created with *Alloc()).
func memtest2() {
	runtime.LockOSThread()
	fmt.Println("memtest2: started")
	i := 0
	for {
		ns.Autoreleasepool(func() {
			o1 := ns.NSObjectAlloc()
			o1.Autorelease()
			s1 := ns.NSStringWithGoString(fmt.Sprintf("string-%d",i))
			_ = s1
			//o1.Retain() // uncomment for leak
		})
		time.Sleep(time.Second/3)
	}
	fmt.Println("memtest2: done")
}

//Test nested Autoreleasepool invocations -- confirms that objects in the
//outer pool are not deallocated by the inner pool.
func memtest3() {

	runtime.LockOSThread() // comment out for crash

	fmt.Println("memtest3: started")
	for { ns.Autoreleasepool(func() {
		arr := ns.NSMutableArrayAlloc().Init()
		arr.Autorelease()
		arr.AddObject(ns.NSStringWithGoString("my string"))

		for { ns.Autoreleasepool(func() {
			str := arr.ObjectAtIndex(0).NSString()
			fmt.Println(str) // does not leak in an autorelease pool
			time.Sleep(time.Second / 2)
		})}
		time.Sleep(time.Second)
	})}
	fmt.Println("memtest3: done")
}

//Test of manual memory management. Lets run multiple goroutines here to
//confirm we can use multiple threads if autorelease pools are not in play.
func memtest4() {
	go memtest4a()
	go memtest4a()
	go memtest4a()
	go memtest4b()
	go memtest4b()
	go memtest4b()
	go memtest4c()
	go memtest4c()
	go memtest4c()
	// Exactly one goroutine (locked to an OS thread) can use an
	// autorelease pool (?)
	go memtest1()
}

func memtest4a() {
	for {
		o1 := ns.NSObjectAlloc()
		o1.Init()
		time.Sleep(time.Second/50)
		o1.Release()
	}
}
func memtest4b() {
	for {
		o1 := ns.NSObjectAlloc() // need to Release
		s1 := ns.NSStringWithGoString("a string")
		arr := ns.NSArrayAlloc().InitWithObjects(s1) // need to Release
		s2 := arr.ObjectAtIndex(0).NSString()

	//If you try to convert an NSString to UTF8String, CString (*Char),
	//or GoString, Objective-C runtime will leak an
	//NSTaggedPointerCStringContainer. Don't know why or how to fix it.
	//There would be no leak if we were using an autorelease pool.
		//u := str.UTF8String() // uncomment for leak
		//fmt.Println(s1) // uncomment for leak

		time.Sleep(time.Second/50)

		o1.Release()
		arr.Release()
		_ = o1
		_ = s2
	}
}
func memtest4c() {
	for {
		o1 := ns.NSArrayAlloc()
		o2 := ns.NSStringAlloc()

		time.Sleep(time.Second/50)
		o1.Release()
		o2.Release()
	}
}


func main() {
	//Uncomment more than one test at a time for a crash.
	//Note: You may not run autorelease pools from multiple goroutines.
	//Within an autorelease pool, do not do anything that can result in a
	//switch to a different thread.

	//go memtest1()
	//go memtest2()
	go memtest3()
	//go memtest4()
	select { }
}
