package wrap

import (
	"fmt"
	"os"
	"path"
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
	Interfaces map[string]Interface

	cCode strings.Builder	  // put cGo code here
	goTypes strings.Builder   // put Go type declarations here
	goCode strings.Builder    // put Go code here
	goHelpers strings.Builder // put Go helper functions here
	Processed map[string]bool
	VaArgs int
}

func NewWrapper(debug bool) *Wrapper {
	Debug = debug
	if Debug { fmt.Println("// Debug mode") }
	ret := &Wrapper{
		Interfaces: map[string]Interface{},
		Processed: map[string]bool{},
		VaArgs: 16,
	}
	ret.cCode.WriteString(`/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
`)
	ret.goTypes.WriteString(`
type Id struct { ptr unsafe.Pointer }
func (o *Id) Ptr() unsafe.Pointer { return o.ptr }
`)
	return ret
}

func (w *Wrapper) Import(ss []string) {
	for _,s := range ss {
		w.cCode.WriteString(`
#import "` + s + `"
`)
	}
}

func (w *Wrapper) SysImport(ss []string) {
	for _,s := range ss {
		w.cCode.WriteString("\n#import <" + s + ">\n")
	}
}

func (w *Wrapper) Pragma(ss []string) {
	for _,s := range ss {
		w.cCode.WriteString("\n#pragma " + s + "\n")
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
	Name, Class string
	Type *types.Type
	ClassMethod bool
	Parameters []Parameter
}

//isVoid() returns true if the method has no return value.
func (m Method) isVoid() bool {
	return m.Type.CType() == "void"
}

//hasFunctionParam() returns true if a method has a function as a parameter.
func (m Method) hasFunctionParam() bool {
	for _,p := range m.Parameters {
		if p.Type.IsFunction() {
			return true
		}
	}
	return false
}

func (w Wrapper) cparamlist(m Method) string {
	ret := make([]string,0)
	if !m.ClassMethod {
		ret = append(ret,"void* obj")
	}
	for _,p := range m.Parameters {
		var tp string
		if p.Type.Node.IsPointer() || p.Type.Variadic {
			tp = "void*"
		} else {
			tp = p.Type.CType()
		}
		ret = append(ret,fmt.Sprintf("%s %s",tp,p.Vname))
	}
	return strings.Join(ret,", ")
}

func (w Wrapper) objcparamlist(m Method) string {
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
}

func (m *Method) gpntp() ([]string,[]*types.Type,string) {
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
	Name string
	Properties map[string]Property
	Methods map[string]Method
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

func (w *Wrapper) add(name string, ns []ast.Node) {
	var i Interface
	var ok bool
	if i,ok = w.Interfaces[name]; !ok {
		i = Interface{
			Name: name,
			Properties: map[string]Property{},
			Methods: map[string]Method{},
		}
	}
	tp := types.NewTypeFromString(name,name)
	types.Wrap(tp.GoType())
	var avail bool
	for _,c := range ns {
		switch x := c.(type) {
		case *ast.ObjCPropertyDecl:
			//fmt.Printf("ObjCPropertyDecl: %s\n",x.Name)
			p := Property{
				Name: x.Name,
				Type: types.NewTypeFromString(x.Type,name),
			}
			//_,avail = w.GetParms(x,name) // TODO
			//if avail {
				i.Properties[p.Name] = p
			//}
		case *ast.ObjCMethodDecl:
			//fmt.Printf("ObjCMethodDecl: %s\n",x.Name)
			m := Method{
				Name: x.Name,
				Type: types.NewTypeFromString(x.Type,name),
				Class: name,
				ClassMethod: x.ClassMethod,
			}
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
	//fmt.Println("Add interface ",i.Name)
	w.Interfaces[i.Name] = i
}

type AvailAttr struct {
	OS, Version string
	Deprecated bool
}

//GetParms returns the parameters of a method declaration and a bool
//indicating whether the given method is available on MacOS and not
//deprecated.
func (w *Wrapper) GetParms(n *ast.ObjCMethodDecl,class string) ([]Parameter,bool) {
	ret := make([]Parameter,0)
	avail := make([]AvailAttr,0)
	j := 0
	for _,c := range n.Children() {
		switch x := c.(type) {
		case *ast.ParmVarDecl:
			p := Parameter{
				Pname: n.Parameters[j],
				Vname: x.Name,
				Type: types.NewTypeFromString(x.Type,class),
			}
			ret = append(ret,p)
			j++
		case *ast.Variadic:
			ret[j-1].Type.Variadic = true
		case *ast.AvailabilityAttr:
			avail = append(avail,
				AvailAttr{
					OS: x.OS,
					Version: x.Version,
					Deprecated: x.Unknown1 != "0",
				})
			//fmt.Println("AvailAttr ",avail,x)
		case *ast.UnavailableAttr:
			avail = append(avail,
				AvailAttr{ OS: "macos", Deprecated: true })
		case *ast.Unknown:
			if Debug { fmt.Printf("GetParms(): ast.Unknown: %s\n",x.Name) }
		}
	}
	// check that the method is available for this OS and not deprecated
	a := func() bool {
		if len(avail) == 0 {
			return true
		}
		for _,x := range avail {
			if x.OS == "macos" && x.Deprecated == false {
				return true
			}
		}
		return false
	}()
	if !a {
		return nil, false
	}
	// check that we found the right number of parameters
	if len(ret) != len(n.Parameters) {
		fmt.Printf("Error in method declaration %s: Wrong number of ParmVarDecl children: %d parameters but %d ParmVarDecl children\n",n.Name,len(n.Parameters),len(ret))
	}
	return ret, true
}

func (w *Wrapper) processTypes(tps []*types.Type) {
	switch len(tps) {
	case 0:
		return
	case 1:
		w.processType(tps[0])
	default:
		for _,tp := range tps {
			w.processType(tp)
		}
		return
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
	if bt.IsFunction() {
		return
	}
	super := types.Super(gt)
	if super != "" {
		types.Wrap(super)
		w.processType(types.NewTypeFromString(super,""))
	}
	w.goTypes.WriteString(bt.GoTypeDecl())
}

func (w *Wrapper) CharHelpers() {
	w.goHelpers.WriteString(`
func CharFromString(s string) *Char {
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
	pInterfaces := map[string]Interface{}
	for _,iface := range toproc {
		pInterfaces[iface] = w.Interfaces[iface]
	}
	//FIXME: sort pInterfaces
	for _,i := range pInterfaces {
		fmt.Printf("Interface %s: %d properties, %d methods\n",
			i.Name, len(i.Properties), len(i.Methods))

		w.goCode.WriteString(fmt.Sprintf(`
func New%s() *%s {
	ret := &%s{}
	ret.ptr = unsafe.Pointer(C.New%s())
	//ret = ret.Init()
	return ret
}
`,i.Name,i.Name,i.Name,i.Name))

		w.cCode.WriteString(fmt.Sprintf(`
%s*
New%s() {
	return [%s alloc];
}
`, i.Name, i.Name, i.Name))

		//FIXME: sort properties
		for _,p := range i.Properties {
			if Debug {
				fmt.Printf("  property: %s (%s)\n", p.Name, p.Type.CType())
			}
		}
		//FIXME: sort methods
		for _,m := range i.Methods {
			if Debug {
				fmt.Printf("  method: %s (%s)\n", m.Name, m.Type)
			}
			if m.Type.IsFunction() {
				continue
			}
			if m.hasFunctionParam() {
				continue
			}
			gname := strings.Title(m.Name)
			if !m.ClassMethod {
				gname = "(o *" + i.Name + ") " + gname
			}
			cname := i.Name + "_" + m.Name

			cmtype := m.Type.CTypeAttrib()
			ns,tps,gplist := m.gpntp()
			w.processTypes(tps)
			w.processType(m.Type)
			grtype := m.Type.GoType()
			if grtype == "Void" {
				grtype = ""
			}
			w.goCode.WriteString(fmt.Sprintf(`
func %s(%s) %s {
`,gname,gplist,grtype))
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
			w.goCode.WriteString(
				types.GoToC(cname,ns,m.Type,tps) + "}\n\n")

			cret := ""
			if !m.isVoid() {
				cret = "return "
			}
			var cobj string
			if m.ClassMethod {
				cobj = i.Name
			} else {
				cobj = "(id)obj"
			}
			w.cCode.WriteString(fmt.Sprintf(`
%s
%s(%s) {
`, cmtype, cname, w.cparamlist(m)))
			if len(tps) > 0 && tps[lparm].Variadic {
				w.cCode.WriteString(fmt.Sprintf(
`	%s* arr = %s;
`, tps[lparm].CType(), ns[lparm]))
			}
			w.cCode.WriteString(fmt.Sprintf(`	%s[%s %s];
}`, cret, cobj, w.objcparamlist(m)))
		}
	}
	of.WriteString("package " + w.Package + "\n\n")
	of.WriteString(w.cCode.String())
	of.WriteString(`
*/
import "C"

import (
	"unsafe"
)
`)
	of.WriteString(w.goTypes.String())
	of.WriteString(w.goHelpers.String())
	of.WriteString(w.goCode.String())
	of.Close()
}
