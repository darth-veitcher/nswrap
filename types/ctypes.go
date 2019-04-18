package types

// Parsers for recognizing type names in C/Objective-C

func TypeName(s string, n *Node) (string, *Node) {
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
	return NodeNamed("DirectDeclarator",
		OneOf(
			Identifier,
			Parenthesized(Declarator),
			// INCOMPLETE
		),
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
		//StructOrUnionSpecifier,
		//EnumSpecifier,
		//TypedefName,
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
	return NodeNamed("StructOrUnionSpecifier",OneOf(
//		Seq(StructOrUnion,Opt(Identifier),StructDeclarationList),
		Nest(StructOrUnion,Identifier),
	))(s,n)
}

func StructOrUnion(s string, n *Node) (string, *Node) {
	return OneOf(
		NodeNamed("Struct",Word("struct")),
		NodeNamed("Union",Word("union")))(s,n)
}

func Generic(s string, n *Node) (string, *Node) {
	return NodeNamed("Generic",TypeName)(s,n)
}

func GenericList(s string, n *Node) (string, *Node) {
	return OneOf(
		Seq(Generic,Lit(","),GenericList),
		Generic,
	)(s,n)
}

func BareTypedefName(s string, n *Node) (string, *Node) {
	return NodeNamed("TypedefName",Identifier)(s,n)
}

func TypedefName(s string, n *Node) (string, *Node) {
	return OneOf(
		Seq(BareTypedefName, AngBracketed(GenericList)),
		Seq(BareTypedefName, NullableAnnotation),
		BareTypedefName,
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

