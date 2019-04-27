package main

import (
	"fmt"
	"gitlab.wow.st/gmp/nswrap/examples/simple/ClassOne"
)

func main() {
	o := ClassOne.NewClassOne().Init()
	fmt.Println("i1 = ",o.Geti1())
	fmt.Println("p1 = ",o.Getp1())
	p1 := o.Getp1()
	fmt.Println("*p1 = ", *p1)
	*p1 = 17
	fmt.Println("*p1 = ", *o.Getp1())
	ns := o.Nstru1()
	np := o.Nstru2()
	fmt.Println(o.Hi1(ns))
	fmt.Println(o.Hi2(np))
}

