package types

import (
	"fmt"
	"strings"
)

var super map[string]string

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

type Type struct {
	Node *Node
	Class string
	ctype string
	wrapped bool
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
	if err != nil {
		return &Type{}
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

func (t *Type) Wrap() *Type {
	t.wrapped = true
	return t
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
	ct = strings.ReplaceAll(ct," ","_")
	return ct
}

func (t *Type) GoType() string {
	ct := swapstars(t.CType())
	ct = strings.Title(ct)
	ct = strings.ReplaceAll(ct," ","")
	ct = strings.ReplaceAll(ct,"Struct","")
	return ct
}

func (t *Type) CType() string {
	if t.ctype != "" { // cache
		return t.ctype
	}
	ct := t.Node.CtypeSimplified()
	ct = strings.ReplaceAll(ct,"instancename",t.Class)
	ct = strings.ReplaceAll(ct,"instancetype",t.Class + " *")
	if len(ct) > 1 && ct[:2] == "id" {
		ct = "NSObject *" + ct[2:]
	}
	if len(ct) > 11 {
		if ct[:12] == "instancename" { ct = t.Class }
		if ct[:12] == "instancetype" { ct = t.Class + " *" }
	}
	t.ctype = ct
	return ct
}

func (t *Type) GoTypeDecl() string {
	if t.wrapped {
		return t.GoInterfaceDecl()
	}
	return fmt.Sprintf(`
type %s %s
`,t.GoType(),t.CGoType())
}

func (t *Type) GoInterfaceDecl() string {
	super := Super(t.Class)
	if super == "" {
		super = "ptr unsafe.Pointer"
	}
	return fmt.Sprintf(`
type %s struct { %s }
`,t.GoType(), super)
}

func (t *Type) CToGo(cval string) string { // cast C value to CGo
	if t.Node.IsPointer() {
		cval = "unsafe.Pointer(" + cval + ")"
	}
	return fmt.Sprintf("(%s)(%s)",t.GoType(),cval)
}

// Call a C function from Go with a given return type and parameter types
func GoToC(name string, pnames []string, rtype *Type, ptypes []*Type) string {
	var ret strings.Builder
	rt := rtype.CType()
	if rt != "void" {
		if rtype.wrapped {
			ret.WriteString("ret := &" + rtype.GoType()[1:] + "{}\n")
			ret.WriteString("ret.ptr = unsafe.Pointer(")
		} else {
			ret.WriteString("return (" + rtype.GoType() + ")(")
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
		if pt.wrapped {
			p = pn + ".ptr"
		} else {
			p = "(" + pt.CGoType() + ")(" + pn + ")"
		}
		parms = append(parms,p)
	}
	ret.WriteString(strings.Join(parms,", "))
	ret.WriteString(")")
	if rt != "void" {
		if rtype.wrapped {
			ret.WriteString(`)
return ret
`)
		} else {
			ret.WriteString(")")
			if rtype.Node.IsPointer() {
				ret.WriteString(")")
			}
		}
	}
	return ret.String()
}

