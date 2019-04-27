package ClassOne

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#import "simple.h"

ClassOne*
NewClassOne() {
	return [ClassOne alloc];
}

int
ClassOne_geti1(void* obj) {
	return [(id)obj geti1];
}
int*
ClassOne_getp1(void* obj) {
	return [(id)obj getp1];
}
int
ClassOne_hi1(void* obj, struct stru in) {
	return [(id)obj hi1:in];
}
int
ClassOne_hi2(void* obj, void* in) {
	return [(id)obj hi2:in];
}
struct stru
ClassOne_nstru1(void* obj) {
	return [(id)obj nstru1];
}
struct stru*
ClassOne_nstru2(void* obj) {
	return [(id)obj nstru2];
}
ClassOne*
ClassOne_init(void* obj) {
	return [(id)obj init];
}
*/
import "C"

import (
	"unsafe"
)

//ClassOne*
type ClassOne struct { NSObject }

//NSObject*
type NSObject struct { ptr unsafe.Pointer }

//int
type Int C.int

//struct stru
type Stru C.struct_stru

func NewClassOne() *ClassOne {
	ret := &ClassOne{}
	ret.ptr = unsafe.Pointer(C.NewClassOne())
	//ret = ret.Init()
	return ret
}

func (o *ClassOne) Geti1() Int {
	return (Int)(C.ClassOne_geti1(o.ptr))
}


func (o *ClassOne) Getp1() *Int {
	return (*Int)(unsafe.Pointer(C.ClassOne_getp1(o.ptr)))
}


func (o *ClassOne) Hi1(in Stru) Int {
	return (Int)(C.ClassOne_hi1(o.ptr, (C.struct_stru)(in)))
}


func (o *ClassOne) Hi2(in *Stru) Int {
	return (Int)(C.ClassOne_hi2(o.ptr, unsafe.Pointer(in)))
}


func (o *ClassOne) Nstru1() Stru {
	return (Stru)(C.ClassOne_nstru1(o.ptr))
}


func (o *ClassOne) Nstru2() *Stru {
	return (*Stru)(unsafe.Pointer(C.ClassOne_nstru2(o.ptr)))
}


func (o *ClassOne) Init() *ClassOne {
	ret := &ClassOne{}
	ret.ptr = unsafe.Pointer(C.ClassOne_init(o.ptr))
	return ret
}

