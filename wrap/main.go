package wrap

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"gitlab.wow.st/gmp/nswrap/ast"
	"gitlab.wow.st/gmp/nswrap/types"
)

var (
	Debug = false
)

type Wrapper struct {
	Package string
	Interfaces map[string]*Interface
	Functions map[string]*Method
	NamedEnums map[string]*Enum
	AnonEnums []*Enum

	cgoFlags strings.Builder  // put cGo directives here
	cCode strings.Builder	  // put cGo code here
	goTypes strings.Builder   // put Go type declarations here
	goConst strings.Builder   // put Go constants (from C enums) here
	goCode strings.Builder    // put Go code here
	goHelpers strings.Builder // put Go helper functions here
	Processed map[string]bool
	VaArgs int
}

func NewWrapper(debug bool) *Wrapper {
	Debug = debug
	if Debug { fmt.Println("// Debug mode") }
	ret := &Wrapper{
		Interfaces: map[string]*Interface{},
		Functions: map[string]*Method{},
		NamedEnums: map[string]*Enum{},
		AnonEnums: []*Enum{},
		Processed: map[string]bool{},
		VaArgs: 16,
	}
	ret.cgoFlags.WriteString(fmt.Sprintf(`/*
#cgo CFLAGS: -x objective-c
`))
	ret.goTypes.WriteString(`
type Id struct { }
func (o *Id) Ptr() unsafe.Pointer { return unsafe.Pointer(o) }
`)
	return ret
}

func (w *Wrapper) Frameworks(ss []string) {
	if len(ss) == 0 {
		return
	}
	for _,s := range ss {
		w.cCode.WriteString(fmt.Sprintf("#import <%s/%s.h>\n",s,s))
	}
	w.cgoFlags.WriteString("#cgo LDFLAGS: -framework " + strings.Join(ss," -framework "))
}

func (w *Wrapper) Import(ss []string) {
	for _,s := range ss {
		w.cCode.WriteString("\n#import \"" + s + "\"\n")
	}
}

func (w *Wrapper) SysImport(ss []string) {
	for _,s := range ss {
		w.cCode.WriteString("\n#import <" + s + ">\n")
	}
}

func (w *Wrapper) Pragma(ss []string) {
	for _,s := range ss {
		w.cgoFlags.WriteString("\n#pragma " + s + "\n")
	}
}

type Property struct {
	Name, Attr string
	Type *types.Type
}

type Parameter struct {
	Pname, Vname string
	Type *types.Type
}

type Method struct {
	Name, Class, GoClass string
	Type *types.Type
	ClassMethod bool
	Parameters []*Parameter
}

type Enum struct {
	Name string
	Type *types.Type
	Constants []string
}

//isVoid() returns true if the method has no return value.
func (m Method) isVoid() bool {
	return m.Type.CType() == "void"
}

//hasFunctionParam() returns true if a method has a function as a parameter.
func (m Method) hasFunctionParam() bool {
	for _,p := range m.Parameters {
		if p.Type.IsFunction() || p.Type.IsFunctionPtr() {
			return true
		}
	}
	return false
}

func (w Wrapper) cparamlist(m *Method) (string,string) {
	ns := make([]string,0)
	ret := make([]string,0)
	if !m.ClassMethod {
		ret = append(ret,"void* o")
	}
	for _,p := range m.Parameters {
		var tp string
		if p.Type.IsPointer() || p.Type.Variadic {
			tp = "void*"
		} else {
			tp = p.Type.CType()
		}
		ns = append(ns,p.Vname)
		ret = append(ret,fmt.Sprintf("%s %s",tp,p.Vname))
	}
	return strings.Join(ns,", "),strings.Join(ret,", ")
}

func (w Wrapper) objcparamlist(m *Method) string {
	if len(m.Parameters) == 0 {
		return m.Name
	}
	first := true
	ret := []string{}
	for _,p := range m.Parameters {
		if first && !p.Type.Variadic {
			ret = append(ret,m.Name + ":" + p.Vname)
			first = false
		} else {
			if p.Type.Variadic {
				str := []string{m.Name + ":arr[0]"}
				for i := 1; i < w.VaArgs; i++ {
					str = append(str,"arr["+strconv.Itoa(i)+"]")
				}
				str = append(str,"nil")
				ret = append(ret, strings.Join(str,", "))
			} else {
				ret = append(ret, p.Pname + ":" + p.Vname)
			}
		}
	}
	return strings.Join(ret," ")
}

//goreserved is a map telling whether a word is a go reserved word that is not
//also a C/Objective-C reserved word.
var goreserved map[string]bool = map[string]bool{
	"range": true,
	"type": true,
}

func (w *Wrapper) gpntp(m *Method) ([]string,[]*types.Type,string) {
	w.processType(m.Type)
	ns := []string{}
	tps := []*types.Type{}
	if !m.ClassMethod {
		ns = append(ns,"o")
		tps = append(tps,types.NewTypeFromString(m.Class + "*",""))
	}
	for _,p := range m.Parameters {
		gname := p.Vname
		if goreserved[gname] {
			gname = gname + "_"
		}
		ns = append(ns,gname)
		tps = append(tps,p.Type)
	}
	w.processTypes(tps)
	ret := []string{}
	i := 0
	if !m.ClassMethod { i = 1 }
	for ; i < len(ns); i++ {
		gt := tps[i].GoType()
		if gt == "*Void" {
			gt = "unsafe.Pointer"
		}
		if tps[i].Variadic {
			gt = "..." + gt
			ns[i] = ns[i] + "s"
		}
		ret = append(ret,ns[i] + " " + gt)
	}
	return ns, tps, strings.Join(ret,", ")
}


type Interface struct {
	Name, GoName string
	Properties map[string]*Property
	Methods map[string]*Method
}

func (w *Wrapper) AddInterface(n *ast.ObjCInterfaceDecl) {
	//fmt.Printf("ast.ObjCInterfaceDecl: %s\n",n.Name)
	w.add(n.Name, n.Children())
}

func (w *Wrapper) AddCategory(n *ast.ObjCCategoryDecl) {
	ns := n.Children()
	if len(ns) > 0 {
		switch x := ns[0].(type) {
		case *ast.ObjCInterface:
			w.add(x.Name, ns[1:])
			return
		}
	}
	fmt.Printf("Not adding methods for %s: interface name not found in first child node of category defclaration\n",n.Name)
}

func (w *Wrapper) AddFunction(n *ast.FunctionDecl) {
	//treat functions as class methods with no class
	tp := types.NewTypeFromString(n.Type,"")
	m := &Method{
		Name: n.Name,
		Type: tp.ReturnType(),
		Class: "",
		ClassMethod: true,
		Parameters: []*Parameter{},
	}
	f := tp.Node.Children[len(tp.Node.Children)-1] // Function node
	if f.Kind != "Function" {
		//fmt.Printf("AddFunction(%s): not a function -- Node type is %s\n%s",n.Name,f.Kind,tp.String())
		return
	}
	//fmt.Printf("FunctionDecl: %s (%s) %s\n",n.Type,m.Type.CType(),n.Name)
	i := 0
	for _,c := range n.Children() {
		switch x := c.(type) {
		case *ast.ParmVarDecl:
			p := &Parameter{
				Vname: x.Name,
				Type: types.NewTypeFromString(x.Type,""),
			}
			m.Parameters = append(m.Parameters,p)
			i++
			//fmt.Printf("  %s\n",p.Type.CType())
		}
	}
	if i > 0 && len(f.Children) > i {
		if e := f.Children[i]; len(e.Children) > 0 {
			//fmt.Println("  Next parameter: ",e.Children[0].String())
			m.Parameters[i-1].Type.Variadic = true
		}
	}
	w.Functions[n.Name] = m
}

//FIXME: copied from nswrap/main.go, should put this in a utils package
func matches(x string, rs []string) bool {
	for _,r := range rs {
		if m,_ := regexp.MatchString("^" + r + "$",x); m {
			return true
		}
	}
	return false
}

func (w *Wrapper) AddEnum(n *ast.EnumDecl,rs []string) {
	if n.Name != "" && !matches(n.Name,rs) {
		return
	}
	//fmt.Printf("Adding enum: (%s) %s\n",n.Type,n.Name)
	var tp *types.Type
	a := (*Avail)(&[]AvailAttr{})
	if n.Type == "" {
		tp = nil
	} else {
		tp = types.NewTypeFromString(n.Type,"")
		//fmt.Printf("  type: %s -> %s\n",n.Type,tp.CType())
	}
	e := &Enum{
		Name: n.Name, // NOTE: may be empty string
		Type: tp,
		Constants: []string{},
	}
	for _,c := range n.Children() {
		switch x := c.(type) {
		case *ast.AvailabilityAttr, *ast.UnavailableAttr:
			a.Add(x)
		case *ast.EnumConstantDecl:
			//fmt.Printf("*ast.EnumConstantDecl: (%s) '%s': %s\n",n.Type,n.Name,x.Name)
			if n.Name == "" && !matches(x.Name,rs) {
				continue
			}
			e.Constants = append(e.Constants,x.Name)
		}
	}
	if a.Available() && len(e.Constants) > 0 {
		if e.Name == "" {
			w.AnonEnums = append(w.AnonEnums,e)
		} else {
			w.NamedEnums[e.Name] = e
		}
		//fmt.Printf("  added\n")
	}
}

func (w *Wrapper) add(name string, ns []ast.Node) {
	var i *Interface
	var ok bool
	goname := types.NewTypeFromString(name,name).GoType()
	types.Wrap(goname)
	if i,ok = w.Interfaces[name]; !ok {
		i = &Interface{
			Name: name,
			GoName: goname,
			Properties: map[string]*Property{},
			Methods: map[string]*Method{},
		}
	}
	var avail bool
	for _,c := range ns {
		switch x := c.(type) {
		case *ast.ObjCPropertyDecl:
			//fmt.Printf("ObjCPropertyDecl: %s\n",x.Name)
			p := &Property{
				Name: x.Name,
				Type: types.NewTypeFromString(x.Type,name),
			}
			//_,avail = w.GetParms(x,name) // TODO
			//if avail {
				i.Properties[p.Name] = p
			//}
		case *ast.ObjCMethodDecl:
			//fmt.Printf("ObjCMethodDecl: %s (%s) %s\n",x.Type,name,x.Name)
			m := &Method{
				Name: x.Name,
				Type: types.NewTypeFromString(x.Type,name),
				Class: name,
				GoClass: goname,
				ClassMethod: x.ClassMethod,
			}
			//fmt.Println(m.Type.Node.String())
			m.Parameters, avail = w.GetParms(x,name)
			if avail {
				i.Methods[m.Name] = m
			}
		case *ast.ObjCProtocol:
			//fmt.Printf("ast.ObjCProtocol: %s\n",x.Name)
		case *ast.ObjCInterface:
			if x.Super {
				//fmt.Printf("ast.ObjCInterface: %s inherits from %s\n",name,x.Name)
				types.SetSuper(name,x.Name)
			}
		case *ast.ObjCTypeParamDecl:
			//fmt.Printf("ObjCTypeParamDecl: %s = %s\n",x.Name,x.Type)
			types.SetTypeParam(name,x.Name,x.Type)
		case *ast.Unknown:
			//fmt.Printf("(*ast.Unkonwn %s: %s)\n",x.Name,x.Content)
		default:
			//fmt.Println(reflect.TypeOf(x))
		}
	}
	//Add class methods from super class
	var supmethods func(*Interface,string)
	supmethods = func(i *Interface, s string) {
		if sup,ok := w.Interfaces[s]; !ok {
			return
		} else {
			for _,m := range sup.Methods {
				if !m.ClassMethod {
					continue
				}
				m2 := &Method{
					Name: m.Name,
					Class: i.Name,
					GoClass: i.GoName,
					Type: m.Type.CloneToClass(i.Name),
					ClassMethod: true,
					Parameters: []*Parameter{},
				}
				for _,p := range m.Parameters {
					m2.Parameters = append(m2.Parameters,
					&Parameter{
						Pname: p.Pname,
						Vname: p.Vname,
						Type: p.Type.CloneToClass(i.Name),
					})
				}
				i.Methods[m.Name] = m2
			}
		}
		supmethods(i,types.Super(s))
	}
	supmethods(i,types.Super(i.Name))
	//fmt.Println("Add interface ",i.Name)
	w.Interfaces[i.Name] = i
}

type AvailAttr struct {
	OS, Version string
	Deprecated bool
}

type Avail []AvailAttr

func (a *Avail) Add(n ast.Node) {
	switch x := n.(type) {
	case *ast.AvailabilityAttr:
		*a = append(*a, AvailAttr{
			OS: x.OS,
			Version: x.Version,
			Deprecated: (x.Unknown1 != "0") || x.IsUnavailable,
		})
	case *ast.UnavailableAttr:
		*a = append(*a, AvailAttr{
			OS: "macos", Deprecated: true })
	}
}

func (a *Avail) Available() bool {
	if len(*a) == 0 {
		return true
	}
	for _,x := range *a {
		if x.OS == "macos" && x.Deprecated == false {
			return true
		}
	}
	return false
}

//GetParms returns the parameters of a method declaration and a bool
//indicating whether the given method is available on MacOS and not
//deprecated.
func (w *Wrapper) GetParms(n ast.Node,class string) ([]*Parameter,bool) {
	ret := make([]*Parameter,0)
	avail := (*Avail)(&[]AvailAttr{})
	var parms []string
	switch x := n.(type) {
	case *ast.ObjCMethodDecl:
		parms = x.Parameters
	case *ast.FunctionDecl:
	default:
		panic("GetParms called with wrong node type")
	}
	j := 0
	for _,c := range n.Children() {
		switch x := c.(type) {
		case *ast.ParmVarDecl:
			p := &Parameter{
				Vname: x.Name,
				Type: types.NewTypeFromString(x.Type,class),
			}
			if parms != nil {
				p.Pname = parms[j]
			}
			ret = append(ret,p)
			j++
		case *ast.Variadic:
			ret[j-1].Type.Variadic = true
		case *ast.AvailabilityAttr, *ast.UnavailableAttr:
			avail.Add(x)
		case *ast.Unknown:
			if Debug { fmt.Printf("GetParms(): ast.Unknown: %s\n",x.Name) }
		}
	}
	// check that the method is available for this OS and not deprecated
	if !avail.Available() {
		return nil, false
	}
	// check that we found the right number of parameters
	//if len(ret) != len(n.Parameters) {
	//	fmt.Printf("Error in method declaration %s: Wrong number of ParmVarDecl children: %d parameters but %d ParmVarDecl children\n",n.Name,len(n.Parameters),len(ret))
	//}
	return ret, true
}

func (w *Wrapper) processTypes(tps []*types.Type) {
	for _,tp := range tps {
		w.processType(tp)
	}
}

func (w *Wrapper) processType(tp *types.Type) {
	bt := tp.BaseType()
	gt := bt.GoType()
	if gt == "" {
		return
	}
	if gt[0] == '*' {
		w.processType(bt.PointsTo())
		return
	}
	if w.Processed[gt] { return }
	w.Processed[gt] = true
	if gt == "Char" {
		w.CharHelpers()
	}
	if gt == "NSEnumerator" {
		w.EnumeratorHelpers()
	}
	if bt.IsFunction() || bt.IsFunctionPtr() {
		return
	}
	super := types.Super(gt)
	if super != "" {
		tp := types.NewTypeFromString(super,"")
		types.Wrap(tp.GoType())
		w.processType(tp)
	}
	w.goTypes.WriteString(bt.GoTypeDecl())
}

func (w *Wrapper) CharHelpers() {
	w.goHelpers.WriteString(`
func CharWithGoString(s string) *Char {
	return (*Char)(unsafe.Pointer(C.CString(s)))
}

func CharFromBytes(b []byte) *Char {
	return (*Char)(unsafe.Pointer(C.CString(string(b))))
}

func (c *Char) String() string {
	return C.GoString((*C.char)(c))
}
`)
}

func (w *Wrapper) EnumeratorHelpers() {
	w.goHelpers.WriteString(`
func (e *NSEnumerator) ForIn(f func(*Id) bool) {
	for o := e.NextObject(); o != nil; o = e.NextObject() {
		if !f(o) { break }
	}
}
`)
}

func (w *Wrapper) ProcessMethod(m *Method) {
	w._processMethod(m,false)
}

func (w *Wrapper) ProcessFunction(m *Method) {
	if m.Type.Node.IsId() {
	//do not wrap functions that return ID because of CGo struct size bug
		return
	}
	w._processMethod(m,true)
}

func (w *Wrapper) _processMethod(m *Method,fun bool) {
	if Debug {
		fmt.Printf("  method: %s (%s)\n", m.Name, m.Type)
	}
	if m.Type.IsFunction() || m.Type.IsFunctionPtr() || m.hasFunctionParam() {
		return
	}
	gname := strings.Title(m.Name)
	switch {
	case !m.ClassMethod:
	case m.Type.GoType() != "*" + m.GoClass:
		gname = m.GoClass + gname
	default:
		//Shorten class method names
		lens1 := len(m.Class)
		i := 0
		if len(gname) < len(m.Class) { i = lens1 - len(gname) }
		for ; i < lens1; i++ {
			if m.Class[i:] == gname[:lens1 - i] { break }
		}
		if lens1 - i >= len(gname) {
			gname = m.GoClass + gname
		} else {
			gname = m.GoClass + gname[lens1-i:]
		}
	}
	receiver := ""
	if !m.ClassMethod {
		receiver = "(o *" + m.GoClass + ") "
	}
	cname := m.Name
	if m.Class != "" {
		cname = m.Class + "_" + cname
	}
	var cmtype string
	if m.Type.IsPointer() {
		// work around cgo bugs with struct size calculation
		cmtype = "void*"
		if x := m.Type.Node.Qualifiers(); x != "" {
			cmtype = x + " " + cmtype
		}
		if x := m.Type.Node.Annotations(); x != "" {
			cmtype = cmtype + " " + x
		}
	} else {
		cmtype = m.Type.CTypeAttrib()
	}
	ns,tps,gplist := w.gpntp(m)
	grtype := m.Type.GoType()
	if grtype == "Void" {
		grtype = ""
	}
	if types.IsGoInterface(grtype) {
		grtype = "*Id"
	}
	if grtype == "BOOL" { // convert objective-c bools to Go bools
		grtype = "bool"
	}
	if gname == grtype { // avoid name conflicts between methods and types
		gname = "Get" + gname
	}
	w.goCode.WriteString(fmt.Sprintf(`
//%s
func %s%s(%s) %s {
`,m.Type.CType(),receiver,gname,gplist,grtype))
	lparm := len(tps)-1
	if len(tps) > 0 && tps[lparm].Variadic {
		vn := ns[lparm]
		vn = vn[:len(vn)-1]
		ns[lparm] = vn
		w.goCode.WriteString(fmt.Sprintf(
`	var %s [%d]unsafe.Pointer
	for i,o := range %ss {
		%s[i] = o.Ptr()
	}
`,vn,w.VaArgs,vn,vn))
	}
	w.goCode.WriteString(`	` +
		types.GoToC(cname,ns,m.Type,tps) + "\n}\n\n")

	cret := ""
	if !m.isVoid() {
		cret = "return "
	}
	var cobj string
	if m.ClassMethod {
		cobj = m.Class
	} else {
		cobj = "(" + m.Class + "*)o"
	}
	cns,cntps := w.cparamlist(m)
	_ = cns
	if fun {
		return
	}
	w.cCode.WriteString(fmt.Sprintf(`
%s
%s(%s) {
`, cmtype, cname, cntps))
	if len(tps) > 0 && tps[lparm].Variadic {
		w.cCode.WriteString(fmt.Sprintf(
`	%s* arr = %s;
`, tps[lparm].CType(), ns[lparm]))
	}
	if fun {
		w.cCode.WriteString(fmt.Sprintf(`	%s%s(%s);
}`, cret, m.Name, cns))
	} else {
		w.cCode.WriteString(fmt.Sprintf(`	%s[%s %s];
}`, cret, cobj, w.objcparamlist(m)))
	}

	// create GoString helper method
	if ok,_ := regexp.MatchString("WithString$",m.Name); ok {
		//fmt.Printf("--%s\n",gname)
		gname2 := gname[:len(gname)-6] + "GoString"
		gps := []string{}
		i := 0
		if !m.ClassMethod { i = 1 }
		for ; i < len(ns); i++ {
			gt := tps[i].GoType()
			//fmt.Printf("  %s\n",gt)
			ns2 := ns[i]
			if gt == "*NSString" {
				gt = "string"
				ns[i] = gStringToNsstring(ns[i])
			}
			gps = append(gps,ns2 + " " + gt)
		}
		gplist = strings.Join(gps,", ")
		obj := ""
		if !m.ClassMethod {
			obj = "o."
			ns = ns[1:]
		}
		w.goCode.WriteString(fmt.Sprintf(`
func %s%s(%s) %s {
	return %s%s(%s)
}
`,receiver,gname2,gplist,grtype,obj,gname,strings.Join(ns,", ")))
	}
}

func gStringToNsstring(s string) string {
	return fmt.Sprintf("NSStringWithUTF8String(CharWithGoString(%s))",s)
}


func (w *Wrapper) ProcessEnum(e *Enum) {
	//fmt.Printf("Processing enum (%s)\n",e.Name)
	gtp := ""
	ctp := ""
	if e.Name != "" {
		gtp = e.Name
		ctp = "C." + e.Name
	} else {
		gtp = e.Type.GoType()
		ctp = e.Type.CGoType()
	}
	if e.Type != nil {
		if !w.Processed[gtp] {
			w.goTypes.WriteString(fmt.Sprintf(`
type %s %s
`,gtp,ctp))
			w.Processed[gtp] = true
		}
	}
	gtp = gtp + " "
	//fmt.Printf("  gtp = %s; ctp = %s\n",gtp,ctp)
	for _,c := range e.Constants {
		w.goConst.WriteString(fmt.Sprintf(`
const %s %s= C.%s
`,c,gtp,c))
	}
}

func (w *Wrapper) Wrap(toproc []string) {
	if w.Package == "" { w.Package = "ns" }
	err := os.MkdirAll(w.Package,0755)
	if err != nil {
		fmt.Printf("Error creating directory '%s'\n%s\n",w.Package,err)
		os.Exit(-1)
	}
	of,err := os.Create(path.Join(w.Package,"main.go"))
	if err != nil {
		fmt.Printf("Error opening file %s\n%s\n",path.Join(w.Package,"main.go"),err)
		os.Exit(-1)
	}
	fmt.Printf("Writing output to %s\n",path.Join(w.Package,"main.go"))
	pInterfaces := map[string]*Interface{}
	for _,iface := range toproc {
		pInterfaces[iface] = w.Interfaces[iface]
	}
	//FIXME: sort pInterfaces
	for _,i := range pInterfaces {
		if i == nil {
			continue
		}
		fmt.Printf("Interface %s: %d properties, %d methods\n",
			i.Name, len(i.Properties), len(i.Methods))

		w.goCode.WriteString(fmt.Sprintf(`
func %sAlloc() *%s {
	return (*%s)(unsafe.Pointer(C.%sAlloc()))
}
`,i.GoName,i.GoName,i.GoName,i.Name))

		w.cCode.WriteString(fmt.Sprintf(`
void*
%sAlloc() {
	return [%s alloc];
}
`, i.Name, i.Name))

		//FIXME: sort properties
		for _,p := range i.Properties {
			if Debug {
				fmt.Printf("  property: %s (%s)\n", p.Name, p.Type.CType())
			}
		}
		//FIXME: sort methods
		for _,m := range i.Methods {
			w.ProcessMethod(m)

		}
	}
	for _,m := range w.Functions {
		//fmt.Printf("Processing function %s %s\n",m.Type.CType(),m.Name)
		w.ProcessFunction(m)
	}
	for _,e := range w.NamedEnums {
		w.ProcessEnum(e)
	}
	for _,e := range w.AnonEnums {
		w.ProcessEnum(e)
	}
	fmt.Printf("%d functions\n", len(w.Functions))
	fmt.Printf("%d enums\n", len(w.NamedEnums) + len(w.AnonEnums))
	of.WriteString("package " + w.Package + "\n\n")
	of.WriteString(w.cgoFlags.String() + "\n")
	of.WriteString(w.cCode.String())
	of.WriteString(`
*/
import "C"

import (
	"unsafe"
)
`)
	of.WriteString(w.goTypes.String())
	of.WriteString(w.goConst.String())
	of.WriteString(w.goHelpers.String())
	of.WriteString(w.goCode.String())
	of.Close()
}
