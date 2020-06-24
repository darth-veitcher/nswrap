package main

import (
	"fmt"
	"git.wow.st/gmp/nswrap/examples/functions/ns"
	"unsafe"
)

func main() {
	var s ns.Stat
	ns.Puts(ns.CharWithGoString("Hi there"))
	ret := ns.Fstat(3, &s)
	fmt.Printf("Fstat: %d\n", ret)

	fmt.Printf("Opening file\n")
	f := ns.Fopen(ns.CharWithGoString("nswrap.yaml"), ns.CharWithGoString("r"))
	ret = ns.Fstat(3, &s)
	fmt.Printf("Fstat: %d\n", ret)
	chr := make([]byte, 4096)
	i := ns.Fread(unsafe.Pointer(&chr[0]), 1, 4096, f)
	if i < 4096 {
		chr = chr[:i]
	}
	fmt.Printf("file contents:\n%s\n", string(chr))
}
