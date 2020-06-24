package ns


/*
#cgo CFLAGS: -x objective-c -fno-objc-arc
#cgo LDFLAGS: -framework Foundation
#pragma clang diagnostic ignored "-Wformat-security"

#import <Foundation/Foundation.h>


*/
import "C"

import (
	"unsafe"
)

//export C1Dealloc
func C1Dealloc(o unsafe.Pointer) {
	C1Mux.RLock()
	cb := C1Lookup[o].Dealloc
	C1Mux.RUnlock()
	if cb == nil { return }
	self := C1{}
	self.ptr = o
	super := C1Supermethods{
		self.SuperDealloc,
		self.SuperRelease,
	}
	cb(self, super)
}

//export C1Release
func C1Release(o unsafe.Pointer) {
	C1Mux.RLock()
	cb := C1Lookup[o].Release
	C1Mux.RUnlock()
	if cb == nil { return }
	self := C1{}
	self.ptr = o
	super := C1Supermethods{
		self.SuperDealloc,
		self.SuperRelease,
	}
	cb(self, super)
}

//export C2MyMethod
func C2MyMethod(o unsafe.Pointer) {
	C2Mux.RLock()
	cb := C2Lookup[o].MyMethod
	C2Mux.RUnlock()
	if cb == nil { return }
	self := C2{}
	self.ptr = o
	cb(self)
}
