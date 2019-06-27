package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
	"unsafe"

	"git.wow.st/gmp/nswrap/examples/strings/ns"
)

func incr() func(bool) (int, float64) {
	i := 0
	b := 0.0
	var mx sync.Mutex
	return func(bad bool) (int, float64) {
		mx.Lock()
		if bad {
			b++
			defer mx.Unlock()
		} else {
			defer func() { i++; mx.Unlock() }()
		}
		if b == 0 {
			return i, 0.0
		} else {
			return i, (b/float64(i)) * 100
		}
	}
}

type tracker struct {
	add, drop func(*ns.Id)
	check func()
	i func(bool) (int, float64)
}

type record struct {
	ptr unsafe.Pointer
	goPtr *ns.Id
	when time.Time
}

func newTracker() (func(*ns.Id), func(*ns.Id), func()) {
	addch := make(chan *ns.Id)
	dropch := make(chan *ns.Id)
	data := []record{}
	var mux sync.Mutex

	go func() {
		for {
			select {
			case x := <-addch:
				mux.Lock()
					data = append(data,record{
						x.Ptr(),
						x,
						time.Now(),
					})
				mux.Unlock()
			case x := <-dropch:
				mux.Lock()
					data = append(data,record{
						nil,
						x,
						time.Now(),
					})
				mux.Unlock()
			}
		}
	}()

	add := func(x *ns.Id) {
		addch<- x
	}
	drop := func(x *ns.Id) {
		dropch<- x
	}
	check := func() {
		live := map[unsafe.Pointer]*ns.Id{}
		bad := false
		mux.Lock()
		for _,r := range data {
			if r.ptr != nil {
				if live[r.ptr] != nil {
					fmt.Printf("COLLISION: %p & %p -> %p\n", r.goPtr, live[r.ptr], r.ptr)
					bad = true
				}
				live[r.ptr] = r.goPtr
			} else {
				delete(live,r.ptr)
			}
		}
		fmt.Printf("Checked %d records -- ",len(data))
		if bad {
			fmt.Printf("failed\n")
		} else {
			fmt.Printf("ok\n")
		}
		mux.Unlock()
	}
	return add,drop,check
}

func mkstrings(t tracker) {
	for {
		//fmt.Printf("main thread: %t\n",ns.NSThreadIsMainThread())
		x,b := t.i(false)
		str := fmt.Sprintf("string %d",x)
		s := ns.NSStringWithGoString(str)
		//t.add(&s.Id)
		for j := 0; j < 10; j++ {
			sout := s.String()
			if str != sout {
				_,b = t.i(true)
				fmt.Printf("%3.2f%% -- %d: '%s' '%s'\n", b, x, str, sout)
			}
			time.Sleep(time.Second/1000)
		}
		if x % 1000 == 0 {
			fmt.Printf("%3.2f%% -- %s\n", b, time.Now().Format("03:04:05.000"))
		}
		//t.drop(&s.Id)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Printf("Starting\n")
	//ns.NSThreadNew().Start()
	fmt.Printf("multithreaded: %t\n", ns.NSThreadIsMultiThreaded())
	//pool := ns.NSAutoreleasePoolAlloc()
	add, drop, check := newTracker()
	i := tracker{add, drop, check, incr()}
	go mkstrings(i)
	go mkstrings(i)
	go mkstrings(i)
	go mkstrings(i)
	go mkstrings(i)
	go mkstrings(i)
	go mkstrings(i)
	go mkstrings(i)
	go func() {
		for {
			runtime.GC()
			time.Sleep(time.Second/100)
		}
	}()

	time.Sleep(time.Second * 600)
	i.check()

	//pool.Drain()
	os.Exit(0)
}
