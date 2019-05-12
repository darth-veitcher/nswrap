package types

// A parser to recognize type names in C/Objective-C

import (
	"regexp"
)

var TypeName func(s string, n *Node) (string, *Node)

func init() {
	instancename := regexp.MustCompile("instancename")
	instancetype := regexp.MustCompile("instancetype")
	cacheable := func(s string) bool {
		return !instancetype.MatchString(s) && !instancename.MatchString(s)
	}

	//memoize the TypeName function for performance
	cache := map[string]*Node{}
	TypeName = func(s string, n *Node) (string, *Node) {
		if n2,ok := cache[s]; ok {
			return "",n2
		}
		s2,n2 := _TypeName(s,n)
		if s2 == "" && cacheable(s) {
			cache[s] = n2
		}
		return s2,n2
	}
	//for debug purposes, the following line can be uncommented, which will
	//memoization memoization
	//TypeName = _TypeName
}

func _TypeName(s string, n *Node) (string, *Node) {
	return ChildOf(NewNode("TypeName"),Seq(
		SpecifierQualifierList,
		Opt(AbstractDeclarator),
	))(s,n)
}

func AbstractDeclarator(s string, n *Node) (string, *Node) {
	return OneOf(Seq(
			Opt(Pointer),
			OneOrMore(DirectAbstractDeclarator)),
		Pointer,
		Block,
	)(s,n)
}

func ParenAbstractDeclarator(s string, n *Node) (string, *Node) {
	return ChildOf(NewNode("Parenthesized"),
			Parenthesized(AbstractDeclarator),
	)(s,n)
}

func ArrayDeclarator(s string, n *Node) (string, *Node) {
	return OneOf(
		ChildOf(NewNode("Array"),
			Bracketed(Opt(TypeQualifierList))),

		// NOTE: Parser does not allow arbitrary 'length' expressions
		ChildOf(NewNode("Array"),
			Bracketed(Seq(
				Opt(TypeQualifierList),
				NodeNamed("Length",Regexp(`[\d]+|\*`))))),

		ChildOf(NewNode("Array"),
			Bracketed(Seq(
				Word("static"),
				Opt(TypeQualifierList),
				NodeNamed("Length",Regexp(`[\d]+`))))),

		ChildOf(NewNode("Array"),
			Bracketed(Seq(
				Opt(TypeQualifierList),
				Word("static"),
				NodeNamed("Length",Regexp(`[\d]+`))))),
	)(s,n)
}

func FunctionDeclarator(s string, n *Node) (string, *Node) {
	return ChildOf(NewNode("Function"),
			Parenthesized(Opt(ParameterList)),
	)(s,n)
}

func DirectAbstractDeclarator(s string, n *Node) (string, *Node) {
	return OneOf(
		ParenAbstractDeclarator,
		ArrayDeclarator,
		FunctionDeclarator,
	)(s,n)
}

func ParameterList(s string, n *Node) (string, *Node) {
	return Seq(
		Opt(OneOrMore(Seq(ParameterDeclaration,Lit(",")))),
		ParameterDeclaration,
	)(s,n)
}

func ParameterDeclaration(s string, n *Node) (string, *Node) {
	return ChildOf(NewNode("ParameterDeclaration"),OneOf(
		NodeNamed("Ellipsis",Lit("...")),
		Seq(DeclarationSpecifiers,Declarator),
		Seq(DeclarationSpecifiers,Opt(AbstractDeclarator)),
	))(s,n)
}

func DeclarationSpecifiers(s string, n *Node) (string, *Node) {
	return OneOf(
		Seq(StorageClassSpecifier,Opt(DeclarationSpecifiers)),
		Seq(TypeSpecifier,Opt(DeclarationSpecifiers)),
		Seq(StructOrUnionSpecifier,Opt(DeclarationSpecifiers)),
		Seq(TypeQualifier,Opt(DeclarationSpecifiers)),
		Seq(TypedefName,Opt(DeclarationSpecifiers)),
	//	Seq(FunctionSpecifier,Opt(DeclarationSpecifiers)),
	)(s,n)
}

func StorageClassSpecifier(s string, n *Node) (string, *Node) {
	return NodeNamed("StorageClassSpecifier",OneOf(
		Word("typedef"),
		Word("extern"),
		Word("static"),
		Word("auto"),
		Word("register"),
	))(s,n)
}

func Declarator(s string, n *Node) (string, *Node) {
	return ChildOf(NewNode("Declarator"),
		Seq(ZeroOrMore(Pointer), DirectDeclarator))(s,n)
}

func DirectDeclarator(s string, n *Node) (string, *Node) {
	return OneOf(
			Identifier,
			Parenthesized(Declarator),
			// INCOMPLETE
	)(s,n)
}

func NullableAnnotation(s string, n *Node) (string, *Node) {
	return NodeNamed("NullableAnnotation",OneOf(
		Word("_Nullable"),
		Word("_Nonnull"),
		Word("_Null_unspecified"),
	))(s,n)
}

func Pointer(s string, n *Node) (string, *Node) {
	return Seq(
		NodeNamed("Pointer",Lit("*")),
		Opt(TypeQualifierList),
		Opt(NullableAnnotation),
		Opt(Pointer),
	)(s,n)
}

//FIXME: not sure how correct this is...
func Block(s string, n *Node) (string, *Node) {
	return Seq(
		NodeNamed("Block",Lit("^")),
		Opt(NullableAnnotation),
	)(s,n)
}

func TypeQualifierList(s string, n *Node) (string, *Node) {
	return OneOrMore(TypeQualifier)(s,n)
}

func SpecifierQualifierList(s string, n *Node) (string, *Node) {
	return NodeNamed("SpecifierQualifierList",
		OneOf(
			Seq(TypeSpecifier,Opt(SpecifierQualifierList)),
			Seq(StructOrUnionSpecifier,Opt(SpecifierQualifierList)),
			Seq(TypedefName,Opt(SpecifierQualifierList)),
			Seq(TypeQualifier,Opt(SpecifierQualifierList)),
		),
	)(s,n)
}

func TypeSpecifier(s string, n *Node) (string, *Node) {
	return NodeNamed("TypeSpecifier",OneOf(
		Word("void"),
		Word("char"),
		Word("short"),
		Word("int"),
		Word("long"),
		Word("float"),
		Word("double"),
		Word("signed"),
		Word("unsigned"),
		Word("_Bool"),
		Word("_Complex"),
		EnumSpecifier,
	))(s,n)
}

func TypeQualifier(s string, n *Node) (string, *Node) {
	return NodeNamed("TypeQualifier",OneOf(
		Word("const"),
		Word("restrict"),
		Word("volatile"),
	))(s,n)
}

func StructOrUnionSpecifier(s string, n *Node) (string, *Node) {
	return OneOf(
//		Seq(StructOrUnion,Opt(Identifier),StructDeclarationList),
		Nest(StructOrUnion,Identifier),
	)(s,n)
}

func StructOrUnion(s string, n *Node) (string, *Node) {
	return OneOf(
		NodeNamed("Struct",Word("struct")),
		NodeNamed("Union",Word("union")))(s,n)
}

func EnumSpecifier(s string, n *Node) (string, *Node) {
	return Nest(
		NodeNamed("Enum",Word("enum")),Identifier)(s,n)
}

func Generic(s string, n *Node) (string, *Node) {
	return NodeNamed("Generic",TypeName)(s,n)
}

func GenericList(s string, n *Node) (string, *Node) {
	return ChildOf(NewNode("GenericList"),Seq(
		Generic,
		ZeroOrMore(Seq(Lit(","),Generic)),
	))(s,n)
}

func BareTypedefName(s string, n *Node) (string, *Node) {
	return NodeNamed("TypedefName",Identifier)(s,n)
}

func TypedefName(s string, n *Node) (string, *Node) {
	return Seq(
		Opt(NodeNamed("KindQualifier",Lit("__kindof"))),
		BareTypedefName,
		Opt(AngBracketed(GenericList)),
		Opt(NullableAnnotation),
	)(s,n)
}

func Identifier(s string, n *Node) (string, *Node) {
	s2,n2 := NodeNamed("Identifier",
		Regexp(`[_a-zA-Z][_0-9a-zA-Z]*`))(s,n)
	if n2 == nil {
		return s,nil
	}
	if reservedwords.MatchString(n2.Content) {
		dbg("Identifier '%s' contains reserved word\n",n2.Content)
		return s,nil
	}
	return s2,n2
}

