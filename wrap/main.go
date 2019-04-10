package wrap

import (
	"fmt"
	//"reflect" 
	"strings"

	"gitlab.wow.st/gmp/clast/ast"
)

type Wrapper struct {
	Interfaces map[string]Interface
	Types map[string]string
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

func (w *Wrapper) AddType(t string) {
	// only for objects and pointers?
	// strip off pointers, < > and blank space
	nt := strings.ReplaceAll(t,"*","")
	if len(nt) > 3 && nt[0:3] == "id<" {
		nt = t[3:len(t)-1]
	}
	nt = strings.ReplaceAll(nt,"<","_")
	nt = strings.ReplaceAll(nt,">","_")
	if _,ok := builtinTypes[nt]; ok { // do not add builtin types
		return
	}
	if len(nt) > 5 && nt[0:5] == "enum " { // FIXME: deal with enums?
		return
	}
	nt = strings.ReplaceAll(nt," ","")
	if nt == "void" {
		return
	}
	w.Types[nt] = t
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
			fmt.Printf("ObjcMethodDecl: %s\n",x.Name)
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
			//fmt.Printf("ast.ObjCInterface: %s\n",x.Name)
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

func (w *Wrapper) Wrap() {
	//var cCode strings.Builder
	var goCode strings.Builder

	for k,t := range w.Types { // FIXME: SORT FIRST
		w := fmt.Sprintf(`
// %s
type %s struct { ptr unsafe.Pointer }
`,t, k)
		goCode.WriteString(w)
	}
	for k,v := range w.Interfaces {
		fmt.Printf("Interface %s: %d properties, %d methods\n",
			k, len(v.Properties), len(v.Methods))
		for _,y := range v.Properties {
			fmt.Printf("  property: %s (%s)\n", y.Name, typeOrType2(y.Type,y.Type2))
		}
		for _,y := range v.Methods {
			fmt.Printf("  method: %s (%s)\n", y.Name, typeOrType2(y.Type,y.Type2))
			for _,z := range y.Parameters {
				fmt.Printf("    %s:%s\n", z.Name, typeOrType2(z.Type,z.Type2))
			}
		}
	}
	fmt.Println(goCode.String())
}
