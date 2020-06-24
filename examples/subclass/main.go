package main

import (
	"fmt"

	"git.wow.st/gmp/nswrap/examples/subclass/ns"
)

func c1release(self ns.C1, super ns.C1Supermethods) {
	fmt.Printf("c1release()\n")
	super.Release()
	fmt.Printf("c1release() done\n")
}

func c2myMethod(self ns.C2) {
	fmt.Printf("c2myMethod()\n")
}

func main() {
	fmt.Printf("Starting\n")

	c1 := ns.C1Alloc()
	c1.ReleaseCallback(c1release)
	c1.Release()

	c2 := ns.C2Alloc()
	c2.MyMethodCallback(c2myMethod)
	c2.Release()

	fmt.Printf("Done\n")
}
