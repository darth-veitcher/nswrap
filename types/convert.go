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

func shouldWrap(gt string) bool {
	return gt != "" && gt[0] == '*' && wrapped[gt[1:]]
}

//goInterfaces records the names of top level Go interfaces. Pointers to these
//are dereferenced to bare interface names.
var goInterfaces map[string]bool

func isGoInterface(gt string) bool {
	return goInterfaces[gt]
}

//TypeParameters maps, for each class, a TypedefName to a type, representing
//the Objective-C type parameters for that class
var TypeParameters map[string]map[string]string

//Typedefs maps from C types to the Type of a typedef with that name.
var typedefs map[string]*Type

func (t *Type) Typedef() *Type {
	return typedefs[t.BaseType().CType()]
}

func init() {
	super = make(map[string]string)
	wrapped = make(map[string]bool)
	goInterfaces = make(map[string]bool)
	TypeParameters = make(map[string]map[string]string)
	typedefs = make(map[string]*Type)
}

func Super(c string) string {
	return super[c]
}

func SetSuper(c, p string) {
	super[c] = p
}

func SetTypeParam(c, n, t string) {
	if TypeParameters[c] == nil {
		TypeParameters[c] = make(map[string]string)
	}
	TypeParameters[c][n] = t
}

func AddTypedef(n,t string) {
	typedefs[n] = NewTypeFromString(t,"")
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

func (t *Type) PointsTo() *Type {
	return NewType(t.Node.PointsTo(), t.Class)
}

func Wrap(s string) {
	// it is the pointers to this type that get wrapped
	wrapped[s] = true
}

func (t *Type) BaseType() *Type {
	ret := NewType( 
		t.Node.BaseType(),
		t.Class,
	)
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
	if len(ct) > 0 && ct[0] == '*' && isGoInterface(ct[1:]) {
		return ct[1:]
	}
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
		extra := t.Node.CtypeSimplified()
		var cgt string
		if td := tp.Typedef(); td != nil {
			cgt = td.CGoType()
			extra = "typedef " + td.Node.Ctype()
		} else {
			cgt = tp.CGoType()
		}
		return fmt.Sprintf(`
//%s (%s)
type %s %s
`,t.Node.Ctype(),extra,gt,cgt)
	}
}

func (t *Type) GoInterfaceDecl() string {
	gt := t.GoType()
	if gt[0] == '*' {
		gt = gt[1:] // dereference wrapped types
	}
	super := Super(gt)
	if super == "" {
		goInterfaces[gt] = true
		return fmt.Sprintf(`
//%s (%s)
type %s interface {
	Ptr() unsafe.Pointer
}
`,t.Node.Ctype(),t.BaseType().GoType(),gt)
	}
	if isGoInterface(super) {
		super = "Id"
	}
	return fmt.Sprintf(`
//%s (%s)
type %s struct { %s }
func (o *%s) Ptr() unsafe.Pointer { return o.ptr }
`,t.Node.Ctype(),t.BaseType().GoType(),gt,super,gt)
}

func (t *Type) IsFunction() bool {
	if td := t.Typedef(); td != nil {
		return td.IsFunction()
	}
	return t.Node.IsFunction()
}

func (t *Type) IsPointer() bool {
	if td := t.Typedef(); td != nil {
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
		rtgt := rtype.GoType()
		if shouldWrap(rtgt) || isGoInterface(rtgt) {
			if isGoInterface(rtgt) {
				rtgt = "Id"
			} else {
				rtgt = rtgt[1:]
			}
			ret.WriteString("	ret := &" + rtgt + "{}\n")
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
		if shouldWrap(pt.GoType()) || isGoInterface(pt.GoType()) {
			p = pn + ".Ptr()"
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
		if shouldWrap(rtype.GoType()) || isGoInterface(rtype.GoType()) {
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

