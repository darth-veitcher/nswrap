package main

//go:generate nswrap

import (
	"fmt"
	ns "git.wow.st/gmp/nswrap/examples/simple/ClassOne"
)

func cb(self ns.ClassThree, super ns.ClassThreeSupermethods) ns.Int {
	fmt.Printf("In Go callback\n")
	return 0
}

func main() {
	o := ns.ClassOneAlloc().Init()
	fmt.Println("i1 = ", o.Geti1())
	fmt.Println("p1 = ", o.Getp1())
	p1 := o.Getp1()
	fmt.Println("*p1 = ", *p1)
	*p1 = 17
	fmt.Println("*p1 = ", *o.Getp1())
	ns1 := o.Nstru1()
	np := o.Nstru2()
	fmt.Println(o.Hi1(ns1))
	fmt.Println(o.Hi2(np))
	o2 := ns.ClassTwoAlloc().Init()
	fmt.Println(o2.Hi1(ns1))
	o3 := ns.ClassThreeAlloc()
	o3.Init()
	o3.Geti1Callback(cb)
	fmt.Println(o3.Hi2(np))
	fmt.Println(o3.Geti1())
}
