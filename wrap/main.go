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
	"float": "float",
	"double": "float",
	"complex float": "C.complexfloat",
	"complex double": "C.complexdouble",
}

var gobuiltinTypes map[string]bool = map[string]bool{
	"byte": true,
	"int": true,
	"float": true,
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

func (w *Wrapper) AddType(t string) {
	if _,ok := builtinTypes[t]; ok {
		return
	}
	nt, err := goType(t)
	if err != nil {
		return
	}
	w.Types[nt] = t
}

func goType(t string) (string, error) {
	// strip off pointers, < > and blank space
	nt := strings.ReplaceAll(t,"*","")
	if len(nt) > 3 && nt[0:3] == "id<" {
		nt = t[3:len(t)-1]
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

func NewWrapper() *Wrapper {
	return &Wrapper{
		Interfaces: map[string]Interface{},
		Types: map[string]string{},
	}
}

type Property struct {
	Name, Type, Type2, Attr string
}

type Parameter struct {
	Name, Type, Type2 string
}

type Method struct {
	Name, Type, Type2 string
	Parameters map[string]Parameter
}

func (m Method) isVoid() bool {
	return typeOrType2(m.Type,m.Type2) == "void"
}

func (m Method) cparamlist() string {
	ret := []string{"void* obj"}
	for k,v := range m.Parameters {
		ret = append(ret,fmt.Sprintf("%s %s",typeOrType2(v.Type,v.Type2),k))
	}
	return strings.Join(ret,", ")
}

func (m Method) objcparamlist() string {
	if len(m.Parameters) == 0 {
		return m.Name
	}
	ret := []string{fmt.Sprintf("%s:",m.Name)}
	for k,_ := range m.Parameters {
		ret = append(ret, fmt.Sprintf("%s:%s",k,k))
	}
	return strings.Join(ret," ")
}

func (m Method) goparamlist() (string,[]string) {
	ret := []string{}
	tps := []string{}
	for k,v := range m.Parameters {
		tp,err := goType(typeOrType2(v.Type,v.Type2))
		if err != nil {
			return "UNSUPPORTED TYPE", []string{}
		}
		tps = append(tps,tp)
		ret = append(ret,fmt.Sprintf("%s %s",k,tp))
	}
	return strings.Join(ret,", "),tps
}

func (m Method) goparamnames() string {
	ret := []string{"o.ptr"}
	for k,_ := range m.Parameters {
		ret = append(ret,k)
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
		w.AddType(name)
	}
	for _,c := range ns {
		switch x := c.(type) {
		case *ast.ObjCPropertyDecl:
			p := Property{
				Name: x.Name,
				Type: x.Type,
				Type2: x.Type2,
			}
			w.AddType(typeOrType2(x.Type,x.Type2))
			i.Properties[p.Name] = p
		case *ast.ObjCMethodDecl:
			//fmt.Printf("ObjCMethodDecl: %s\n",x.Name)
			m := Method{
				Name: x.Name,
				Type: x.Type,
				Type2: x.Type2,
			}
			w.AddType(typeOrType2(x.Type,x.Type2))
			m.Parameters = w.GetParms(x)
			i.Methods[m.Name] = m
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

func (w *Wrapper) GetParms(n *ast.ObjCMethodDecl) map[string]Parameter {
	ps := make([]Parameter,0)
	avail := make([][]string,len(n.Children()))
	var c ast.Node
	i := -1
	for i,c = range n.Children() {
		switch x := c.(type) {
		case *ast.ParmVarDecl:
			i++
			p := Parameter{
				Type: x.Type,
				Type2: x.Type2,
			}
			ps = append(ps,p)
			avail = append(avail,[]string{})
		case *ast.AvailabilityAttr:
			avail[i] = append(avail[i],x.OS)
		}
	}
	isAvail := func(l []string) bool {
		if len(l) == 0 {
			return true
		}
		for _,x := range l {
			if x == "macos" {
				return true
			}
		}
		return false
	}
	ret := make(map[string]Parameter)
	j := 0
	for i,p := range ps {
		if isAvail(avail[i]) {
			ret[n.Parameters[j]] = p
			j++
		}
	}
	for _,p := range ret {
		w.AddType(typeOrType2(p.Type,p.Type2))
	}
	if j != len(n.Parameters) {
		fmt.Printf("Error in method declaration %s: Wrong number of ParmVarDecl children: %d parameters but %d ParmVarDecl children\n",n.Name,len(n.Parameters),i)
	}
	return ret
}

func typeOrType2(t1, t2 string) string {
	if t2 != "" {
		return t2
	}
	return t1
}

func (w *Wrapper) ProcessType(gotype string) {
	if gotype == "" {
		return
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
#cgo LDFLAGS: -framework Foundation -framework CoreBluetooth
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
			cmtype := typeOrType2(y.Type,y.Type2)
			if !y.isVoid() {
				var err error
				grtype,err = goType(cmtype)
				if err != nil {
					grtype = fmt.Sprintf("// goType(%s): NOT IMPLEMENTED\n",cmtype)
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
			w.ProcessType(grtype)
			gplist, gptypes := y.goparamlist()
			w.ProcessTypes(gptypes)
			w.goCode.WriteString(fmt.Sprintf(`
func (o *%s) %s(%s) %s%s {
	%sC.%s_%s(%s)%s
}`,v.Name, gname, gplist, grptr, grtype, gcast, v.Name, y.Name, y.goparamnames(),gcast2))

			cret := ""
			if !y.isVoid() {
				cret = "return "
			}
			w.cCode.WriteString(fmt.Sprintf(`
%s
%s_%s(%s) {
	%s[(id)obj %s];
}`, cmtype, v.Name, y.Name, y.cparamlist(), cret, y.objcparamlist()))
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
