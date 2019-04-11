package wrap

import (
	"fmt"
	//"reflect" 
	"strings"

	"gitlab.wow.st/gmp/clast/ast"
)

var (
	Debug = false
)

type Wrapper struct {
	Interfaces map[string]Interface
	Types map[string]string
	cCode strings.Builder // put cGo code here
	goTypes strings.Builder // put Go type declarations here
	goCode strings.Builder // put Go code here
	Processed map[string]bool
}

// translate C builtin types to CGo
var builtinTypes map[string]string = map[string]string{
	"char": "byte",
	"signed char": "byte",
	"unsigned char": "byte",
	"short": "int",
	"unsigned short": "int",
	"int": "int",
	"unsigned int": "int",
	"long": "int",
	"unsigned long": "int",
	"long long": "int",
	"unsigned long long": "int",
	"float": "float64",
	"double": "float64",
	"complex float": "C.complexfloat",
	"complex double": "C.complexdouble",
}

var gobuiltinTypes map[string]bool = map[string]bool{
	"byte": true,
	"int": true,
	"float64": true,
}

/*
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
}*/

func (w *Wrapper) AddType(t,class string) {
	if _,ok := builtinTypes[t]; ok {
		return
	}
	nt, err := goType(t,class)
	if err != nil {
		return
	}
	w.Types[nt] = t
}

func goType(t,class string) (string, error) {
	// strip off pointers, < > and blank space
	nt := strings.ReplaceAll(t,"*","")
	if len(nt) > 3 && nt[0:3] == "id<" {
		nt = t[3:len(t)-1]
	}
	nt = strings.ReplaceAll(nt,"const ","")
	nt = strings.ReplaceAll(nt," _Nullable","")
	nt = strings.ReplaceAll(nt," _Nonnull","")
	if t == "instancetype" {
		return class, nil
	}
	nt = strings.ReplaceAll(nt,"<","__")
	nt = strings.ReplaceAll(nt,">","__")
	nt = strings.ReplaceAll(nt,",","_")
	if x,ok := builtinTypes[nt]; ok { // do not add builtin types
		return x, nil
	}
	if len(nt) > 5 && nt[0:5] == "enum " { // FIXME: deal with enums?
		return "", fmt.Errorf("goType(): enum")
	}
	nt = strings.ReplaceAll(nt," ","")
	if nt == "void" {
		return "", fmt.Errorf("goType(): void")
	}
	if nt[len(nt)-1] == ')' { // skip function pointers
		return "", fmt.Errorf("goType(): function pointer")
	}
	return nt, nil
}

func cType(t, class string) string {
	nt := strings.ReplaceAll(t," _Nullable","")
	nt = strings.ReplaceAll(nt," _Nonnull","")
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
	}
}

type Property struct {
	Name, Type, Type2, Attr string
}

type Parameter struct {
	Pname, Vname, Type, Type2 string
}

type Method struct {
	Name, Type, Type2, Class string
	Parameters []Parameter
}

func (m Method) isVoid() bool {
	return typeOrType2(m.Type,m.Type2) == "void"
}

func (w Wrapper) isObject(tp string) bool { // takes a goType
	if _,ok := w.Interfaces[tp]; ok {
		return true
	}
	return false
}

func (w Wrapper) cparamlist(m Method) string {
	ret := []string{"void* obj"}
	for _,p := range m.Parameters {
		tp := typeOrType2(p.Type,p.Type2)
		gtp,_ := goType(tp,m.Class)
		if w.isObject(gtp) {
			tp = "void*"
		}
		ret = append(ret,fmt.Sprintf("%s %s",tp,p.Vname))
	}
	return strings.Join(ret,", ")
}

func (m Method) objcparamlist() string {
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

var goreserved map[string]bool = map[string]bool{
	"range": true,
}

func (w Wrapper) goparamlist(m Method) (string,[]string,bool) {
	ret := []string{}
	tps := []string{}
	for _,p := range m.Parameters {
		gname := p.Vname
		tp,err := goType(typeOrType2(p.Type,p.Type2),m.Class)
		if err != nil {
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
	ret := []string{"o.ptr"}
	for _,p := range m.Parameters {
		gname := p.Vname
		if goreserved[gname] {
			gname = gname + "_"
		}
		gt,_ := goType(typeOrType2(p.Type,p.Type2),m.Class)
		if w.isObject(gt) {
			gname = gname + ".ptr"
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
	switch x := ns[0].(type) {
	case *ast.ObjCInterface:
		w.add(x.Name, ns[1:])
	default:
		fmt.Printf("Not adding methods for %s: interface name not found in first child node of category defclaration\n",n.Name)
	}
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
				Type2: x.Type2,
			}
			//_,avail = w.GetParms(x,name) // TODO
			//if avail {
				w.AddType(typeOrType2(x.Type,x.Type2),name)
				i.Properties[p.Name] = p
			//}
		case *ast.ObjCMethodDecl:
			//fmt.Printf("ObjCMethodDecl: %s\n",x.Name)
			m := Method{
				Name: x.Name,
				Type: x.Type,
				Type2: x.Type2,
				Class: name,
			}
			m.Parameters, avail = w.GetParms(x,name)
			if avail {
				w.AddType(typeOrType2(x.Type,x.Type2),name)
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
	w.Interfaces[i.Name] = i
}

type AvailAttr struct {
	OS, Version string
	Deprecated bool
}

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
				Type2: x.Type2,
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
	if len(ret) != len(n.Parameters) {
		fmt.Printf("Error in method declaration %s: Wrong number of ParmVarDecl children: %d parameters but %d ParmVarDecl children\n",n.Name,len(n.Parameters),len(ret))
	}
	return ret, true
}

func typeOrType2(t1, t2 string) string {
	if t2 != ""  && t2 != "id" {
		return t2
	}
	return t1
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
	}
	var fields string
// if there is an Interface known for this type, decide if it is the
// root object and if so give it a pointer element:
	if i,ok := w.Interfaces[gotype]; ok && i.IsRoot() {
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

func (w *Wrapper) ProcessTypes(tps []string) {
	for _,tp := range tps {
		w.ProcessType(tp)
	}
}

func (w *Wrapper) Wrap(toproc string) {

	w.Processed = make(map[string]bool)

	w.cCode.WriteString(`/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#import <Foundation/Foundation.h>
`)

	pInterfaces := map[string]Interface {toproc: w.Interfaces[toproc]}
	for k,v := range pInterfaces {
		if Debug {
			fmt.Printf("Interface %s: %d properties, %d methods\n",
			k, len(v.Properties), len(v.Methods))
		}

		w.goCode.WriteString(fmt.Sprintf(`
func New%s() *%s {
	ret := &%s{}
	ret.ptr = unsafe.Pointer(C.New%s())
	return ret
}
`,v.Name,v.Name,v.Name,v.Name))

		w.cCode.WriteString(fmt.Sprintf(`
%s*
New%s() {
	return [[%s alloc] init];
}
`, v.Name, v.Name, v.Name))

		for _,y := range v.Properties {
			if Debug {
				fmt.Printf("  property: %s (%s)\n", y.Name, typeOrType2(y.Type,y.Type2))
			}
		}
		for _,y := range v.Methods {
			if Debug {
				fmt.Printf("  method: %s (%s)\n", y.Name, typeOrType2(y.Type,y.Type2))
			}
			gname := strings.Title(y.Name)

			grtype := ""
			grptr := ""
			cmtype := cType(typeOrType2(y.Type,y.Type2),v.Name)
			if !y.isVoid() {
				var err error
				grtype,err = goType(cmtype,v.Name)
				if err != nil {
					grtype = fmt.Sprintf("// goType(%s): NOT IMPLEMENTED\n",cmtype)
					continue
				}
			}
			gcast := ""
			gcast2 := ""
			if grtype != "" {
				if _,ok := w.Interfaces[grtype]; ok {
					grptr = "*"
					gcast = "ret := &" + grtype + "{}\n	ret.ptr = unsafe.Pointer("
					gcast2 = ")\n	return ret"
				} else {
					gcast = fmt.Sprintf("return (%s)(",grtype)
					gcast2 = ")"
				}
			}
			if grtype == "id" { // can't return id
				continue
			}
			w.ProcessType(grtype)
			gplist, gptypes, ok := w.goparamlist(y)
			if !ok {
				continue
			}
			w.ProcessTypes(gptypes)
			w.goCode.WriteString(fmt.Sprintf(`
func (o *%s) %s(%s) %s%s {
	%sC.%s_%s(%s)%s
}`,v.Name, gname, gplist, grptr, grtype, gcast, v.Name, y.Name, w.goparamnames(y),gcast2))

			cret := ""
			if !y.isVoid() {
				cret = "return "
			}
			w.cCode.WriteString(fmt.Sprintf(`
%s
%s_%s(%s) {
	%s[(id)obj %s];
}`, cmtype, v.Name, y.Name, w.cparamlist(y), cret, y.objcparamlist()))
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
