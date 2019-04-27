package types

import (
	"fmt"
	"strings"
)

//super is a map recording which class is the parent of each other class
var super map[string]string

//wrapped is a map recording whether a given GoType is to be "wrapped" in a
//go struct.
var wrapped map[string]bool

//TypeParameters maps, for each class, a TypedefName to a type, representing
//the Objective-C type parameters for that class
var TypeParameters map[string]map[string]string

var Typedefs map[string]*Type

func Super(c string) string {
	if super == nil {
		super = make(map[string]string)
	}
	return super[c]
}

func SetSuper(c, p string) {
	if super == nil {
		super = make(map[string]string)
	}
	super[c] = p
}

func SetTypeParam(c, n, t string) {
	if TypeParameters == nil {
		TypeParameters = make(map[string]map[string]string)
	}
	if TypeParameters[c] == nil {
		TypeParameters[c] = make(map[string]string)
	}
	TypeParameters[c][n] = t
}

func AddTypedef(n,t string) {
	if Typedefs == nil {
		Typedefs = make(map[string]*Type)
	}
	Typedefs[n] = NewTypeFromString(t,"")
}

type Type struct {
	Node *Node
	Class string
	ctype string
}

func NewType(n *Node, c string) *Type {
	return &Type{
		Node: n,
		Class: c,
		ctype: "",
	}
}

func NewTypeFromString(t,c string) *Type {
	n,err := Parse(t)
	if n.IsId() {
	//if n.CtypeSimplified() == "id" {
		n,err = Parse("NSObject*")
	}
	if err != nil {
		return &Type{}
	}
	if TypeParameters[c] != nil {
		recur := false
		for k,v := range TypeParameters[c] {
			recur = n.renameTypedefs(k,v)
		}
		if recur {
			return NewTypeFromString(n.Ctype(),c)
		}
	}
	return &Type{
		Node: n,
		Class: c,
		ctype: "",
	}
}

func (t *Type) String() string {
	return t.Node.String()
}

func Wrap(s string) {
	if wrapped == nil {
		wrapped = make(map[string]bool)
	}
	// it is the pointers to this type that get wrapped
	wrapped["*" + s] = true
}

func (t *Type) BaseType() *Type {
	ret := NewType( 
		t.Node.BaseType(),
		t.Class,
	)
	if ret.CType() == ret.Class + " *" { // "instancename"
		ret.ctype = ret.Class
	}
	return ret
}

func swapstars(s string) string {
	for i := len(s) - 1; i > 0 && s[i] == '*'; {
		s = "*" + s[:i]
	}
	return strings.TrimSpace(s)
}

func (t *Type) CGoType() string {
	ct := swapstars("C." + t.CType())
	ct = strings.ReplaceAll(ct,"unsigned ","u")
	ct = strings.ReplaceAll(ct,"signed ","u")
	ct = strings.ReplaceAll(ct,"long ","long")
	ct = strings.ReplaceAll(ct,"complex ","complex")
	ct = strings.ReplaceAll(ct," ","_")
	return ct
}

func (t *Type) GoType() string {
	return _goType(t.CType())
}

func _goType(ct string) string {
	ct = swapstars(ct)
	ct = strings.Title(ct)
	ct = strings.ReplaceAll(ct," ","")
	ct = strings.ReplaceAll(ct,"Struct","")
	return ct
}


func (t *Type) CType() string {
	return t._CType(false)
}

func (t *Type) CTypeAttrib() string {
	return t._CType(true)
}

func (t *Type) _CType(attrib bool) string {
	//if !attrib && c.ctype != "" ... FIXME?
	if t.ctype != "" { // cache
		return t.ctype
	}
	var ct string
	if attrib {
		ignore := map[string]bool { "GenericList": true }
		ct = t.Node._Ctype(ignore)
	} else {
		ct = t.Node.CtypeSimplified()
	}
	ct = strings.ReplaceAll(ct,"instancename",t.Class)
	ct = strings.ReplaceAll(ct,"instancetype",t.Class + " *")
	if len(ct) > 1 && ct[:2] == "id" {
		ct = "NSObject*" + ct[2:]
	}
	if len(ct) > 11 {
		if ct[:12] == "instancename" { ct = t.Class + ct[12:] }
		if ct[:12] == "instancetype" { ct = t.Class + ct[12:] + " *" }
	}
	if attrib {
		t._CType(false)
	} else {
		t.ctype = ct
	}
	return ct
}

func (t *Type) GoTypeDecl() string {
	if wrapped[t.GoType()] {
		return t.GoInterfaceDecl()
	}
	tp := t.BaseType()
	if tp.Node.IsId() {
		return ""
	}
	gt := tp.GoType()
	switch gt {
	case "", "Void":
		return ""
	default:
		return fmt.Sprintf(`
//%s
type %s %s
`,t.Node.CtypeSimplified(),gt,tp.CGoType())
	}
}

func (t *Type) GoInterfaceDecl() string {
	gt := t.GoType()
	if gt[0] == '*' {
		gt = gt[1:] // dereference wrapped types
	}
	x := ""
	super := Super(gt)
	if super == "" {
		super = "ptr unsafe.Pointer"
	}
	return fmt.Sprintf(`
//%s
%stype %s struct { %s }
`,t.CTypeAttrib(),x,gt,super)
}

func (t *Type) IsFunction() bool {
	if td,ok := Typedefs[t.BaseType().CType()]; ok {
		return td.IsFunction()
	}
	return t.Node.IsFunction()
}

func (t *Type) IsPointer() bool {
	if td,ok := Typedefs[t.BaseType().CType()]; ok {
		return td.IsPointer()
	}
	return t.Node.IsPointer()
}

func (t *Type) CToGo(cval string) string { // cast C value to CGo
	if t.IsPointer() {
		cval = "unsafe.Pointer(" + cval + ")"
	}
	return fmt.Sprintf("(%s)(%s)",t.GoType(),cval)
}

// Call a C function from Go with a given return type and parameter types
func GoToC(name string, pnames []string, rtype *Type, ptypes []*Type) string {
	var ret strings.Builder
	rt := rtype.CType()
	if rt != "void" {
		if wrapped[rtype.GoType()] {
			ret.WriteString("	ret := &" + rtype.GoType()[1:] + "{}\n")
			ret.WriteString("	ret.ptr = unsafe.Pointer(")
		} else {
			ret.WriteString("	return (" + rtype.GoType() + ")(")
			if rtype.Node.IsPointer() {
				ret.WriteString("unsafe.Pointer(")
			}
		}
	}
	ret.WriteString("C." + name + "(")
	parms := []string{}
	for i := 0; i < len(pnames); i++ {
		pn,pt := pnames[i],ptypes[i]
		p := pn
		if wrapped[pt.GoType()] {
			p = pn + ".ptr"
		} else {
			if pt.Node.IsPointer() {
				p = "unsafe.Pointer(" + pn + ")"
			} else {
				p = "(" + pt.CGoType() + ")(" + pn + ")"
			}
		}
		parms = append(parms,p)
	}
	ret.WriteString(strings.Join(parms,", "))
	ret.WriteString(")")
	if rt != "void" {
		if wrapped[rtype.GoType()] {
			ret.WriteString(`)
	return ret
`)
		} else {
			ret.WriteString(")")
			if rtype.Node.IsPointer() {
				ret.WriteString(")")
			}
			ret.WriteString("\n")
		}
	}
	return ret.String()
}

