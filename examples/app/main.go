package main

import (
	"gitlab.wow.st/gmp/nswrap/examples/app/ns"
)

func main() {
	a := ns.NSApplicationSharedApplication()
	a.SetActivationPolicy(ns.NSApplicationActivationPolicyRegular)
	//w := ns.NSWindowAlloc()
	w := ns.NSWindowAlloc().InitWithContentRect(
		ns.NSMakeRect(200,200,600,600),
		ns.NSWindowStyleMaskTitled,
		ns.NSBackingStoreBuffered,
		0,
		nil,
	)
	w.SetTitle(ns.NSStringWithGoString("Hi World"))
	w.MakeKeyAndOrderFront(w)
	a.Run()
}

