package types

/* Parsers for recognizing type names in C/Objective-C

type-name:
	specifier-qualifier-list abstract-declarator<opt>
abstract-declarator:
	pointer
	pointer<opt> direct-abstract-declarator
direct-abstract-declarator:
	( abstract-declarator )
	direct-abstract-declarator<opt> [ type-qualifier-list<opt> assignment-expression<opt> ]
	direct-abstract-declarator<opt> [ static type-qualifier-list<opt> assignment-expression ]
	direct-abstract-declarator<opt> [ type-qualifier-list static assignment-expression ]
	direct-abstract-declarator<opt> [ * ]
	direct-abstract-declarator<opt> ( parameter-type-list<opt> )
pointer:
	* type-qualifier-list<opt>
	* type-qualifier-list<opt> pointer
parameter-type-list:
	parameter-list
	parameter-list , ...
parameter-list:
	parameter-declaration
	parameter-list , parameter-declaration
parameter-declaration:
	declaration-specifiers declarator
	declaration-specifiers abstract-declarator<opt>
type-qualifier-list:
	type-qualifier
	type-qualifier-list type-qualifier
specifier-qualifier-list:
	type-specifier specifier-qualifier-list<opt>
	type-qualifier specifier-qualifier-list<opt>
type-specifier:
	void
	char
	short
	int
	long
	float
	double
	signed
	unsigned
	_Bool
	_Complex
	struct-or-union-specifier
	enum-specifier
	typedef-name
type-qualifier:
	const
	restrict
	volatile
struct-or-union-specifier:
	// DON'T DO struct-or-union identifier<opt> { struct-declaration-list }
	struct-or-union identifier
struct-or-union:
	struct
	union
struct-declaration-list:
	struct-declaration
	struct-declaration-list struct-declaration
struct-declaration:
	specifier-qualifier-list struct-declarator-list ;
struct-declarator-list:
	struct-declarator
	struct-declarator-list , struct-declarator
struct-declarator:
	declarator
	declarator<opt>: constant-expression
identifier:
	identifier-non-digit
	identifier identifier-nondigit
	identifier digit
identifier-nondigit:
	nondigit
	universal-character-name
nondigit:
	_ [a-zA-Z]
digit:
	[0-9]
*/

var TypeName func(s string, n *Node) (string, *Node)

func init() {
	cache := map[string]*Node{}
	TypeName = func(s string, n *Node) (string, *Node) {
		if n2,ok := cache[s]; ok {
			return "",n2
		}
		s2,n2 := _TypeName(s,n)
		if s2 == "" {
			cache[s] = n2
		}
		return s2,n2
	}
	TypeName = _TypeName
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
		Id,
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

func Id(s string, n *Node) (string, *Node) {
	return Seq(
		NodeNamed("Id",Lit("id")),
		Opt(TypeQualifierList),
		Opt(NullableAnnotation),
	)(s,n)
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
		//StructOrUnionSpecifier,
		EnumSpecifier,
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

