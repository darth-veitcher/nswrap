// An example of manual memory management (nogc directive in nswrap.yaml)
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
		time.Sleep(time.Second / 2)
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
			//Does not leak strings within an autorelease pool
			s1 := ns.NSStringWithGoString(fmt.Sprintf("string-%d", i))
			_ = s1
			//o1.Retain() // uncomment for leak
			i++
		})
		time.Sleep(time.Second / 3)
	}
	fmt.Println("memtest2: done")
}

//Test nested Autoreleasepool invocations -- confirms that objects in the
//outer pool are not deallocated by the inner pool.
func memtest3() {

	runtime.LockOSThread() // comment out for crash

	fmt.Println("memtest3: started")
	i := 0
	for {
		ns.Autoreleasepool(func() {
			arr := ns.NSMutableArrayAlloc().Init()
			arr.Autorelease()
			arr.AddObject(ns.NSStringWithGoString(fmt.Sprintf("my string %d",i)))
			s1 := ns.NSStringWithGoString(fmt.Sprintf("my other string %d",i))
			fmt.Printf("%s\n",arr.ObjectAtIndex(0).NSString())
			_ = s1

			for x := 0; x < 3; x++ {
				ns.Autoreleasepool(func() {
					str := arr.ObjectAtIndex(0).NSString()
					fmt.Printf("%d->%s\n",x,str) // does not leak in an autorelease pool
					time.Sleep(time.Second / 5)
				})
			}
			time.Sleep(time.Second/2)
			i++
		})
	}
	fmt.Println("memtest3: done")
}

//Test of manual memory management.
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
	// running multiple separate threads with autorelease pools...
	go memtest1()
	go memtest2()
	go memtest3()
	go memtest1()
	go memtest2()
	go memtest3()
	go memtest1()
	go memtest2()
	go memtest3()
	go memtest1()
	go memtest2()
	go memtest3()
}

func memtest4a() {
	for {
		o1 := ns.NSObjectAlloc()
		o1.Init()
		time.Sleep(time.Second / 50)
		o1.Release()
	}
}
func memtest4b() {
	i := 0
	for {
		o1 := ns.NSObjectAlloc() // need to Release

		// These object constructors will always leak. In the case of an
		// immutable string, it is not an issue unless a large number of different
		// strings are going to be made.
		//s1 := ns.NSStringWithGoString(fmt.Sprintf("a string %d",i)) // uncomment for leak
		s1 := ns.NSStringWithGoString("a string")
		i++
		arr := ns.NSArrayAlloc().InitWithObjects(s1) // need to Release arr
		s2 := arr.ObjectAtIndex(0).NSString()

		// If you try to convert an NSString to UTF8String, CString (*Char),
		// or GoString, Objective-C runtime creates an autoreleased
		// NSTaggedPointerCStringContainer internally and there is no way to
		// release it unless you called your method from within an autorelease
		// pool. The following two calls cause a leak.
		//u := s1.UTF8String() // uncomment for leak
		//fmt.Println(s1) // uncomment for leak

		time.Sleep(time.Second / 50)

		o1.Release()
		arr.Release()

		//s1.Release() // does not prevent the leak caused by
		// NSStringWithGoString: s1 is autoreleased by the Objective-C
		// runtime, these methods will always leak without autorelease
		// pools.

		_ = o1
		_ = s2
	}
}
func memtest4c() {
	for {
		o1 := ns.NSArrayAlloc()
		o2 := ns.NSStringAlloc()

		time.Sleep(time.Second / 50)
		o1.Release()
		o2.Release()
	}
}

func main() {
	//Uncomment more than one test at a time for a crash.
	//Note: You may not run autorelease pools from multiple goroutines.
	//Within an autorelease pool, do not do anything that can result in a
	//switch to a different thread.

	go memtest4()
	select {}
}
