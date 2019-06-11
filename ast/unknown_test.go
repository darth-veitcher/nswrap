package ast

import (
	"fmt"
	"testing"
)

func TestUnknown(t *testing.T) {
	i := 1
	runNodeTest(t,
		Parse(`SomeNode 0x7faa18a445d8 <line:66:45> "asdf" aoeu`),
		testNode{&Unknown{
			Addr:       0x7faa18a445d8,
			Name:       "SomeNode",
			Pos:        NewPositionFromString("line:66:45"),
			Content:    ` "asdf" aoeu`,
			ChildNodes: []Node{},
		},
			0x7faa18a445d8,
			NewPositionFromString("line:66:45"),
			[]Node{},
		},
		&i,
	)
	runNodeTest(t,
		Parse(`AlignedAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`AllocSizeAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`AlwaysInlineAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ArcWeakrefUnavailableAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ArraySubscriptExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`AsmLabelAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`AttributedType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`AvailabilityAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`BinaryOperator 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`BlockCommandComment 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`BlockPointerType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`BreakStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`BuiltinType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CFAuditedTransferAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CFConsumedAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CFReturnsNotRetainedAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CFReturnsRetainedAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CStyleCastExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CallExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CaseStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CharacterLiteral 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CompoundAssignOperator 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CompoundLiteralExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`CompoundStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ConditionalOperator 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ConstantArrayType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ContinueStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ConvertVectorExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`DecayedType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`DeclRefExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`DeclRefExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`DeclStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`DefaultStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`DeprecatedAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`DisableTailCallsAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`DoStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ElaboratedType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`EmptyDecl 0x7faa18a445d8 <line:66:45> col:12 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`Enum 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`EnumConstantDecl 0x7faa18a445d8 <line:66:45> col: 32 invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`EnumDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`EnumType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`EnumExtensibilityAttr 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`Field 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`FieldDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`FlagEnumAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`FloatingLiteral 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ForStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`FormatArgAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`FormatAttr 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`FullComment 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`FunctionDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`FunctionProtoType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`GCCAsmStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`GotoStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`HTMLEndTagComment 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`HTMLStartTagComment 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`IBActionAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`IBOutletAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`IfStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ImplicitCastExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ImplicitValueInitExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`IncompleteArrayType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`IndirectFieldDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`InitListExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`InlineCommandComment 0x7faa18a445d8 <line:66:45>`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`IntegerLiteral 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`LabelStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`MallocAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`MaxFieldAlignmentAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`MayAliasAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`MemberExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`MinVectorWidthAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ModeAttr 0x7faa18a445d8 <line:66:45>`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NoDebugAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NoEscapeAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NoInlineAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NoThrowAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NonNullAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NotTailCalledAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NSConsumedAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NSConsumesSelfAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NSErrorDomainAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`NSReturnsRetainedAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCBoolLiteralExpr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCBoxableAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCBridgeAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCBridgeMutableAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCBridgeRelatedAttr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCCategoryDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCDesignatedInitializerAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCExceptionAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCExplicitProtocolImplAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCIndependentClassAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCInterface 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCInterfaceDecl 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCInterfaceType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCIvarDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCMessageExpr invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCMethod 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCMethodDecl 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCObjectPointerType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCObjectType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCPropertyDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCProtocol 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCProtocolDecl invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCRequiresSuperAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCReturnsInnerPointerAttr 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCRootClassAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ObjCTypeParamDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`OffsetOfExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`PackedAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ParagraphComment 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ParamCommandComment 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ParenExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ParenType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ParmVarDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`PointerType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`PredefinedExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`PureAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`QualType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`Record 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`RecordDecl 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`RecordType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`RestrictAttr 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ReturnStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ReturnsTwiceAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`SentinelAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`ShuffleVectorExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`StmtExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`StringLiteral 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`SwiftBridgedTypedefAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`SwiftErrorAttr 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`SwiftNameAttr 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`SwiftNewtypeAttr 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`SwiftPrivateAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`SwitchStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`TargetAttr 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`TextComment 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`TranslationUnitDecl 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`TransparentUnionAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`Typedef 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`TypedefDecl 0x7faa18a445d8 <line:66:45> invalid invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`TypedefType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`UnaryExprOrTypeTraitExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`UnaryOperator 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`UnavailableAttr 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`UnusedAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`UsedAttr 0x7faa18a445d8 invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`VAArgExpr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`VarDecl 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`VectorType 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`VerbatimBlockComment 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`VerbatimBlockLineComment 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`VerbatimLineComment 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`VisibilityAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`WarnUnusedResultAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`WeakAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`WeakImportAttr 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	runNodeTest(t,
		Parse(`WhileStmt 0x7faa18a445d8 <line:66:45> invalid`),
		testNode{&Unknown{}, 0, NewPositionFromString(""), nil},
		&i,
	)
	t.Run(fmt.Sprintf("Example%d", i), func(t *testing.T) {
		node := Parse(`NullStmt`)
		if node != nil {
			t.Errorf("Parse(NullStmt) did not return nil\n")
		}
	})
}
