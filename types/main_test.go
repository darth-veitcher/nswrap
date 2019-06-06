package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	i := 1
	chk := func(input string,n *Node,expected string) {
		t.Run(fmt.Sprintf("Parse%d",i), func(t *testing.T) {
			if n.String() != expected {
				t.Errorf("Mismatch parsing %s -- Got:\n%s\nExpected:\n%s\n",input,n.String(),expected)
			}
		})
		i++
	}

	runParseTest := func(input, expected string) {
		n,err := Parse(input)
		if err != nil {
			t.Errorf("Parse error parsing %s\n",input)
		}
		chk(input,n,expected)
	}

	runSigTest := func(input, expected string) {
		rem,n := MethodSignature(input,NewNode("AST"))
		if rem != "" {
			t.Errorf("Parse error parsing %s. Remainder = %s\n",input,rem)
		}
		chk(input,n,expected)
	}

	runParseTest(`int`,
`<TypeName> ''
-<TypeSpecifier> 'int'
`)
	runParseTest(`int`,
`<TypeName> ''
-<TypeSpecifier> 'int'
`)
	runParseTest(`char*`,
`<TypeName> ''
-<TypeSpecifier> 'char'
-<Pointer> '*'
`)
	runParseTest(`struct Str`,
`<TypeName> ''
-<Struct> 'struct'
--<Identifier> 'Str'
`)
	runParseTest(`uint`,
`<TypeName> ''
-<TypedefName> 'uint'
`)
	runParseTest(`uind`,
`<TypeName> ''
-<TypedefName> 'uind'
`)
	runParseTest(`void ()`,
`<TypeName> ''
-<TypeSpecifier> 'void'
-<Function> ''
`)
	runParseTest(`void (*)(struct NSRange)`,
`<TypeName> ''
-<TypeSpecifier> 'void'
-<Parenthesized> ''
--<Pointer> '*'
-<Function> ''
--<ParameterDeclaration> ''
---<Struct> 'struct'
----<Identifier> 'NSRange'
`)
	runParseTest(`void (* _Nullable)(int)`,
`<TypeName> ''
-<TypeSpecifier> 'void'
-<Parenthesized> ''
--<Pointer> '*'
--<NullableAnnotation> '_Nullable'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeSpecifier> 'int'
`)
	runParseTest(`NSRange[3]`,
`<TypeName> ''
-<TypedefName> 'NSRange'
-<Array> ''
--<Length> '3'
`)
	runParseTest(`struct _NSRange[3]`,
`<TypeName> ''
-<Struct> 'struct'
--<Identifier> '_NSRange'
-<Array> ''
--<Length> '3'
`)
	runParseTest(`struct _NSRange (*)(int, char[])`,
`<TypeName> ''
-<Struct> 'struct'
--<Identifier> '_NSRange'
-<Parenthesized> ''
--<Pointer> '*'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeSpecifier> 'int'
--<ParameterDeclaration> ''
---<TypeSpecifier> 'char'
---<Array> ''
`)
	runParseTest(`struct _NSRange (*)(int, char[],NSRange)`,
`<TypeName> ''
-<Struct> 'struct'
--<Identifier> '_NSRange'
-<Parenthesized> ''
--<Pointer> '*'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeSpecifier> 'int'
--<ParameterDeclaration> ''
---<TypeSpecifier> 'char'
---<Array> ''
--<ParameterDeclaration> ''
---<TypedefName> 'NSRange'
`)
	runParseTest(`struct _NSRange (*)(int, char[],struct _NSRange)`,
`<TypeName> ''
-<Struct> 'struct'
--<Identifier> '_NSRange'
-<Parenthesized> ''
--<Pointer> '*'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeSpecifier> 'int'
--<ParameterDeclaration> ''
---<TypeSpecifier> 'char'
---<Array> ''
--<ParameterDeclaration> ''
---<Struct> 'struct'
----<Identifier> '_NSRange'
`)
	runParseTest(`mytype`,
`<TypeName> ''
-<TypedefName> 'mytype'
`)
	runParseTest(`const int`,
`<TypeName> ''
-<TypeQualifier> 'const'
-<TypeSpecifier> 'int'
`)
	runParseTest(`int[*]`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Array> ''
--<Length> '*'
`)
	runParseTest(`int[3][5]`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Array> ''
--<Length> '3'
-<Array> ''
--<Length> '5'
`)
	runParseTest(`int *`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Pointer> '*'
`)
	runParseTest(`int * _Nonnull`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Pointer> '*'
-<NullableAnnotation> '_Nonnull'
`)
	runParseTest(`int **`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Pointer> '*'
-<Pointer> '*'
`)
	runParseTest(`int *[3]`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Pointer> '*'
-<Array> ''
--<Length> '3'
`)
	runParseTest(`int (*)[3]`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Parenthesized> ''
--<Pointer> '*'
-<Array> ''
--<Length> '3'
`)
	runParseTest(`int (*)[*]`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Parenthesized> ''
--<Pointer> '*'
-<Array> ''
--<Length> '*'
`)
	runParseTest(`int *()`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Pointer> '*'
-<Function> ''
`)
	runParseTest(`int (*)(void)`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Parenthesized> ''
--<Pointer> '*'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeSpecifier> 'void'
`)
	runParseTest(`int (*)(int, unsigned int)`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Parenthesized> ''
--<Pointer> '*'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeSpecifier> 'int'
--<ParameterDeclaration> ''
---<TypeSpecifier> 'unsigned'
---<TypeSpecifier> 'int'
`)
	runParseTest(`int (* _Nullable)(int, unsigned int)`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Parenthesized> ''
--<Pointer> '*'
--<NullableAnnotation> '_Nullable'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeSpecifier> 'int'
--<ParameterDeclaration> ''
---<TypeSpecifier> 'unsigned'
---<TypeSpecifier> 'int'
`)
	runParseTest(`int (*const [])(unsigned int, ...)`,
`<TypeName> ''
-<TypeSpecifier> 'int'
-<Parenthesized> ''
--<Pointer> '*'
--<TypeQualifier> 'const'
--<Array> ''
-<Function> ''
--<ParameterDeclaration> ''
---<TypeSpecifier> 'unsigned'
---<TypeSpecifier> 'int'
--<ParameterDeclaration> ''
---<Ellipsis> '...'
`)
	runParseTest(`BOOL (* _Nullable)()`,
`<TypeName> ''
-<TypedefName> 'BOOL'
-<Parenthesized> ''
--<Pointer> '*'
--<NullableAnnotation> '_Nullable'
-<Function> ''
`)
	runParseTest(`BOOL (* _Nullable)(const void * _Nonnull, const void * _Nonnull)`,
`<TypeName> ''
-<TypedefName> 'BOOL'
-<Parenthesized> ''
--<Pointer> '*'
--<NullableAnnotation> '_Nullable'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeQualifier> 'const'
---<TypeSpecifier> 'void'
---<Pointer> '*'
---<NullableAnnotation> '_Nonnull'
--<ParameterDeclaration> ''
---<TypeQualifier> 'const'
---<TypeSpecifier> 'void'
---<Pointer> '*'
---<NullableAnnotation> '_Nonnull'
`)
	runParseTest(`BOOL (* _Nullable)(const void * _Nonnull, const void * _Nonnull, int (* _Nullable)(const void * _Nonnull))`,
`<TypeName> ''
-<TypedefName> 'BOOL'
-<Parenthesized> ''
--<Pointer> '*'
--<NullableAnnotation> '_Nullable'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeQualifier> 'const'
---<TypeSpecifier> 'void'
---<Pointer> '*'
---<NullableAnnotation> '_Nonnull'
--<ParameterDeclaration> ''
---<TypeQualifier> 'const'
---<TypeSpecifier> 'void'
---<Pointer> '*'
---<NullableAnnotation> '_Nonnull'
--<ParameterDeclaration> ''
---<TypeSpecifier> 'int'
---<Parenthesized> ''
----<Pointer> '*'
----<NullableAnnotation> '_Nullable'
---<Function> ''
----<ParameterDeclaration> ''
-----<TypeQualifier> 'const'
-----<TypeSpecifier> 'void'
-----<Pointer> '*'
-----<NullableAnnotation> '_Nonnull'
`)
	runParseTest(`BOOL (* _Nullable)(const void * _Nonnull, const void * _Nonnull, NSUInteger (* _Nullable)(const void * _Nonnull))`,
`<TypeName> ''
-<TypedefName> 'BOOL'
-<Parenthesized> ''
--<Pointer> '*'
--<NullableAnnotation> '_Nullable'
-<Function> ''
--<ParameterDeclaration> ''
---<TypeQualifier> 'const'
---<TypeSpecifier> 'void'
---<Pointer> '*'
---<NullableAnnotation> '_Nonnull'
--<ParameterDeclaration> ''
---<TypeQualifier> 'const'
---<TypeSpecifier> 'void'
---<Pointer> '*'
---<NullableAnnotation> '_Nonnull'
--<ParameterDeclaration> ''
---<TypedefName> 'NSUInteger'
---<Parenthesized> ''
----<Pointer> '*'
----<NullableAnnotation> '_Nullable'
---<Function> ''
----<ParameterDeclaration> ''
-----<TypeQualifier> 'const'
-----<TypeSpecifier> 'void'
-----<Pointer> '*'
-----<NullableAnnotation> '_Nonnull'
`)
	runParseTest(`NSEnumerator<ObjectType> * _Nonnull`,
`<TypeName> ''
-<TypedefName> 'NSEnumerator'
-<GenericList> ''
--<Generic> ''
---<TypedefName> 'ObjectType'
-<Pointer> '*'
-<NullableAnnotation> '_Nonnull'
`)
	runParseTest(`NSArray<NSString *> * _Nonnull`,
`<TypeName> ''
-<TypedefName> 'NSArray'
-<GenericList> ''
--<Generic> ''
---<TypedefName> 'NSString'
---<Pointer> '*'
-<Pointer> '*'
-<NullableAnnotation> '_Nonnull'
`)
	runParseTest(`id  _Nonnull (^ _Nonnull)(id _Nullable, NSArray<NSExpression *> * _Nonnull, NSMutableDictionary * _Nullable)`,
`<TypeName> ''
-<TypedefName> 'id'
-<NullableAnnotation> '_Nonnull'
-<Parenthesized> ''
--<Block> '^'
--<NullableAnnotation> '_Nonnull'
-<Function> ''
--<ParameterDeclaration> ''
---<TypedefName> 'id'
---<NullableAnnotation> '_Nullable'
--<ParameterDeclaration> ''
---<TypedefName> 'NSArray'
---<GenericList> ''
----<Generic> ''
-----<TypedefName> 'NSExpression'
-----<Pointer> '*'
---<Pointer> '*'
---<NullableAnnotation> '_Nonnull'
--<ParameterDeclaration> ''
---<TypedefName> 'NSMutableDictionary'
---<Pointer> '*'
---<NullableAnnotation> '_Nullable'
`)
	runSigTest(`(void)pressed`,
`<MethodSignature> ''
-<TypeName> ''
--<TypeSpecifier> 'void'
-<Identifier> 'pressed'
`)
	runSigTest(`(void)performSelector:(SEL)aSelector target:(id)target argument:(id)arg order:(NSUInteger)order modes:(NSArray<NSRunLoopMode> *)modes`,
`<MethodSignature> ''
-<TypeName> ''
--<TypeSpecifier> 'void'
-<Identifier> 'performSelector'
-<MethodParameter> ''
--<TypeName> ''
---<TypedefName> 'SEL'
--<Identifier> 'aSelector'
-<MethodParameter> ''
--<Identifier> 'target'
--<TypeName> ''
---<TypedefName> 'id'
--<Identifier> 'target'
-<MethodParameter> ''
--<Identifier> 'argument'
--<TypeName> ''
---<TypedefName> 'id'
--<Identifier> 'arg'
-<MethodParameter> ''
--<Identifier> 'order'
--<TypeName> ''
---<TypedefName> 'NSUInteger'
--<Identifier> 'order'
-<MethodParameter> ''
--<Identifier> 'modes'
--<TypeName> ''
---<TypedefName> 'NSArray'
---<GenericList> ''
----<Generic> ''
-----<TypedefName> 'NSRunLoopMode'
---<Pointer> '*'
--<Identifier> 'modes'
`)
}

func TestFuncs(t *testing.T) {
	i := 1
	var n *Node
	var err error
	f := func(input string, actual, expected interface{}) {
		t.Run(fmt.Sprintf("TestFunc%d",i),func (t *testing.T) {
			if err != nil {
				t.Errorf("Error parsing %s\n",input)
			}
			if actual == nil && expected == nil { return }
			if !reflect.DeepEqual(actual,expected) {
				t.Errorf("Test failed for %s\n",input)
			}
		})
		i++
	}
	str := "int*"
	n,err = Parse(str)
	f(str,n.HasFunc(),false)
	str = "int (*)(void)"
	n,err = Parse(str)
	f(str,n.HasFunc(),true)
	n,err = nil,nil
	f("",n.HasFunc(),false)
	f("",n.IsId(),false)
	f("",n.IsFunction(),false)
	f("",n.IsInstancetype(),false)
	f("",n.IsArray(),false)
	f("",n.IsStruct(),false)
	if n.BaseType() != nil {
		t.Errorf("BaseType() for nil node is not nil\n")
	}
	if n.PointsTo() != nil {
		t.Errorf("PointsTo() for nil node is not nil\n")
	}
	if n.ReturnType() != nil {
		t.Errorf("ReturnType() for nil node is not nil\n")
	}
	f("",n.String(),"")
	f("",n.Qualifiers(),"")
	f("",n.Annotations(),"")

	str = "int"
	n,err = Parse(str)
	f(str,n.isIndirect("Pointer"),false)
	if n.ReturnType() != nil {
		t.Errorf("Return type for non-function is not nil\n")
	}

	str = "int*"
	n,err = Parse(str)
	f(str,n.isIndirect("Pointer"),true)
	f(str,n.isIndirect("Array"),false)

	str = "int* _Nullable"
	n,err = Parse(str)
	f(str,n.isIndirect("Pointer"),true)

	str = "int* _Nonnull"
	n,err = Parse(str)
	f(str,n.isIndirect("Pointer"),true)
	f(str,n.CType(),"int* _Nonnull")
	f(str,n.CTypeSimplified(),"int*")
	f(str,n.Annotations(),"_Nonnull")

	n,err = Parse("int[]")
	n2,err2 := Parse("int")
	if err2 != nil {
		t.Errorf("Cannot parse int")
	}
	f(str,n.isIndirect("Pointer"),false)
	f(str,n.isIndirect("Array"),true)
	f(str,n.ArrayOf(),n2)
	f(str,n.CType(),"int[]")
	f(str,n.CTypeSimplified(),"int[]")

	str = "const int*[]"
	n,err = Parse(str)
	f(str,n.isIndirect("Array"),true)
	f(str,n.Qualifiers(),"const")

	str = "int (*)[]"
	n,err = Parse(str)
	f(str,n.isIndirect("Array"),true)
	f(str,n.isIndirect("Pointer"),true)

	str = "int (* _Nullable * _Nonnull)[]"
	n,err = Parse(str)
	f(str,n.isIndirect("Pointer"),true)

	str = "int (*)(void)"
	n,err = Parse(str)
	n2,err2 = Parse("int(void)")
	if err2 != nil {
		t.Errorf("Failed to parse int(void)")
	}
	f(str,n.PointsTo(),n2)
	f(str,n.IsPointer(),true)
	f(str,n.IsFunction(),false)
	f(str,n.CType(),"int (*)(void)")

	str = "int *(void)"
	n,err = Parse(str)
	f(str,n.IsFunction(),true)
	n2,_ = Parse("int*")
	f(str,n.ReturnType(),n2)

	str = "__kindof NSRange"
	n,err = Parse(str)
	f(str,n.IsStruct(),false)
	f(str,n.IsInstancetype(),false)

	str = "struct NSRange"
	n,err = Parse(str)
	f(str,n.IsStruct(),true)
	f(str,n.IsFunction(),false)

	str = "id"
	n,err = Parse(str)
	f(str,n.IsId(),true)
	f(str,n.IsPointer(),true)

	str = "int * _Nullable * _Nonnull"
	n,err = Parse(str)
	n2,_ = Parse("int")
	f(str,n.BaseType(),n2)

	str = "__kindof id"
	n,err = Parse(str)
	f(str,n.IsId(),true)
	f(str,n.IsPointer(),true)


	str = "NSArray<ObjectType> *"
	n,err = Parse(str)
	f(str,n.CType(),"NSArray <ObjectType>*")
	f(str,n.CTypeSimplified(),"NSArray*")

	str = "NSRange *(int, NSRange *)"
	n,err = Parse(str)
	f(str,n.renameTypedefs("NSRange","struct NSRange"),true)
	f(str,n.CType(),"struct NSRange* (int, struct NSRange*)")

	str = "int*%&#$"
	n,err = Parse(str)
	if err == nil {
		fmt.Printf("%s\n",n.String())
		t.Errorf("Parse should have failed for %s\n",str)
	}

	var r func(i int) *Node
	r = func(i int) *Node {
		n := &Node{}
		if i > 0 {
			n.Children = []*Node{ r(i-1) }
		}
		return n
	}
	n = r(150)
	str = n.String()
	if str[len(str)-8:] != "too deep" {
		fmt.Printf(str[len(str)-8:])
		t.Errorf("Deep recursion did not produce an error from (*Node)String()\n")
	}
}

func ExampleDebug() {
	Debug = true
	dbg("one %s two %d","string",15)
	// Output: one string two 15
	Debug = false
}
