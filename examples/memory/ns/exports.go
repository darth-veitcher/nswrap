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

//export MyClassDealloc
func MyClassDealloc(o unsafe.Pointer) {
	cb := MyClassLookup[o].Dealloc
	if cb == nil { return }
	cb()
}

//export MyClassRelease
func MyClassRelease(o unsafe.Pointer) {
	cb := MyClassLookup[o].Release
	if cb == nil { return }
	self := MyClass{}
	self.ptr = o
	super := MyClassSupermethods{
		self.SuperRelease,
	}
	cb(super)
}
