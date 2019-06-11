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

//export MyClassRelease
func MyClassRelease(o unsafe.Pointer) {
	MyClassMux.RLock()
	cb := MyClassLookup[o].Release
	MyClassMux.RUnlock()
	if cb == nil { return }
	self := MyClass{}
	self.ptr = o
	super := MyClassSupermethods{
		self.SuperRelease,
	}
	cb(super)
}

//export MyClassDealloc
func MyClassDealloc(o unsafe.Pointer) {
	MyClassMux.RLock()
	cb := MyClassLookup[o].Dealloc
	MyClassMux.RUnlock()
	if cb == nil { return }
	cb()
}
