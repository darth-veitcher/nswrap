package types

import (
	"fmt"
	"regexp"
	"strings"
)

//super is a map recording which class is the parent of each other class
var super map[string]string

//wrapped is a map recording whether a given GoType is to be "wrapped" in a
//go struct.
var wrapped map[string]bool

func ShouldWrap(gt string) bool {
	return wrapped[gt]
}

//goInterfaces records the names of top level Go interfaces.
var goInterfaces map[string]bool

func IsGoInterface(gt string) bool {
	return goInterfaces[gt]
}


//TypeParameters maps, for each class, a TypedefName to a type, representing
//the Objective-C type parameters for that class
var TypeParameters map[string]map[string]string

//Typedefs maps from C types to the Type of a typedef with that name.
var typedefs map[string]*Type

func (t *Type) Typedef() *Type {
	return typedefs[t.CType()]
}

var (
	r_id *regexp.Regexp
	r_instancename *regexp.Regexp
	r_instancetype *regexp.Regexp
)

func init() {
	super = make(map[string]string)
	wrapped = make(map[string]bool)
	goInterfaces = make(map[string]bool)
	TypeParameters = make(map[string]map[string]string)
	typedefs = make(map[string]*Type)

	r_id = regexp.MustCompile(`\bid\b`)
	r_instancename = regexp.MustCompile(`\binstancename\b`)
	r_instancetype = regexp.MustCompile(`\binstancetype\b`)
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

func AddTypedef(n string, tp *Type) {
	//fmt.Printf("AddTypedef(): %s -> %s\n",n,t)
	typedefs[n] = tp
}

type Type struct {
	Node *Node
	Class string
	Variadic bool
}

func (t *Type) CloneToClass(c string) *Type {
	return &Type{
		Node: t.Node,
		Class: c,
		Variadic: t.Variadic,
	}
}

func clean(n *Node,c string) (*Node,bool) {
	if n == nil {
		return nil,false
	}
	ret := NewNode(n.Kind,n.Content)
	ret.Children = n.Children
	//fmt.Printf("clean(%s,%s)\n",n.CType(),c)
	recur := false
	if TypeParameters[c] != nil {
		for k,v := range TypeParameters[c] {
			recur = ret.renameTypedefs(k,v)
		}
	}
	if recur {
		clean(n, c)
		return ret,true
	}
	return n,false
}

func NewType(n *Node, c string) *Type {
	n2,_ := clean(n, c)
	return &Type{
		Node: n2,
		Class: c,
	}
}

func NewTypeFromString(t,c string) *Type {
	//fmt.Printf("t/c: %s/%s\n",t,c)
	n,err := Parse(t)
	//fmt.Printf("%p %s",n,n.String())
	if err != nil {
		return &Type{}
	}
	if n2,ok := clean(n, c); ok {
		//found type parameters, re-parse
		return NewTypeFromString(n2.CType(),c)
	}
	return &Type{
		Node: n,
		Class: c,
	}
}

func (t *Type) String() string {
	return t.Node.String()
}

func (t *Type) PointsTo() *Type {
	if td := t.Typedef(); td != nil {
		return td.PointsTo()
	}
	if pt := t.Node.PointsTo(); pt != nil {
		return NewType(t.Node.PointsTo(), t.Class)
	} else {
		return nil
	}
}

func Wrap(s string) {
	wrapped[s] = true
}

func (t *Type) BaseType() *Type {
	if t == nil {
		return nil
	}
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
	if IsGoInterface(ct) {
		return ct
	}
	if ct == "Id" {
		ct = "Id"
	}
	if len(ct) > 1 && ShouldWrap(ct[1:]) {
		return ct[1:]
	}
	if len(ct) > 4 && ct[len(ct)-4:len(ct)] == "Void" {
		ct = ct[:len(ct)-5] + "unsafe.Pointer"
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
	if t == nil {
		//fmt.Println("nil sent to _CType()")
		return ""
	}
	var ct string
	if attrib {
		ignore := map[string]bool { "GenericList": true }
		ct = t.Node._CType(ignore)
	} else {
		ct = t.Node.CTypeSimplified()
	}
	ct = r_id.ReplaceAllString(ct,"NSObject*")
	ct = r_instancename.ReplaceAllString(ct,t.Class)
	ct = r_instancetype.ReplaceAllString(ct,t.Class + "*")
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
		var cgt string
		if td := tp.Typedef(); td != nil {
			cgt = td.CGoType()
		} else {
			cgt = tp.CGoType()
		}
		return fmt.Sprintf(`
type %s %s
`,gt,cgt)
	}
}

func (t *Type) GoInterfaceDecl() string {
	ct := t.CType()
	gt := t.GoType()
	super := Super(ct)
	if super == "" {
		goInterfaces[gt] = true
		return fmt.Sprintf(`
type %s interface {
	Ptr() unsafe.Pointer
}
`,gt)
	}
	if IsGoInterface(super) {
		super = "Id"
	}
	return fmt.Sprintf(`
type %s struct { %s }
func (o %s) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) %s() %s {
	ret := %s{}
	ret.ptr = o.ptr
	return ret
}
`,gt,super,gt,gt,gt,gt)
}

func (t *Type) IsFunctionPtr() bool {
	if t == nil {
		return false
	}
	if td := t.Typedef(); td != nil {
		return td.IsFunctionPtr()
	}
	if pt := t.PointsTo(); pt != nil {
		return pt.IsFunction()
	}
	return false
}

func (t *Type) IsFunction() bool {
	if t == nil {
		fmt.Println("nil sent to IsFunction()")
		return false
	}
	if td := t.Typedef(); td != nil {
		return td.IsFunction()
	}
	return t.Node.IsFunction()
}

func (t *Type) IsValist() bool {
	if t == nil {
		return false
	}
	if t.GoType() == "__va_list_tag" { // OS dependent
		return true
	}
	return false
}

func (t *Type) ReturnType() *Type {
	if rt := t.Node.ReturnType(); rt != nil {
		return NewType(rt,t.Class)
	}
	return nil
}

func (t *Type) IsPointer() bool {
	if td := t.Typedef(); td != nil {
		return td.IsPointer()
	}
	return t.Node.IsPointer()
}

// cast C value to CGo
func (t *Type) CToGo(cval string) string {
	if t.IsPointer() {
		cval = "unsafe.Pointer(" + cval + ")"
	}
	return fmt.Sprintf("(%s)(%s)",t.GoType(),cval)
}

// Call a C function from Go with a given return type and parameter types
func GoToC(name string, pnames, snames []string, rtype *Type, ptypes []*Type, fun bool) string {
	if rtype == nil {
		fmt.Println("nil sent to GoToC")
		return ""
	}
	var ret strings.Builder
	rt := rtype.CType()
	rtgt := rtype.GoType()
	if IsGoInterface(rtgt) {
		rtgt = "Id"
	}
	sw := ShouldWrap(rtgt) || rtgt == "Id"
	if rt != "void" {
		if sw {
			ret.WriteString(fmt.Sprintf(
`ret := %s{}
	ret.ptr = `,rtgt))
		} else {
			if rtgt == "BOOL" {
				ret.WriteString("ret := (")
				rtgt = "bool"
			} else {
				ret.WriteString("ret := (" + rtgt + ")(")
			}
			if rtype.IsPointer() {
				ret.WriteString("unsafe.Pointer(")
			}
		}
	}
	ret.WriteString("C." + name + "(")
	parms := []string{}
	for i := 0; i < len(pnames); i++ {
		pn,pt := pnames[i],ptypes[i]
		p := pn
		if (ShouldWrap(pt.GoType()) || IsGoInterface(pt.GoType())) && !pt.Variadic {
			p = pn + ".Ptr()"
		} else {
			switch {
			case snames[i] != "":
				p = "unsafe.Pointer(&" + snames[i] + "[0])"
			case pt.Variadic:
				p = "unsafe.Pointer(&" + p + ")"
			case pt.IsPointer() && !fun:
				p = "unsafe.Pointer(" + pn + ")"
			default:
				p = "(" + pt.CGoType() + ")(" + pn + ")"
			}
		}
		parms = append(parms,p)
	}
	ret.WriteString(strings.Join(parms,", "))
	ret.WriteString(")")
	if rt != "void" && !sw {
		ret.WriteString(")")
		if rtype.IsPointer() {
			ret.WriteString(")")
		}
	}
	if rt == "BOOL" {
		ret.WriteString(" != 0")
	}
	for i,sname := range snames {
		if sname == "" { continue }
		ret.WriteString(fmt.Sprintf(`
	for i := 0; i < len(*%s); i++ {
		if %s[i] == nil {
			(*%s) = (*%s)[:i]
			break
		}
		(*%s)[i].ptr = %s[i]
	}`,pnames[i],sname,pnames[i],pnames[i],pnames[i],sname))
	}
	if rt != "void" {
		ret.WriteString(`
	return ret`)
	}
	return ret.String()
}

