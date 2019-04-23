package wrap

import (
	"fmt"
	//"reflect" 
	"strings"

	"gitlab.wow.st/gmp/clast/ast"
	"gitlab.wow.st/gmp/clast/types"
)

var (
	Debug = false
)

type cStruct struct {
	goName, cName string
}

type Wrapper struct {
	Interfaces map[string]Interface
	Types map[string]string
	gtMap map[string]string
	ctMap map[string]*types.Node 	 // map from go type as a string.
	goStructTypes map[string]cStruct // map from gotype to struct descr.
	goInterfaceTypes []string

	cCode strings.Builder	// put cGo code here
	goTypes strings.Builder // put Go type declarations here
	goCode strings.Builder	// put Go code here
	Processed map[string]bool
}

var gobuiltinTypes map[string]bool = map[string]bool{
	"byte": true,
	"int": true,
	"float64": true,
}

// translate C builtin types to CGo
var builtinTypes map[string]string = map[string]string{
	"char": "C.char",
	"signed char": "C.schar",
	"unsigned char": "C.uchar",
	"short": "C.short",
	"unsigned short": "C.ushort",
	"int": "C.int",
	"unsigned int": "C.uint",
	"long": "C.long",
	"unsigned long": "C.ulong",
	"long long": "C.longlong",
	"unsigned long long": "C.ulonglong",
	"float": "C.float",
	"double": "C.double",
	"complex float": "C.complexfloat",
	"complex double": "C.complexdouble",
}

//AddType registers a type that needs a Go wrapper struct
func (w *Wrapper) AddType(t, class string) {
	//fmt.Printf("Type: %s\n",t)
	if _,ok := builtinTypes[t]; ok {
		return
	}
	nt := w.goType(t,class)
	if Debug {
		fmt.Printf("AddType(): (%s) -> %s\n",t,nt)
	}
	if nt == "" {
		return
	}
	if _,ok := w.Interfaces[nt]; !ok { // not an interface
		return
	}
	w.Types[nt] = t
}

func (w *Wrapper) goType(t,class string) string {
	n,err := types.Parse(t)
	if err != nil {
		//fmt.Printf("Cannot parse type %s\n",t)
		return ""
	}
	if n.HasFunc() {
		//fmt.Printf("Function types not supported (%s)\n",t)
		return ""
	}
	ct := n.CtypeSimplified()
	ret := ""
	if nt, ok := w.gtMap[ct]; ok { // we already know this type
		return nt
	}
	if x := n.PointsTo(); x != nil {
		pt := x.CtypeSimplified()
//		if _,ok := w.Interfaces[pt]; ok {
//		// pointer to Objective-C interface, stop here
//			//fmt.Printf("goType(): %s -> %s\n",t,pt)
//			w.gtMap[ct] = pt
//			w.goInterfaceTypes = append(w.goInterfaceTypes,pt)
//			return pt
//		}
//		pt = x.BaseType().CtypeSimplified()
//		if _,ok := w.Interfaces[pt]; ok {
//		// ultimately points to an interface, so need a wrapper
//			w.AddType(pt,class)
//			ret = w.goType(pt,class) + "Ptr"
//			//fmt.Printf("goType(): %s -> %s\n",t,ret)
//			w.gtMap[ct] = ret
//			return ret
//		}
		// pointer to non-interface type
		ret = "*" + w.goType(pt,class)
		w.gtMap[ct] = ret
		w.ctMap[ret] = n
		return ret
	}
	if x := n.ArrayOf(); x != nil {
		pt := x.CtypeSimplified()
		w.AddType(pt,class)
		ret = w.goType(pt,class) + "Arr"
		//fmt.Printf("goType(): %s -> %s\n",t,ret)
		w.gtMap[ct] = ret
		w.ctMap[ret] = n
		return ret
	}
	if ct == "id" { ct = "NSObject" }
	if ct == "instancename" { ct = class }
	if bt,ok := builtinTypes[ct]; ok {
		ct = bt
	}
	if _,ok := w.Interfaces[ct]; ok {
		// pointer to Objective-C interface, stop here
		//fmt.Printf("goType(): %s -> %s\n",t,pt)
		w.gtMap[ct] = ct
		w.goInterfaceTypes = append(w.goInterfaceTypes,ct)
		w.ctMap[ct] = n
		return ct
	}
	if n.IsStruct() {
		gt := strings.Title(ct)
		gt = strings.ReplaceAll(gt, " ", "")
		w.gtMap[ct] = gt
		w.goStructTypes[gt] = cStruct{ goName: gt, cName: ct }
		//fmt.Printf("goType(): %s -> %s\n",t,gt)
		w.ctMap[gt] = n
		return gt
	}
	//fmt.Printf("goType(): %s -> %s\n",t,ct)
	w.gtMap[ct] = ct
	w.ctMap[ct] = n
	return ct
}

//FIXME: to be deleted
func cType(t, class string) string {
	n, err := types.Parse(t)
	if err != nil {
		//fmt.Printf("Cannot parse type %s\n",t)
		return "NOT IMPLEMENTED"
	}
	nt := n.CtypeSimplified()
	if nt == "id" {
		return "NSObject"
	}
	if nt == "instancetype" {
		return class + " *"
	}
	return nt
}

func NewWrapper(debug bool) *Wrapper {
	Debug = debug
	if Debug { fmt.Println("// Debug mode") }
	return &Wrapper{
		Interfaces: map[string]Interface{},
		Types: map[string]string{},
		gtMap: map[string]string{},
		ctMap: map[string]*types.Node{},
		goStructTypes: map[string]cStruct{},
		Processed: map[string]bool{},
	}
}

type Property struct {
	Name, Type, Attr string
}

type Parameter struct {
	Pname, Vname, Type string
}

type Method struct {
	Name, Type, Class string
	ClassMethod bool
	Parameters []Parameter
}

func (m Method) isVoid() bool {
	return m.Type == "void"
}

func (w Wrapper) isObject(tp string) bool { // takes a goType
	if _,ok := w.Interfaces[tp]; ok {
		return true
	}
	return false
}

func (w Wrapper) cparamlist(m Method) string {
	ret := make([]string,0)
	if !m.ClassMethod {
		ret = append(ret,"void* obj")
	}
	for _,p := range m.Parameters {
		tp := p.Type
		gtp := w.goType(tp,m.Class)
		//w.goType(tp,m.Class)
		if w.isObject(gtp) {
			tp = "void*"
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
		if first {
			ret = append(ret,m.Name + ":" + p.Vname)
			first = false
		} else {
			ret = append(ret, p.Pname + ":" + p.Vname)
		}
	}
	return strings.Join(ret," ")
}

//goreserved is a map telling whether a word is a go reserved word that is not
//also a C/Objective-C reserved word.
var goreserved map[string]bool = map[string]bool{
	"range": true,
}

func (w Wrapper) goparamlist(m Method) (string,[]string,bool) {
	ret := []string{}
	tps := []string{}
	for _,p := range m.Parameters {
		gname := p.Vname
		w.AddType(p.Type,m.Class)
		tp := w.goType(p.Type,m.Class)
		if tp == "" {
			return "UNSUPPORTED TYPE", []string{},false
		}
		if w.isObject(tp) {
			tp = "*" + tp
		}
		if goreserved[gname] {
			gname = gname + "_"
		}
		tps = append(tps,tp)
		ret = append(ret,fmt.Sprintf("%s %s",gname,tp))
	}
	return strings.Join(ret,", "),tps,true
}

func (w Wrapper) goparamnames(m Method) string {
	ret := make([]string,0)
	if !m.ClassMethod {
		ret = append(ret,"o.ptr")
	}
	for _,p := range m.Parameters {
		gname := p.Vname
		if goreserved[gname] {
			gname = gname + "_"
		}
		gt := w.goType(p.Type,m.Class)
		if gt == "" {
			return "UNSUPPORTED TYPE " + p.Type
		}
		if w.isObject(gt) {
			gname = gname + ".ptr"
		} else {
			n := w.ctMap[gt]
			star := ""
			for pt := n.PointsTo(); pt != nil; pt = pt.PointsTo() {
				star = star + "*"
			}
			ct := n.BaseType().CtypeSimplified()
			if n.IsStruct() {
				ct = strings.ReplaceAll(ct,"struct ","")
				ct = "struct_" + ct
			}
			if gt[0] == '*' { // wrap pointers in unsafe.Pointer()
				gname = "unsafe.Pointer(" + gname + ")"
			}
			gname = "(" + star + "C." + ct + ")(" + gname + ")"
		}
		ret = append(ret,gname)
	}
	return strings.Join(ret, ", ")
}

type Interface struct {
	Name, Super string
	Properties map[string]Property
	Methods map[string]Method
}

func (i Interface) IsRoot() bool {
	if i.Super == "" || i.Super == i.Name {
		return true
	}
	return false
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
		w.AddType(name,name)
	}
	var avail bool
	for _,c := range ns {
		switch x := c.(type) {
		case *ast.ObjCPropertyDecl:
			//fmt.Printf("ObjCPropertyDecl: %s\n",x.Name)
			p := Property{
				Name: x.Name,
				Type: x.Type,
			}
			//_,avail = w.GetParms(x,name) // TODO
			//if avail {
//				w.AddType(x.Type,name)
				i.Properties[p.Name] = p
			//}
		case *ast.ObjCMethodDecl:
			//fmt.Printf("ObjCMethodDecl: %s\n",x.Name)
			m := Method{
				Name: x.Name,
				Type: x.Type,
				Class: name,
				ClassMethod: x.ClassMethod,
			}
			m.Parameters, avail = w.GetParms(x,name)
			if avail {
//				w.AddType(x.Type,name)
				i.Methods[m.Name] = m
			}
		case *ast.ObjCProtocol:
			//fmt.Printf("ast.ObjCProtocol: %s\n",x.Name)
		case *ast.ObjCInterface:
			if x.Super {
				//fmt.Printf("ast.ObjCInterface: %s inherits from %s\n",name,x.Name)
				i.Super = x.Name
			}
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
				Type: x.Type,
			}
			ret = append(ret,p)
			j++
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

func (w *Wrapper) ProcessType(gotype string) {
	if gotype == "" {
		return
	}
	if gotype[0] == '*' {
		gotype = gotype[1:]
	}
	if _,ok := gobuiltinTypes[gotype]; ok {
		return
	}
	ctype := w.Types[gotype]
	if _,ok := builtinTypes[ctype]; ok {
		return
	}
	if Debug {
		fmt.Printf("Processing %s (%s)\n",gotype,ctype)
	}
	if w.Processed[gotype] {
		return
	}
	if i,ok := w.Interfaces[gotype]; ok {
		if Debug {
			fmt.Printf("Have interface for %s. super = %s\n",gotype,i.Super)
		}
		if i.Name != i.Super {
			w.ProcessType(i.Super)
		}
		var fields string
// if there is an Interface known for this type, decide if it is the
// root object and if so give it a pointer element:
		if i.IsRoot() {
			fields = "ptr unsafe.Pointer"
		} else {
			fields = i.Super // embed superclass
		}
		w.goTypes.WriteString(fmt.Sprintf(`
// %s
type %s struct { %s }
`,ctype, gotype, fields))
		w.Processed[gotype] = true
	}
	if s,ok := w.goStructTypes[gotype]; ok {
		ct := strings.ReplaceAll(s.cName," ","_")
		w.goTypes.WriteString(fmt.Sprintf(`
type %s %s
`,s.goName,"C." + ct))
		w.Processed[gotype] = true
	}
}

func (w *Wrapper) ProcessTypes(tps []string) {
	for _,tp := range tps {
		w.ProcessType(tp)
	}
}

func (w *Wrapper) Wrap(toproc []string) {

	w.cCode.WriteString(`/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#import <Foundation/Foundation.h>
`)

	pInterfaces := map[string]Interface{}
	for _,iface := range toproc {
		pInterfaces[iface] = w.Interfaces[iface]
	}
	for iname,i := range pInterfaces {
		if Debug {
			fmt.Printf("Interface %s: %d properties, %d methods\n",
			iname, len(i.Properties), len(i.Methods))
		}

		w.goCode.WriteString(fmt.Sprintf(`
func New%s() *%s {
	ret := &%s{}
	ret.ptr = unsafe.Pointer(C.New%s())
	ret = ret.Init()
	return ret
}
`,i.Name,i.Name,i.Name,i.Name))

		w.cCode.WriteString(fmt.Sprintf(`
%s*
New%s() {
	return [%s alloc];
}
`, i.Name, i.Name, i.Name))

		for _,p := range i.Properties {
			if Debug {
				fmt.Printf("  property: %s (%s)\n", p.Name, p.Type)
			}
			w.AddType(p.Type,i.Name)
		}
		for _,m := range i.Methods {
			if Debug {
				fmt.Printf("  method: %s (%s)\n", m.Name, m.Type)
			}
			gname := strings.Title(m.Name)

			grtype := ""
			grptr := ""
			w.AddType(m.Type,i.Name)
			cmtype := cType(m.Type,i.Name)
			if !m.isVoid() {
				grtype = w.goType(cmtype,i.Name)
				if grtype == "" {
					grtype = fmt.Sprintf("// goType(%s): NOT IMPLEMENTED\n",cmtype)
					continue
				}
			}
			gcast := ""
			gcast2 := ""
			if grtype != "" {
				if grtype[0] == '*' { // pointer return type
					if _,ok := w.Interfaces[grtype[1:]]; ok {
						//grptr = "*"
						gcast = "ret := &" + grtype[1:] + "{}\n	ret.ptr = unsafe.Pointer("
						gcast2 = ")\n	return ret"
					} else {
						gcast = "return (" + grtype + ")(unsafe.Pointer("
						gcast2 = "))"
					}
				} else {
					gcast = fmt.Sprintf("return (%s)(",grtype)
					gcast2 = ")"
				}
			}
			if grtype == "id" { // can't return id
				continue
			}
			w.ProcessType(grtype)
			gplist, gptypes, ok := w.goparamlist(m)
			if !ok {
				continue
			}
			w.ProcessTypes(gptypes)
			w.goCode.WriteString(fmt.Sprintf(`
func (o *%s) %s(%s) %s%s {
	%sC.%s_%s(%s)%s
}`,i.Name, gname, gplist, grptr, grtype, gcast, i.Name, m.Name, w.goparamnames(m),gcast2))

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
%s_%s(%s) {
	%s[%s %s];
}`, cmtype, i.Name, m.Name, w.cparamlist(m), cret, cobj, w.objcparamlist(m)))
		}
	}
	fmt.Println(`package main
`)
	fmt.Println(w.cCode.String())
	fmt.Println(`
*/
import "C"

import (
	"unsafe"
)
`)
	fmt.Println(w.goTypes.String())
	fmt.Println(w.goCode.String())
}
