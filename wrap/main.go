package wrap

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"git.wow.st/gmp/nswrap/ast"
	"git.wow.st/gmp/nswrap/types"
)

var (
	Debug = false
// Arc flag is for debugging only, your builds will break if you turn it on
	Arc = false
)

type Wrapper struct {
	Package string
	Interfaces map[string]*Interface
	Functions map[string]*Method
	NamedEnums map[string]*Enum
	AnonEnums []*Enum
	Delegates map[string]map[string][]string
	Subclasses map[string]*Subclass
	Protocols map[string]*Protocol

	cgoFlags strings.Builder  // put cGo directives here
	cImports strings.Builder  // put C imports and sysimports here
	cCode strings.Builder	  // put cGo code here
	goTypes strings.Builder   // put Go type declarations here
	goConst strings.Builder   // put Go constants (from C enums) here
	goCode strings.Builder    // put Go code here
	goExports strings.Builder // put exported Go functions here
	goHelpers strings.Builder // put Go helper functions here

	ProcessedTypes map[string]bool
	ProcessedClassMethods map[string]bool
	Vaargs int
}

func NewWrapper(debug bool) *Wrapper {
	Debug = debug
	if Debug { fmt.Println("// Debug mode") }
	ret := &Wrapper{
		Interfaces: map[string]*Interface{},
		Functions: map[string]*Method{},
		NamedEnums: map[string]*Enum{},
		AnonEnums: []*Enum{},
		Protocols: map[string]*Protocol{},
		Subclasses: map[string]*Subclass{},
		ProcessedTypes: map[string]bool{},
		ProcessedClassMethods: map[string]bool{},
		Vaargs: 16,
	}
	arc := " -fno-objc-arc"
	if Arc {
		arc = " -fobjc-arc"
	}
	ret.cgoFlags.WriteString(fmt.Sprintf(`/*
#cgo CFLAGS: -x objective-c%s
`,arc))
	ret.goTypes.WriteString(`
type Id struct {
	ptr unsafe.Pointer
}
func (o Id) Ptr() unsafe.Pointer { return o.ptr }
`)
	return ret
}

func (w *Wrapper) Frameworks(ss []string) {
	if len(ss) == 0 {
		return
	}
	for _,s := range ss {
		w.cImports.WriteString(fmt.Sprintf("#import <%s/%s.h>\n",s,s))
	}
	w.cgoFlags.WriteString("#cgo LDFLAGS: -framework " + strings.Join(ss," -framework "))
}

func (w *Wrapper) Import(ss []string) {
	for _,s := range ss {
		w.cImports.WriteString("\n#import \"" + s + "\"\n")
	}
}

func (w *Wrapper) SysImport(ss []string) {
	for _,s := range ss {
		w.cImports.WriteString("\n#import <" + s + ">\n")
	}
}

func (w *Wrapper) Pragma(ss []string) {
	for _,s := range ss {
		w.cgoFlags.WriteString("\n#pragma " + s + "\n")
	}
}

func (w *Wrapper) Delegate(ds map[string]map[string][]string) {
	w.Delegates = ds
}

func (w *Wrapper) Subclass(ds map[string]map[string][]string) {
	for k,v := range ds {
		sc := &Subclass{
			Overrides: []string{},
			NewMethods: []string{},
		}
		if len(ds) == 0 {
			fmt.Printf("No superclass specified for subclass %s\n",k)
			os.Exit(-1)
		}
		if len(ds) > 1 {
			fmt.Printf("Multiple inheritance not permitted for subclass %s\n",k)
			os.Exit(-1)
		}
		sc.Name = k
		for x,y := range(v) {
			sc.Super = x
			for _,m := range y {
				switch m[0] {
				case '-','+':
					sc.NewMethods = append(sc.NewMethods,m)
				default:
					sc.Overrides = append(sc.Overrides,m)
				}
			}
		}
		w.Subclasses[sc.Name] = sc
	}
}

type Subclass struct {
	Name,Super string
	Overrides []string
	NewMethods []string
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
	Name, GoName, Class, GoClass string
	Type *types.Type
	ClassMethod bool
	Parameters []*Parameter
}

//Fully disambiguated method name (m.GoName + all parameter names)
func (m *Method) LongName() string {
	ret := m.GoName
	for _,p := range m.Parameters {
		ret = ret + p.Pname
	}
	return ret
}

func (m *Method) HasUnsupportedType() bool {
	return	m.Type.IsFunction() ||
		m.Type.IsFunctionPtr() ||
		m.hasUnsupportedParam()
}

type Enum struct {
	Name string
	Type *types.Type
	Constants []string
}

type Protocol struct {
	InstanceMethods, ClassMethods *MethodCollection
}

type MethodCollection struct {
	Class string
	Methods []*Method
}

func NewMethodCollection(class string) *MethodCollection {
	ret := &MethodCollection{
		Class: class,
		Methods: []*Method{},
	}
	return ret
}

type ByParams []*Method

func (a ByParams) Len() int { return len(a) }
func (a ByParams) Swap(i,j int) { a[i], a[j] = a[j], a[i] }
func (a ByParams) Less(i, j int) bool { return len(a[i].Parameters) < len(a[j].Parameters) }

//Disambiguate overloaded method names
func Disambiguate(mc *MethodCollection) {
	lst := map[string][]*Method{}
	for _,m := range mc.Methods {
		lst2 := lst[m.Name]
		if lst2 == nil {
			lst2 = []*Method{m}
		}  else {
			lst2 = append(lst2,m)
		}
		lst[m.Name] = lst2
	}
	mc.Methods = []*Method{}
	used := map[string]bool{}
	for _,v := range lst {
		sort.Sort(ByParams(v))
		for _,m := range v {
			if len(v) < 2 || len(m.Parameters) < 2 {
				mc.Methods = append(mc.Methods,m)
				continue
			}
			i := 2
			pname := m.Name + strings.Title(m.Parameters[1].Pname)
			for ; used[pname] && i < len(m.Parameters); i++ {
				pname = pname + strings.Title(m.Parameters[i].Pname)
			}
			used[pname] = true
			m.GoName = strings.Title(pname)
			mc.Methods = append(mc.Methods,m)
		}
	}
}


//isVoid() returns true if the method has no return value.
func (m Method) isVoid() bool {
	return m.Type.CType() == "void"
}

//hasUnsupportedParam() returns true if a method has a function as a parameter.
func (m Method) hasUnsupportedParam() bool {
	for _,p := range m.Parameters {
		if p.Type.IsFunction() || p.Type.IsFunctionPtr() {
			return true
		}
		if pt := p.Type.PointsTo(); pt.IsValist() {
			return true
		}
	}
	return false
}

func (w Wrapper) cparamlist(m *Method) (string,string,string) {
	ns := make([]string,0)
	ret := make([]string,0)
	tps := []string{"void*"}
	if !m.ClassMethod {
		ret = append(ret,"void* o")
	}
	for _,p := range m.Parameters {
		var tp string
		wp := types.ShouldWrap(p.Type.GoType())
		if wp || p.Type.IsPointer() || p.Type.Variadic {
			tp = "void*"
		} else {
			tp = p.Type.CType()
		}
		ns = append(ns,p.Vname)
		ret = append(ret,fmt.Sprintf("%s %s",tp,p.Vname))
		tps = append(tps,tp)
	}
	return strings.Join(ns,", "),strings.Join(ret,", "),strings.Join(tps,", ")
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
				str := []string{p.Pname + ", arr[0]"}
				for i := 1; i < w.Vaargs; i++ {
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
	"len": true,
}

func (w *Wrapper) gpntp(m *Method) ([]string,[]string,[]*types.Type,string) {
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
	snames := make([]string,len(ns))
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
		if len(gt) > 2 && gt[:2] == "**" && types.ShouldWrap(gt[2:]) {
			x := gt[2:]
			if types.IsGoInterface(x) {
				x = "Id"
			}
			gt = "*[]" + x
			snames[i] = "goSlice" + strconv.Itoa(i)
		}
		ret = append(ret,ns[i] + " " + gt)
	}
	return ns, snames, tps, strings.Join(ret,", ")
}


type Interface struct {
	Name, GoName string
	InstanceMethods, ClassMethods *MethodCollection
	Properties map[string]*Property
	Protocols []string // Protocols impelemented by this Interface
	ProcessedInstanceMethods map[string]bool
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
		GoName: strings.Title(n.Name),
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
		case *ast.Variadic:
			p := &Parameter{
				Vname: "object",
				Type: types.NewTypeFromString("NSObject*",""),
			}
			p.Type.Variadic = true
			m.Parameters = append(m.Parameters,p)
			i++
		}
	}
	if i > 0 && len(f.Children) > i {
		if e := f.Children[i]; len(e.Children) > 0 {
			//fmt.Println("  Next parameter: ",e.Children[0].String())
			//m.Parameters[i-1].Type.Variadic = true
		}
	}
	w.Functions[n.Name] = m
}

func (w *Wrapper) AddProtocol(n *ast.ObjCProtocolDecl) {
	p := w.Protocols[n.Name]
	if p == nil {
		//fmt.Printf("Adding protocol %s\n",n.Name)
		p = &Protocol{ }
		//p.GoName = types.NewTypeFromString(n.Name,n.Name).GoType()
		p.ClassMethods = NewMethodCollection(n.Name)
		p.InstanceMethods = NewMethodCollection(n.Name)
	}
	//fmt.Printf("Protocol %s\n",p.Name)
	for _,c := range n.Children() {
		switch x := c.(type) {
		case *ast.ObjCMethodDecl:
			if Arc {
				switch x.Name {
				case "retain","release","autorelease":
					continue
				}
			}
			if x.ClassMethod {
				w.AddMethod(p.ClassMethods,x)
			} else {
				w.AddMethod(p.InstanceMethods,x)
			}
		}
	}
	Disambiguate(p.InstanceMethods)
	Disambiguate(p.ClassMethods)
	w.Protocols[n.Name] = p
}

func (w *Wrapper) AddMethod(p *MethodCollection, x *ast.ObjCMethodDecl) {
	m := &Method{
		Name: x.Name,
		GoName: strings.Title(x.Name),
		Type: types.NewTypeFromString(x.Type,p.Class),
		Class: p.Class,
		GoClass: strings.Title(p.Class),
		ClassMethod: x.ClassMethod,
	}
	//fmt.Printf("  -- Method %s\n",m.Name)
	var avail bool
	m.Parameters, avail = w.GetParms(x,p.Class)
	if avail {
		//fmt.Printf("%s: Adding %s (%d)\n",p.Class,m.Name,len(m.Parameters))
		//fmt.Printf("--Method name is %s\n",m.Name)
		p.Methods = append(p.Methods,m)
	}
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
		case *ast.AvailabilityAttr, *ast.UnavailableAttr, *ast.DeprecatedAttr:
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

//Add an Interface or add a Category to an Interface
func (w *Wrapper) add(name string, ns []ast.Node) {
	var i *Interface
	var ok bool
	goname := strings.Title(types.NewTypeFromString(name,name).GoType())
	types.Wrap(goname)
	if i,ok = w.Interfaces[name]; !ok {
		i = &Interface{ }
		i.Name = name
		i.GoName = goname
		i.Properties = map[string]*Property{}
		i.InstanceMethods = NewMethodCollection(name)
		i.ClassMethods = NewMethodCollection(name)
		i.Protocols = []string{}
		i.ProcessedInstanceMethods = map[string]bool{}
		m := &Method{
			Name: "class",
			GoName: "Class",
			Class: i.Name,
			GoClass: i.Name,
			Type: types.NewTypeFromString("Class",i.Name),
			ClassMethod: true,
			Parameters: []*Parameter{},
		}
		i.ClassMethods.Methods = []*Method{m}
	}
	//var avail bool
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
			if name == "NSObject" && x.Name == "initialize" {
				continue
			}
			if x.ClassMethod {
				w.AddMethod(i.ClassMethods,x)
			} else {
				w.AddMethod(i.InstanceMethods,x)
			}
		case *ast.ObjCProtocol:
			//fmt.Printf("ast.ObjCProtocol: %s\n",x.Name)
			i.Protocols = append(i.Protocols,x.Name)
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
			//depth first
			supmethods(i,types.Super(s))
			for _,m := range sup.ClassMethods.Methods {
				m2 := &Method{
					Name: m.Name,
					GoName: m.GoName,
					Class: i.Name,
					GoClass: i.GoName,
					Type: m.Type.CloneToClass(i.Name),
					ClassMethod: true,
					Parameters: []*Parameter{},
				}
				for _,p := range m.Parameters {
					p2 := &Parameter{
						Pname: p.Pname,
						Vname: p.Vname,
						Type: p.Type.CloneToClass(i.Name),
					}
					m2.Parameters = append(m2.Parameters,p2)
				}
				found := false
				longname := m2.LongName()
				for n,x := range i.ClassMethods.Methods {
					if x.LongName() == longname {
						i.ClassMethods.Methods[n] = m2
						found = true
					}
				}
				if !found {
					i.ClassMethods.Methods = append(i.ClassMethods.Methods,m2)
				}
			}
		}
	}
	supmethods(i,types.Super(i.Name))
	//fmt.Println("Add interface ",i.Name)
	Disambiguate(i.ClassMethods)
	Disambiguate(i.InstanceMethods)
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
	case *ast.UnavailableAttr, *ast.DeprecatedAttr:
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
			//ret[j-1].Type.Variadic = true
			p := &Parameter{
				Vname: "object",
				Type: types.NewTypeFromString("NSObject*",""),
			}
			p.Type.Variadic = true
			ret = append(ret,p)
			j++
		case *ast.AvailabilityAttr, *ast.UnavailableAttr, *ast.DeprecatedAttr:
			avail.Add(x)
		case *ast.Unknown:
			if Debug { fmt.Printf("GetParms(): ast.Unknown: %s\n",x.Name) }
		}
	}
	// check that the method is available for this OS and not deprecated
	if !avail.Available() {
		return nil, false
	}
	return ret, true
}

func (w *Wrapper) AddTypedef(n,t string) {
	tp := types.NewTypeFromString(t,"")
	gt := tp.GoType()
	//fmt.Printf("Typedef %s -> %s\n",n,t)
	if types.ShouldWrap(gt) {
		//fmt.Printf("  should wrap\n")
		//fmt.Printf("  processing type for %s (%s)\n",n,gt)
		types.Wrap(n)
		types.SetSuper(n,gt)
		w._processType(tp,"*" + n)
	} else {
		types.AddTypedef(n,tp)
	}
}

func (w *Wrapper) processTypes(tps []*types.Type) {
	for _,tp := range tps {
		w.processType(tp)
	}
}

func (w *Wrapper) processType(tp *types.Type) {
	bt := tp.BaseType()
	gt := bt.GoType()
	w._processType(bt,gt)
}

func (w *Wrapper) _processType(bt *types.Type, gt string) {
	//fmt.Printf("processType: gt = %s bt = %s\n",gt,bt)
	if gt == "" {
		return
	}
	if gt[0] == '*' {
		w.processType(bt.PointsTo())
		return
	}
	if w.ProcessedTypes[gt] { return }
	w.ProcessedTypes[gt] = true
	if gt == "Char" {
		w.CharHelpers()
	}
	if gt == "SEL" {
		w.SelectorHelpers()
	}
	if gt == "NSAutoreleasePool" {
		w.AutoreleaseHelpers()
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

func CharWithBytes(b []byte) *Char {
	return (*Char)(unsafe.Pointer(C.CString(string(b))))
}

func (c *Char) String() string {
	return C.GoString((*C.char)(c))
}
`)
}

func (w *Wrapper) StringHelpers() {
	w.goHelpers.WriteString(`
func (o NSString) String() string {
	return o.UTF8String().String()
}
`)
}

func (w *Wrapper) EnumeratorHelpers() {
	w.goHelpers.WriteString(`
func (e NSEnumerator) ForIn(f func(Id) bool) {
	for o := e.NextObject(); o.Ptr() != nil; o = e.NextObject() {
		if !f(o) { break }
	}
}
`)
}

func (w *Wrapper) AutoreleaseHelpers() {
	//not sure why this is not coming up automatically...
	w.cCode.WriteString(`
void* _Nonnull
NSAutoreleasePool_init(void* o) {
	return [(NSAutoreleasePool*)o init];
}
`)
	w.goHelpers.WriteString(`
func (o NSAutoreleasePool) Init() NSAutoreleasePool {
	ret := NSAutoreleasePool{}
	ret.ptr = C.NSAutoreleasePool_init(o.Ptr())
	return ret
}

func Autoreleasepool(f func()) {
	pool := NSAutoreleasePoolAlloc().Init()
	f()
	pool.Drain()
}
`)
}

func (w *Wrapper) SelectorHelpers() {
	w.cCode.WriteString(`
void*
selectorFromString(char *s) {
	return NSSelectorFromString([NSString stringWithUTF8String:s]);
}
`)
	w.goHelpers.WriteString(`
func Selector(s string) SEL {
	return (SEL)(unsafe.Pointer(C.selectorFromString(C.CString(s))))
}
`)
}

func (w *Wrapper) ProcessMethod(m *Method) {
	w._processMethod(m,false)
}

func (w *Wrapper) ProcessMethodForClass(m *Method, class string) {
	goclass := strings.Title(types.NewTypeFromString(class,class).GoType())
	m2 := &Method{
		Name: m.Name, GoName: m.GoName, Class: class, GoClass: goclass,
		Type: m.Type.CloneToClass(class),
		ClassMethod: m.ClassMethod,
		Parameters: make([]*Parameter,len(m.Parameters)),
	}
	for i,p := range m.Parameters {
		m2.Parameters[i] = &Parameter{
			Pname: p.Pname, Vname: p.Vname,
			Type: p.Type.CloneToClass(class),
		}
	}
	w._processMethod(m2,false)
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
	if m.HasUnsupportedType() {
		return
	}
	w.processType(m.Type)
	gname := m.GoName
	gname = strings.ReplaceAll(gname,"_"," ")
	gname = strings.Title(gname)
	gname = strings.ReplaceAll(gname," ","")
	receiver := ""
	var cname string
	if fun {
		cname = m.Name
	} else {
		cname = gname
	}
	//fmt.Printf("Method %s (GoClass %s)\n",cname,m.GoClass)
	switch {
	case !m.ClassMethod:
		if types.IsGoInterface(m.GoClass) {
			receiver = "(o Id) "
		} else {
			receiver = "(o " + m.GoClass + ") "
		}
		//Disambiguate instance methods with same name as a class method
		cname = "inst_" + cname
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
	if m.Class != "" {
		cname = m.Class + "_" + cname
	}
	var cmtype string
	if m.Type.IsPointer() || types.ShouldWrap(m.Type.GoType()) {
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
	ns,snames,tps,gplist := w.gpntp(m)
	grtype := m.Type.GoType()
	if grtype == "Void" {
		grtype = ""
	}
	if types.IsGoInterface(grtype) {
		grtype = "Id"
	}
	if grtype == "BOOL" { // convert objective-c bools to Go bools
		grtype = "bool"
	}
	if gname == grtype { // avoid name conflicts between methods and types
		gname = "Get" + gname
	}
	if m.ClassMethod {
		if w.ProcessedClassMethods[gname] {
			return
		}
		w.ProcessedClassMethods[gname] = true
	} else {
		i, ok := w.Interfaces[m.Class]
		if !ok {
			fmt.Printf("Can't find interface %s for method %s\n",m.Class,m.Name)
			os.Exit(-1)
		}
		if i.ProcessedInstanceMethods[gname] {
			return
		}
		i.ProcessedInstanceMethods[gname] = true
	}

	w.goCode.WriteString(fmt.Sprintf(`
func %s%s(%s) %s {
`,receiver,gname,gplist,grtype))
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
`,vn,w.Vaargs,vn,vn))
	}
	for i,n := range ns {
		if snames[i] == "" {
			continue
		}
		w.goCode.WriteString(fmt.Sprintf(`
	%s := make([]unsafe.Pointer,cap(*%s))
	for i := 0; i < len(*%s); i++ {
		%s[i] = (*%s)[i].Ptr()
	}
`,snames[i],n,n,snames[i],n))
	}
	w.goCode.WriteString(`	` +
		types.GoToC(cname,ns,snames,m.Type,tps,fun) + "\n}\n")

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
	cns,cntps,_ := w.cparamlist(m)
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
			if gt == "NSString" {
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
	w.processType(e.Type)
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
		if !w.ProcessedTypes[gtp] {
			w.goTypes.WriteString(fmt.Sprintf(`
type %s %s
`,gtp,ctp))
			w.ProcessedTypes[gtp] = true
		}
	}
	gtp = gtp + " "
	//fmt.Printf("  gtp = %s; ctp = %s\n",gtp,ctp)
	for _,c := range e.Constants {
		w.goConst.WriteString(fmt.Sprintf(`const %s %s= C.%s
`,c,gtp,c))
	}
	w.goConst.WriteString("\n")
}

func (w *Wrapper) MethodFromSig(sig,class string) *Method {
	ret := &Method{ Parameters: []*Parameter{} }
	if len(sig) == 0 {
		return ret
	}
	if sig[0] == '+' {
		ret.ClassMethod = true
	}
	sig = sig[1:]
	rem,n := types.MethodSignature(sig,types.NewNode("AST"))
	if len(rem) > 0 {
		fmt.Printf("Failed to parse method signature %s (%s)\n",sig,rem)
		os.Exit(-1)
	}
	i := 0 // count MethodParameters
	for _,c := range n.Children {
		switch c.Kind {
		case "TypeName":
			tp := types.NewType(c,class)
			ret.Type = tp
		case "Identifier":
			ret.Name = c.Content
			ret.GoName = strings.Title(c.Content)
		case "MethodParameter":
			p := &Parameter{}
			for _,d := range c.Children {
				switch d.Kind {
				case "TypeName":
					tp := types.NewType(d,class)
					p.Type = tp
				case "Identifier":
					if i == 0 || p.Pname != "" {
						p.Vname = d.Content
					} else {
						p.Pname = d.Content
					}
				}
			}
			i++
			ret.Parameters = append(ret.Parameters,p)
		}
	}
	return ret
}

func (w *Wrapper) ProcessSubclass(sname string, sc *Subclass) {
	gname := strings.Title(sname)
	types.Wrap(gname)
	types.SetSuper(gname,sc.Super)
	ps := map[string][]string{}
	ps[sc.Super] = sc.Overrides
	nms := make([]*Method,len(sc.NewMethods))
	for i,sig := range sc.NewMethods {
		nms[i] = w.MethodFromSig(sig,sname)
	}
	w._ProcessDelSub(sname,ps,nms,true)
}

func (w *Wrapper) ProcessDelegate(dname string, ps map[string][]string) {
	w._ProcessDelSub(dname,ps,nil,false)
}

//NOTE: The delegate wrapper does not support variadic callback functions.
func (w *Wrapper) _ProcessDelSub(dname string, ps map[string][]string, nms []*Method, sub bool) {
	//To create (per delegate):
	//1. ObjC interface
	//2. ObjC implementation
	//3. ObjC constructor function
	//4. Go type
	//5. Go constructor
	//6. Go dispatch database for callbacks
	//7. Go superclass dispatch function
	//To create (per method):
	//1. ObjC function prototypes for go exports
	//2. Go callback registration functions
	//3. Go exported callback function wrappers
	//4. Go wrapper functions for superclass methods

	//organize output into string builders
	var cprotos, ccode, gotypes, gocode, goexports strings.Builder

	//set up array of methods for this delegate
	methods := []*Method{}
	sms := 0 // the number of methods that have super-methods
	gnames := []string{} // go names for methods
	pnames := make([]string,len(ps))
	var supr string
	i := 0
	for pname,pats := range ps {
		pnames[i] = pname
		i++
		var ms []*Method
		if sub {
			interf := w.Interfaces[pname]
			supr = interf.GoName
			if i > 1 {
				fmt.Printf("Multiple inheritance is not permitted:\n    subclass %s already inherits from %s\n",dname,supr)
				os.Exit(-1)
			}
			if interf == nil {
				fmt.Printf("Failed to find interface %s for subclass %s\n",pname,dname)
				os.Exit(-1)
			}
			//fmt.Printf("  subclass for %s\n",pname)
			ms = []*Method{}
			var addmeths func(s string)
			addmeths = func(s string) {
				if sup := types.Super(s); w.Interfaces[sup] != nil {
					addmeths(sup)
				}
				//fmt.Printf("Adding methods for %s\n",s)
				for _,m := range w.Interfaces[s].InstanceMethods.Methods {
					ms = append(ms,m)
				}
			}
		//for subclasses, add all superclass methods, depth first
			addmeths(interf.Name)
		} else { // not a subclass
			proto := w.Protocols[pname]
			if proto == nil {
				fmt.Printf("Failed to find protocol %s for delegate %s\n",pname,dname)
				os.Exit(-1)
			}
			//fmt.Printf("  proto %s\n",pname)
			ms = proto.InstanceMethods.Methods
			fmt.Printf("Protocol %s\n",pname)
			types.SetSuper(dname,"Id")
			supr = "Id"
		}
		for _,m := range ms {
	//note:we may have capitalized the first character to make a GoName...
			if !matches(string(m.Name[0])+m.GoName[1:],pats) {
				continue
			}
			if m.HasUnsupportedType() {
				continue
			}
			methods = append(methods,m)
			gnames = append(gnames,m.GoName)
			if sub { sms = len(methods) }
		}
	}
	//add new methods being defined for the subclass
	if sub {
		for _,m := range nms {
			methods = append(methods,m)
			gnames = append(gnames,strings.Title(m.Name))
		}
	}

	methprotos := make([]string,len(methods)) // objc method prototypes
	smethprotos := make([]string,sms) // super method prototypes
	sfunprotos := make([]string,sms) // super method prototypes
	gname := strings.Title(dname) // go name for this Delegate
	vnames := make([][]string,len(methods)) // objc variable names
	vpnames := make([][]string,len(methods)) // objc parameter:variable names
	gtypes := make([][]string,len(methods)) // go parameter types for each method
	getypes := make([][]string,len(methods)) // parameter types for go export
	grtypes := make([]string,len(methods)) // go retrun types for each method
	cgtypes := make([]string,len(methods)) // cgo return types
	crtypes := make([]string,len(methods)) // c return types for each method

	//1. ObjC interface
	if sub {
		fmt.Printf("Subclass %s <%s>: %d overrides, %d new methods\n",dname,strings.Join(pnames,", "),sms, len(nms))
	} else {
		fmt.Printf("Delegate %s <%s>: %d methods\n",dname,strings.Join(pnames,", "),len(methods))
	}
	for i,m := range methods {
		w.processType(m.Type)
		vnames[i] = make([]string,len(m.Parameters)+1)
		vpnames[i] = make([]string,len(m.Parameters))
		getypes[i] = make([]string,len(m.Parameters)+1)
		vnames[i][0] = "self"
		getypes[i][0] = "unsafe.Pointer"
		gtypes[i] = make([]string,len(m.Parameters)+1)
		gtypes[i][0] = gname + "Supermethods"
		//fmt.Printf("%s: %s\n",dname,m.Name)
		var parms string
		var cparms string
		if len(m.Parameters) == 0 {
			parms = ""
			cparms = "void* self"
			vpnames[i] = []string{m.Name}
		} else {
			pm := m.Parameters[0]
			w.processType(pm.Type)
			parms = fmt.Sprintf(":(%s)%s",pm.Type.Node.CType(),pm.Vname)
			cparms = fmt.Sprintf("void* self, %s %s",pm.Type.Node.CType(),pm.Vname)
			vnames[i][1] = pm.Vname
			vpnames[i][0] = pm.Pname + ":" + pm.Vname
			gtypes[i][1] = pm.Type.GoType()
			if pm.Type.IsPointer() {
				getypes[i][1] = "unsafe.Pointer"
			} else {
				getypes[i][1] = gtypes[i][1]
			}
		}
		for j := 1; j < len(m.Parameters); j++ {
			pm := m.Parameters[j]
			w.processType(pm.Type)
			parms = parms + fmt.Sprintf(" %s:(%s)%s",pm.Pname,pm.Type.Node.CType(),pm.Vname)
			cparms = cparms + fmt.Sprintf(", %s %s",pm.Type.Node.CType(),pm.Vname)
			vnames[i][j+1] = pm.Vname
			vpnames[i][j] = pm.Pname + ":" + pm.Vname
			gtypes[i][j+1] = pm.Type.GoType()
			var getp string
			if pm.Type.IsPointer() {
				getp = "unsafe.Pointer"
			} else {
				getp = gtypes[i][j+1]
			}
			getypes[i][j+1] = getp
		}
		methprotos[i] = fmt.Sprintf(
`- (%s)%s%s;`,m.Type.Node.CType(),m.Name,parms)
		ct := m.Type.Node.CType()
		if i < sms {
			smethprotos[i] = fmt.Sprintf(
`- (%s)super_%s%s;`,ct,m.Name,parms)
		}
		if ct == "instancetype" {
			ct = gname + "*"
		}
		if i < sms {
			sfunprotos[i] = fmt.Sprintf(
`%s %s_super_%s(%s);`,ct,dname,m.Name,cparms)
		}
		if x := m.Type.GoType(); x == "Void" {
			grtypes[i] = ""
		} else {
			grtypes[i] = " " + x
		}
		crtypes[i] = m.Type.CTypeAttrib()
		if m.Type.IsPointer() {
			cgtypes[i] = "unsafe.Pointer"
		} else {
			crtypes[i] = m.Type.CTypeAttrib()
			cgtypes[i] = m.Type.CGoType()
		}
}
	var supcls string
	var protos string
	if sub {
		supcls = pnames[0]
		protos = ""
	} else {
		supcls = "NSObject"
		protos = "<" + strings.Join(pnames,", ") + ">"
	}
	ccode.WriteString(fmt.Sprintf(`
@interface %s : %s %s
{ }
%s
`,dname,supcls,protos,strings.Join(methprotos,"\n")))
	if sub {
		ccode.WriteString(strings.Join(smethprotos,"\n"))
	}
	ccode.WriteString(`
@end
`)
	if sub {
		ccode.WriteString(strings.Join(sfunprotos,"\n"))
	}

	//2. ObjC implementation
	methdecls := make([]string,len(methods))
	smethdecls := make([]string,len(methods))
	sfundecls := make([]string,len(methods))
	for i,mp := range methprotos {
		mp := mp[:len(mp)-1]
		var smp, sfp string
		if sub && i < sms {
			smp = smethprotos[i][:len(smethprotos[i])-1]
			sfp = sfunprotos[i][:len(sfunprotos[i])-1]
		}
		var ret string
		if crtypes[i] != "void" {
			ret = "return "
		}
		methdecls[i] = fmt.Sprintf(`
%s
{
	%s%s(%s);
}
`,mp,ret,gname + gnames[i],strings.Join(vnames[i],", "))
		methdecls[i] = fmt.Sprintf(`
%s
{
	%s%s(%s);
}
`,mp,ret,gname + gnames[i],strings.Join(vnames[i],", "))
		if sub && i < sms {
			smethdecls[i] = fmt.Sprintf(`
%s
{
	%s[super %s];
}
`,smp,ret,strings.Join(vpnames[i]," "))
			sfundecls[i] = fmt.Sprintf(`
%s
{
	%s[(%s*)self super_%s];
}
`,sfp,ret,dname,strings.Join(vpnames[i]," "))
		}
	}
	ccode.WriteString(fmt.Sprintf(`
@implementation %s
%s
`,dname,strings.Join(methdecls,"\n")))
	if sub {
		ccode.WriteString(strings.Join(smethdecls,"\n"))
	}
	ccode.WriteString(`
@end
`)
	if sub {
		ccode.WriteString(strings.Join(sfundecls,"\n"))
	}

	//3. ObjC constructor function
	ccode.WriteString(fmt.Sprintf(`
void*
%sAlloc() {
	return [[%s alloc] autorelease];
}
`,dname,dname))

	//4. Go type
	
	gotypes.WriteString(
		types.NewTypeFromString(gname,supr).GoInterfaceDecl())

	//5. Go constructor
	gocode.WriteString(fmt.Sprintf(`
func %sAlloc() %s {
	ret := %s{}
	ret.ptr = unsafe.Pointer(C.%sAlloc())
	return ret
}
`,gname,gname,gname,dname))

	//6. Go dispatch database for callbacks
	dispitems := make([]string,len(gnames))
	sdispitems := make([]string,sms)
	for i,n := range gnames {
		if !sub || sms == 0 {
			gtypes[i] = gtypes[i][1:]
		}
		dispitems[i] = fmt.Sprintf(
`	%s func(%s)%s`,n,strings.Join(gtypes[i],", "),grtypes[i])
		if sub && i < sms {
			sdispitems[i] = fmt.Sprintf(
`	%s func(%s)%s`,n,strings.Join(gtypes[i][1:],", "),grtypes[i])
		}
	}
	gocode.WriteString(fmt.Sprintf(`
type %sDispatch struct {
%s
}
var %sLookup map[unsafe.Pointer]%sDispatch =
	map[unsafe.Pointer]%sDispatch{}
`,gname,strings.Join(dispitems,"\n"),gname,gname,gname))
	if sub && sms > 0 {
		gocode.WriteString(fmt.Sprintf(`
type %sSupermethods struct {
%s
}
`,gname,strings.Join(sdispitems,"\n")))
	}
	//To create (per method):
	cprotos.WriteString("\n\n")
	for i,m := range methods {
		//1. ObjC function prototypes for go exports
		_,_,ctps := w.cparamlist(m)
		cprotos.WriteString(fmt.Sprintf(
`%s %s%s(%s);
`,crtypes[i],gname,gnames[i],ctps))
		//2. Go callback registration functions
		gocode.WriteString(fmt.Sprintf(`
func (d %s) %sCallback(f func(%s)%s) {
	dispatch := %sLookup[d.Ptr()]
	dispatch.%s = f
	%sLookup[d.Ptr()] = dispatch
}
`,gname,gnames[i],strings.Join(gtypes[i],", "),grtypes[i],gname,gnames[i],gname))
		//3. Go exported callback function wrappers
		earglist := []string{"o unsafe.Pointer"}
		garglist := []string{}
		gargconv := []string{}
		if sub && sms > 0 {
			garglist = []string{"super"}
		}
		for j := 1; j < len(vnames[i]); j++ {
			earglist = append(earglist,vnames[i][j] + " " + getypes[i][j])
			var gt2 string
			if sub {
				gt2 = gtypes[i][j]
			} else {
				gt2 = gtypes[i][j-1]
			}
			if types.IsGoInterface(gt2) {
				gt2 = "Id"
			}
			if types.ShouldWrap(gt2) || gt2 == "Id" {
				garglist = append(garglist,fmt.Sprintf(
`a%d`,j))
				gargconv = append(gargconv,fmt.Sprintf(
`	a%d := %s{}; a%d.ptr = %s`,j,gt2,j,vnames[i][j]))
			} else {
				garglist = append(garglist,fmt.Sprintf(
`(%s)(%s)`,gt2,vnames[i][j]))
			}
		}
		retdecl := ""
		retname := ""
		retn := ""
		retnparen := ""
		crtype := ""
		if cgtypes[i] != "C.void" {
			retdecl = "var ret " + cgtypes[i] + "\n\t"
			retname = " ret"
			if cgtypes[i] == "unsafe.Pointer" {
				retn = "return unsafe.Pointer("
				crtype = " unsafe.Pointer"
				if types.ShouldWrap(m.Type.GoType()) {
					retnparen = ".Ptr()"
				}
			} else {
				retn = "return (" + cgtypes[i] + ")("
				crtype = " " + cgtypes[i]
			}
			retnparen = retnparen + ")"
		}
		sdispentries := make([]string,sms)
		for i,_ := range sdispentries {
			sdispentries[i] = fmt.Sprintf(
`		self.Super%s`,gnames[i])
		}
		sper := ""
		if sub && sms > 0 {
			sper = fmt.Sprintf(
`	self := (*%s)(o)
	super := %sSupermethods{
%s,
	}
	`,gname,gname,strings.Join(sdispentries,",\n"))
		}
		if len(gargconv) > 0 {
			retn = "\n	" + retn
		} else {
			retn = "	" + retn
		}
		goexports.WriteString(fmt.Sprintf(`
//export %s%s
func %s%s(%s)%s {
	%scb := %sLookup[o].%s
	if cb == nil { return%s }
%s%s%scb(%s)%s
}
`,gname,gnames[i],gname,gnames[i],strings.Join(earglist,", "),crtype,retdecl,gname,gnames[i],retname,sper,strings.Join(gargconv,"\n"),retn,strings.Join(garglist,", "),retnparen))
		//4. Go wrapper functions for superclass methods
		if !sub || i >= sms { continue } // for subclasses only
		grtype := m.Type.GoType()
		if grtype == "Void" {
			grtype = ""
		}
		if types.IsGoInterface(grtype) {
			grtype = "Id"
		}
		if grtype == "BOOL" {
			grtype = "bool"
		}
		if sub {
			gocode.WriteString(fmt.Sprintf(`
func (o %s) Super%s(%s) %s {
`,gname,gnames[i],strings.Join(earglist[1:],", "), grtype))
			ns,snames,tps,_ := w.gpntp(m)
			lparm := len(tps)-1
			if len(tps) > 0 && tps[lparm].Variadic {
				vn := ns[lparm]
				vn = vn[:len(vn)-1]
				ns[lparm] = vn
				gocode.WriteString(fmt.Sprintf(
`       var %s [%d]unsafe.Pointer
	for i,o := range %ss {
		%s[i] = o.Ptr()
	}
`,vn,w.Vaargs,vn,vn))
			}
			gocode.WriteString(`  ` + types.GoToC(dname + "_super_"+m.Name,ns,snames,m.Type,tps,false) + "\n}\n")
		}
	}
	w.cCode.WriteString(cprotos.String())
	w.cCode.WriteString(ccode.String())
	w.goTypes.WriteString(gotypes.String())
	w.goCode.WriteString(gocode.String())
	w.goExports.WriteString(goexports.String())
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
	ef,err := os.Create(path.Join(w.Package,"exports.go"))
	if err != nil {
		fmt.Printf("Error opening file %s\n%s\n",path.Join(w.Package,"exports.go"),err)
		os.Exit(-1)
	}
	fmt.Printf("Writing output to %s\n",path.Join(w.Package,"main.go"))
	pInterfaces := map[string]*Interface{}
	//Note: the following code eliminates duplicates, so it is acceptable
	//to have duplicate interfaces in 'toproc'
	for _,iface := range toproc {
		pInterfaces[iface] = w.Interfaces[iface]
	}
	//FIXME: sort pInterfaces
	for _,i := range pInterfaces {
		if i == nil {
			continue
		}
		w.processType(types.NewTypeFromString(i.GoName,""))
		if i.Name == "NSString" {
			w.StringHelpers()
		}
		if i.Name == "NSEnumerator" {
			w.EnumeratorHelpers()
		}
		gname := i.GoName
		if types.IsGoInterface(i.GoName) {
			gname = "Id"
		}
		fmt.Printf("Interface %s: %d properties, %d class methods, %d instance methods\n",
			i.Name, len(i.Properties), len(i.ClassMethods.Methods), len(i.InstanceMethods.Methods))

		w.goCode.WriteString(fmt.Sprintf(`
func %sAlloc() %s {
	ret := %s{}
	ret.ptr = unsafe.Pointer(C.%sAlloc())
	return ret
}
`,i.GoName,gname,gname,i.Name))

		if i.Name != "NSAutoreleasePool" {
			w.cCode.WriteString(fmt.Sprintf(`
void*
%sAlloc() {
	return [[%s alloc] autorelease];
}
`, i.Name, i.Name))
		} else {
		//who autoreleases the autorelease pools?
			w.cCode.WriteString(fmt.Sprintf(`
void*
%sAlloc() {
	return [%s alloc];
}
`, i.Name, i.Name))
		}

		//FIXME: sort properties
		for _,p := range i.Properties {
			if Debug {
				fmt.Printf("  property: %s (%s)\n", p.Name, p.Type.CType())
			}
		}
		for _,m := range i.ClassMethods.Methods {
			w.ProcessMethod(m)
		}
		for _,m := range i.InstanceMethods.Methods {
			w.ProcessMethod(m)
		}
		// add methods for Protocols that this interface implements
		for _,p := range i.Protocols {
			prot,ok := w.Protocols[p]
			if !ok {
				fmt.Printf("Failed to find protocol %s for interface %s\n",p,i.Name)
				os.Exit(-1)
			}
			for _,m := range prot.ClassMethods.Methods {
				w.ProcessMethodForClass(m,i.Name)
			}
			for _,m := range prot.InstanceMethods.Methods {
				w.ProcessMethodForClass(m,i.Name)
			}
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
	for n,p := range w.Delegates {
		w.ProcessDelegate(n,p)
	}
	for n,s := range w.Subclasses {
		w.ProcessSubclass(n,s)
	}
	fmt.Printf("%d functions\n", len(w.Functions))
	fmt.Printf("%d enums\n", len(w.NamedEnums) + len(w.AnonEnums))
	of.WriteString("package " + w.Package + "\n\n")

	of.WriteString(w.cgoFlags.String() + "\n")
	of.WriteString(w.cImports.String() + "\n")

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

	if len(w.Delegates) > 0 || len(w.Subclasses) > 0 {
		ef.WriteString("package " + w.Package + "\n\n")
		ef.WriteString(w.cgoFlags.String() + "\n")
		ef.WriteString(w.cImports.String() + "\n")
		ef.WriteString(`
*/
import "C"

import (
	"unsafe"
)
`)
		ef.WriteString(w.goExports.String())
		ef.Close()
	} else {
		ef.Close()
		err := os.Remove(path.Join(w.Package,"exports.go"))
		if err != nil {
			fmt.Printf("Error removing 'exports.go'. %s\n",err)
			os.Exit(-1)
		}
	}
}
