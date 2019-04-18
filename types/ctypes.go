package types

// Parsers for recognizing type names in C/Objective-C

func TypeName(s string, n *Node) (string, *Node) {
	return ChildOf(NewNode("TypeName"),SeqC(
		SpecifierQualifierList,
		Opt(AbstractDeclarator),
	))(s,n)
}

func AbstractDeclarator(s string, n *Node) (string, *Node) {
	return OneOf(SeqC(
			Opt(Pointer),
			Children(OneOrMore(DirectAbstractDeclarator))),
		Pointer,
	)(s,n)
}

func DirectAbstractDeclarator(s string, n *Node) (string, *Node) {
	return OneOf(
			ChildOf(NewNode("Parenthesized"),Parenthesized(AbstractDeclarator)),
			NodeNamed("Array",Bracketed(Opt(TypeQualifierList))),
			NodeNamed("Array",Bracketed(SeqC(Opt(TypeQualifierList),NodeNamed("Length",Regexp(`[\d]+|\*`))))), // NOTE: Does not allow arbitrary expressions
			NodeNamed("Array",Bracketed(SeqC(Word("static"),Opt(TypeQualifierList),NodeNamed("Length",Regexp(`[\d]+`))))), // NOTE: Does not allow arbitrary expressions
			NodeNamed("Array",Bracketed(SeqC(Opt(TypeQualifierList),Word("static"),NodeNamed("Length",Regexp(`[\d]+`))))), // NOTE: Does not allow arbitrary expressions
			ChildOf(NewNode("Function"),Parenthesized(Opt(ParameterList))),
	)(s,n)
}

func ParameterList(s string, n *Node) (string, *Node) {
	return SeqC(
		Opt(Children(OneOrMore(SeqC(ParameterDeclaration,Lit(","))))),
		ParameterDeclaration,
	)(s,n)
}

func ParameterDeclaration(s string, n *Node) (string, *Node) {
	return ChildOf(NewNode("ParameterDeclaration"),OneOf(
		NodeNamed("Ellipsis",Lit("...")),
		SeqC(DeclarationSpecifiers,Declarator),
		SeqC(DeclarationSpecifiers,Opt(AbstractDeclarator)),
	))(s,n)
}

func DeclarationSpecifiers(s string, n *Node) (string, *Node) {
	return OneOf(
		SeqC(StorageClassSpecifier,Opt(DeclarationSpecifiers)),
		SeqC(TypeSpecifier,Opt(DeclarationSpecifiers)),
		SeqC(StructOrUnionSpecifier,Opt(DeclarationSpecifiers)),
		SeqC(TypeQualifier,Opt(DeclarationSpecifiers)),
		SeqC(TypedefName,Opt(DeclarationSpecifiers)),
	//	SeqC(FunctionSpecifier,Opt(DeclarationSpecifiers)),
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
		SeqC(ZeroOrMore(Pointer), DirectDeclarator))(s,n)
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
	return SeqC(
		NodeNamed("Pointer",Lit("*")),
		Opt(TypeQualifierList),
		Opt(NullableAnnotation),
		Opt(Pointer),
	)(s,n)
}

func TypeQualifierList(s string, n *Node) (string, *Node) {
	return NodeNamed("TypeQualifierList",
		Children((OneOrMore(TypeQualifier))),
	)(s,n)
}

func SpecifierQualifierList(s string, n *Node) (string, *Node) {
	return NodeNamed("SpecifierQualifierList",
		OneOf(
			SeqC(TypeSpecifier,Opt(SpecifierQualifierList)),
			SeqC(StructOrUnionSpecifier,Opt(SpecifierQualifierList)),
			SeqC(TypedefName,Opt(SpecifierQualifierList)),
			SeqC(TypeQualifier,Opt(SpecifierQualifierList)),
		),
	)(s,n)
	//	OneOrMore(OneOf(TypeQualifier,TypeSpecifier)))(s,n)
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
		SeqC(StructOrUnion,Opt(Identifier),StructDeclarationList),
		NestC(StructOrUnion,Identifier),
	))(s,n)
}

func StructOrUnion(s string, n *Node) (string, *Node) {
	return OneOf(
		NodeNamed("Struct",Word("struct")),
		NodeNamed("Union",Word("union")))(s,n)
}

func StructDeclarationList(s string, n *Node) (string, *Node) {
	return NodeNamed("StructDeclarationList",OneOrMore(StructDeclaration))(s,n)
}

func StructDeclaration(s string, n *Node) (string, *Node) {
	return NodeNamed("StructDeclaration",Seq(
		SpecifierQualifierList,
		StructDeclaratorList,
		Lit(";"),
	))(s,n)
}

func StructDeclaratorList(s string, n *Node) (string, *Node) {
	return NodeNamed("StructDeclaratorList",Seq(
		Opt(OneOrMore(Seq(StructDeclarator,Lit(",")))),
		StructDeclarator,
	))(s,n)
}

func StructDeclarator(s string, n *Node) (string, *Node) {
	return NodeNamed("StructDeclarator",Declarator)(s,n)
}

func Generic(s string, n *Node) (string, *Node) {
	return NodeNamed("Generic",TypeName)(s,n)
}

func GenericList(s string, n *Node) (string, *Node) {
	return OneOf(
		SeqC(Generic,Lit(","),GenericList),
		Generic,
	)(s,n)
}

func TypedefName(s string, n *Node) (string, *Node) {
	return NodeNamed("TypedefName",OneOf(
		SeqC(NodeNamed("TypedefName",Identifier),AngBracketed(GenericList)),
		Identifier,
	))(s,n)
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

