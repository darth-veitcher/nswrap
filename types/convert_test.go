package types

import (
	"fmt"
	"reflect"
	"testing"
)

func isNil(x interface{}) bool {
	return x == nil ||
		(reflect.ValueOf(x).Kind() == reflect.Ptr &&
		reflect.ValueOf(x).IsNil())
}

func TestType(t *testing.T) {
	i := 1

	var str string
	var n *Node
	var tp *Type

	chk := func(actual, expected interface{}) {
		t.Run(fmt.Sprintf("TestType%d",i),func(t *testing.T) {
			if isNil(actual) && isNil(expected) { return }
			if !reflect.DeepEqual(actual,expected) {
				t.Errorf("Test failed for %s\n",str)
				fmt.Println("Actual:\n",actual)
				fmt.Println("Expected:\n",expected)
			}
		})
	}

	chk_newtype := func() {
		tp = NewTypeFromString(str,"")
		chk(tp,&Type{ Node: n, Class: "" })
	}

	//tests on nil Type pointers:
	chk(tp.BaseType(),nil)
	chk(tp.CType(),"")
	chk(tp.IsFunction(),false)
	chk(tp.IsPointer(),false)
	chk(tp.IsFunctionPtr(),false)
	chk(tp.IsValist(),false)

	str = "int"
	n = &Node{ "TypeName", "", []*Node{
		&Node{ "TypeSpecifier", "int", []*Node{} } } }
	chk_newtype()
	tint := tp

	str = "NSObject"
	n = &Node{ "TypeName", "", []*Node{
		&Node{ "TypedefName", "NSObject", []*Node{} } } }
	chk_newtype()
	nso := tp

	str = "NSString"
	n = &Node{ "TypeName", "", []*Node{
		&Node{ "TypedefName", "NSString", []*Node{} } } }
	chk_newtype()
	nst := tp

	str = "NSString*"
	n = &Node{ "TypeName", "", []*Node{
		&Node{ "TypedefName", "NSString", []*Node{} },
		&Node{ "Pointer", "*", []*Node{} } } }
	chk_newtype()
	chk(tp.IsPointer(),true)
	chk(tp.Typedef(),nil)
	tpNSString := tp

	str = "myTypedef"
	AddTypedef("myTypedef",tp)
	n = &Node { "TypeName", "", []*Node{
		&Node{ "TypedefName", "myTypedef", []*Node{} }}}
	chk_newtype()
	chk(tp.Typedef(),tpNSString)

	str = "const NSArray <ObjectType * _Nonnull> *"
	n = &Node{ "TypeName", "", []*Node{
		&Node{ "TypeQualifier", "const", []*Node{} },
		&Node{ "TypedefName", "NSArray", []*Node{} },
		&Node{ "GenericList", "", []*Node{
			&Node{ "Generic", "", []*Node{
				&Node{ "TypedefName", "ObjectType", []*Node{} },
				&Node{ "Pointer", "*", []*Node{} },
				&Node{ "NullableAnnotation", "_Nonnull", []*Node{} } } } } },
		&Node{ "Pointer", "*", []*Node{} } } }
	chk_newtype()
	chk(tp.CType(),"NSArray*")
	chk(tp.CTypeAttrib(),"const NSArray*")
	chk(tp.CGoType(),"*C.NSArray")
	chk(tp.GoType(),"*NSArray")

	SetTypeParam("MyClass","ObjectType","MyClass")
	str = "id<ObjectType *>"
	n = &Node{ "TypeName", "", []*Node{
		&Node{ "TypedefName", "id", []*Node{} },
		&Node{ "GenericList", "", []*Node{
			&Node{ "Generic", "", []*Node{
				&Node{ "TypedefName", "ObjectType", []*Node{} },
				&Node{ "Pointer", "*", []*Node{} } } } } } } }
	chk_newtype()
	tp = NewType(n,"MyClass")
	chk(tp.CType(),"NSObject*")
	chk(tp.String(),
`<TypeName> ''
-<TypedefName> 'id'
-<GenericList> ''
--<Generic> ''
---<TypedefName> 'MyClass'
---<Pointer> '*'
`)
	x,_ := clean(nil,"MyClass")
	chk(x,nil)
	SetSuper("NSString","NSObject")
	chk(Super("NSString"),"NSObject")

	tp2 := tp.CloneToClass("NSObject")
	chk(tp2.Class,"NSObject")

	str = "you can't parse this"
	tp = NewTypeFromString(str,"")
	tp2 = &Type{}
	chk(tp,tp2)

	str = "id<ObjectType *>"
	tp2 = NewTypeFromString(str,"MyClass")
	chk(tp.BaseType(),tp)

	str = "id<ObjectType *>*"
	tp = NewTypeFromString(str,"MyClass")
	chk(tp.BaseType(),tp2)
	chk(tp.PointsTo(),tp2)
	AddTypedef("myTypedef",tp)

	str = "myTypedef"
	n = &Node{ "TypeName", "", []*Node{
		&Node{ "TypedefName", "myTypedef", []*Node{} } } }
	chk_newtype()
	chk(tp.PointsTo(),tp2)
	chk(tp2.PointsTo(),nil)

	chk(tp.GoTypeDecl(),`
type MyTypedef **C.NSObject
`)
	chk(tp2.GoTypeDecl(),"")
	chk(nst.GoTypeDecl(),`
type NSString C.NSString
`)
	str = "void"
	n = &Node{ "TypeName", "", []*Node{
		&Node{ "TypeSpecifier", "void", []*Node{} } } }
	chk_newtype()
	chk(tp.GoTypeDecl(),"")
	void := tp

	str = "BOOL"
	n,_ = Parse(str)
	chk_newtype()
	bl := tp

	str = "void**"
	n,_ = Parse(str)
	chk_newtype()
	chk(tp.GoType(),"*unsafe.Pointer")
	chk(tp.CToGo("var"),"(*unsafe.Pointer)(unsafe.Pointer(var))")
	voidpp := tp

	Wrap("NSObject")
	str = "NSObject*"
	n,_ = Parse(str)
	chk_newtype()
	chk(tp.GoType(),"NSObject")

	Wrap("NSString")
	chk(nso.GoTypeDecl(),`
type NSObject interface {
	Ptr() unsafe.Pointer
}
`)
	chk(nso.GoType(),"NSObject")

	chk(nst.GoTypeDecl(),`
type NSString struct { Id }
func (o NSString) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSString() NSString {
	ret := NSString{}
	ret.ptr = o.ptr
	return ret
}
`)
	str = "int(void)"
	n,_ = Parse(str)
	chk_newtype()
	chk(tp.IsFunction(),true)
	chk(tp.IsFunctionPtr(),false)
	chk(tp.ReturnType().CType(),"int")

	str = "int *(void)"
	n,_ = Parse(str)
	chk_newtype()
	chk(tp.IsFunction(),true)
	chk(tp.ReturnType().CType(),"int*")
	chk(tp.ReturnType().GoType(),"*Int")
	fn := tp

	AddTypedef("myTypedef",fn)
	str = "myTypedef"
	n,_ = Parse(str)
	chk_newtype()
	chk(tp.IsFunction(),true)
	chk(tp.IsFunctionPtr(),false)
	chk(tp.ReturnType().CType(),"int*")

	str = "int (*)(void)"
	n,_ = Parse(str)
	chk_newtype()
	chk(tp.IsFunction(),false)
	chk(tp.IsPointer(),true)
	chk(tp.IsFunctionPtr(),true)
	chk(tp.ReturnType(),nil)
	fnptr := tp

	AddTypedef("myTypedef",fnptr)
	str = "myTypedef"
	n,_ = Parse(str)
	chk_newtype()
	chk(tp.IsFunction(),false)
	chk(tp.IsPointer(),true)
	chk(tp.IsFunctionPtr(),true)
	chk(tp.ReturnType(),nil)

	chk(tp.IsValist(),false)
	str = "__va_list_tag"
	n,_ = Parse(str)
	chk_newtype()
	chk(tp.IsValist(),true)

	str = "GoToC"
	var rtype *Type
	ptypes := []*Type{ nst, nso, tint, voidpp }
	pnames := []string{"p1","p2","p3", "p4"}
	snames := []string{"","","",""}

	chk_gotoc := func(expected string) {
		chk(GoToC("myFun",pnames,snames,rtype,ptypes,false),expected)
	}

	chk_gotoc("")

	rtype = void
	chk_gotoc(`C.myFun(p1.Ptr(), p2.Ptr(), (C.int)(p3), unsafe.Pointer(p4))`)

	rtype = bl
	chk_gotoc(
`ret := (C.myFun(p1.Ptr(), p2.Ptr(), (C.int)(p3), unsafe.Pointer(p4))) != 0
	return ret`)

	rtype = voidpp
	chk_gotoc(
`ret := (*unsafe.Pointer)(unsafe.Pointer(C.myFun(p1.Ptr(), p2.Ptr(), (C.int)(p3), unsafe.Pointer(p4))))
	return ret`)

	rtype = nst
	chk_gotoc(
`ret := NSString{}
	ret.ptr = C.myFun(p1.Ptr(), p2.Ptr(), (C.int)(p3), unsafe.Pointer(p4))
	return ret`)

	rtype = nso
	chk_gotoc(
`ret := Id{}
	ret.ptr = C.myFun(p1.Ptr(), p2.Ptr(), (C.int)(p3), unsafe.Pointer(p4))
	return ret`)

	ptypes[1].Variadic = true
	chk_gotoc(
`ret := Id{}
	ret.ptr = C.myFun(p1.Ptr(), unsafe.Pointer(&p2), (C.int)(p3), unsafe.Pointer(p4))
	return ret`)
	ptypes[1].Variadic = false

	snames[1] = "p2p"
	chk_gotoc(
`ret := Id{}
	ret.ptr = C.myFun(p1.Ptr(), p2.Ptr(), (C.int)(p3), unsafe.Pointer(p4))
	(*p2) = (*p2)[:cap(*p2)]
	for i := 0; i < len(*p2); i++ {
		if p2p[i] == nil {
			(*p2) = (*p2)[:i]
			break
		}
		(*p2)[i].ptr = p2p[i]
	}
	return ret`)
	snames[1] = ""
	snames[2] = "p3p"
	chk_gotoc(
`ret := Id{}
	ret.ptr = C.myFun(p1.Ptr(), p2.Ptr(), unsafe.Pointer(&p3p[0]), unsafe.Pointer(p4))
	(*p3) = (*p3)[:cap(*p3)]
	for i := 0; i < len(*p3); i++ {
		if p3p[i] == nil {
			(*p3) = (*p3)[:i]
			break
		}
		(*p3)[i].ptr = p3p[i]
	}
	return ret`)
	chk(GoToC("myFun",pnames,snames,rtype,ptypes,true),
`ret := Id{}
	ret.ptr = C.myFun(p1.Ptr(), p2.Ptr(), unsafe.Pointer(&p3p[0]), (**C.void)(p4))
	(*p3) = (*p3)[:cap(*p3)]
	for i := 0; i < len(*p3); i++ {
		if p3p[i] == nil {
			(*p3) = (*p3)[:i]
			break
		}
		(*p3)[i].ptr = p3p[i]
	}
	return ret`)
}