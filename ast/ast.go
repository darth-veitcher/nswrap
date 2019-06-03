// Package ast parses the clang AST output into AST structures.
package ast

import (
	"fmt"
	"strconv"
	"strings"

	"git.wow.st/gmp/nswrap/util"
)

var TrackPositions bool = false

// Node represents any node in the AST.
type Node interface {
	Address() Address
	Children() []Node
	AddChild(node Node)
	Position() Position
}

// Address contains the memory address (originally outputted as a hexadecimal
// string) from the clang AST. The address are not predictable between run and
// are only useful for identifying nodes in a single AST.
//
// The Address is used like a primary key when storing the tree as a flat
// structure.
type Address uint64

// ParseAddress returns the integer representation of the hexadecimal address
// (like 0x7f8a1d8ccfd0). If the address cannot be parsed, 0 is returned.
func ParseAddress(address string) Address {
	if !TrackPositions {
		return 0
	}
	addr, _ := strconv.ParseUint(address, 0, 64)

	return Address(addr)
}

// Parse takes the coloured output of the clang AST command and returns a root
// node for the AST.
func Parse(fullline string) Node {
	line := fullline

	// This is a special case. I'm not sure if it's a bug in the clang AST
	// dumper. It should have children.
	if line == "array filler" {
		return parseArrayFiller(line)
	}

	parts := strings.SplitN(line, " ", 2)
	nodeName := parts[0]

	if nodeName == "super" || nodeName == "getter" { // special ObjC case
		parts = strings.SplitN(parts[1], " ", 2)
		nodeName += " " + parts[0]
	}
	// skip node name
	if len(parts) > 1 {
		line = parts[1]
	}

	switch nodeName {
	case "AlignedAttr":
		return parseAlignedAttr(line)
	case "AllocSizeAttr":
		return parseAllocSizeAttr(line)
	case "AlwaysInlineAttr":
		return parseAlwaysInlineAttr(line)
	case "ArraySubscriptExpr":
		return parseArraySubscriptExpr(line)
	case "ArcWeakrefUnavailableAttr":
		return parseArcWeakrefUnavailableAttr(line)
	case "AsmLabelAttr":
		return parseAsmLabelAttr(line)
	case "AttributedType":
		return parseAttributedType(line)
	case "AvailabilityAttr":
		return parseAvailabilityAttr(line)
	case "BinaryOperator":
		return parseBinaryOperator(line)
	case "BlockCommandComment":
		return parseBlockCommandComment(line)
	case "BlockPointerType":
		return parseBlockPointerType(line)
	case "BreakStmt":
		return parseBreakStmt(line)
	case "BuiltinType":
		return parseBuiltinType(line)
	case "CallExpr":
		return parseCallExpr(line)
	case "ConvertVectorExpr":
		return parseConvertVectorExpr(line)
	case "CaseStmt":
		return parseCaseStmt(line)
	case "CFAuditedTransferAttr":
		return parseCFAuditedTransferAttr(line)
	case "CFConsumedAttr":
		return parseCFConsumedAttr(line)
	case "CFReturnsRetainedAttr":
		return parseCFReturnsRetainedAttr(line)
	case "CFReturnsNotRetainedAttr":
		return parseCFReturnsNotRetainedAttr(line)
	case "CharacterLiteral":
		return parseCharacterLiteral(line)
	case "CompoundLiteralExpr":
		return parseCompoundLiteralExpr(line)
	case "CompoundStmt":
		return parseCompoundStmt(line)
	case "ConditionalOperator":
		return parseConditionalOperator(line)
	case "ConstAttr":
		return parseConstAttr(line)
	case "ConstantArrayType":
		return parseConstantArrayType(line)
	case "ContinueStmt":
		return parseContinueStmt(line)
	case "CompoundAssignOperator":
		return parseCompoundAssignOperator(line)
	case "CStyleCastExpr":
		return parseCStyleCastExpr(line)
	case "DeclRefExpr":
		return parseDeclRefExpr(line)
	case "DeclStmt":
		return parseDeclStmt(line)
	case "DefaultStmt":
		return parseDefaultStmt(line)
	case "DeprecatedAttr":
		return parseDeprecatedAttr(line)
	case "DisableTailCallsAttr":
		return parseDisableTailCallsAttr(line)
	case "DoStmt":
		return parseDoStmt(line)
	case "ElaboratedType":
		return parseElaboratedType(line)
	case "EmptyDecl":
		return parseEmptyDecl(line)
	case "Enum":
		return parseEnum(line)
	case "EnumConstantDecl":
		return parseEnumConstantDecl(line)
	case "EnumDecl":
		return parseEnumDecl(line)
	case "EnumExtensibilityAttr":
		return parseEnumExtensibilityAttr(line)
	case "EnumType":
		return parseEnumType(line)
	case "Field":
		return parseField(line)
	case "FieldDecl":
		return parseFieldDecl(line)
	case "FlagEnumAttr":
		return parseFlagEnumAttr(line)
	case "FloatingLiteral":
		return parseFloatingLiteral(line)
	case "FormatAttr":
		return parseFormatAttr(line)
	case "FormatArgAttr":
		return parseFormatArgAttr(line)
	case "FunctionDecl":
		return parseFunctionDecl(line)
	case "FullComment":
		return parseFullComment(line)
	case "FunctionProtoType":
		return parseFunctionProtoType(line)
	case "ForStmt":
		return parseForStmt(line)
	case "HTMLStartTagComment":
		return parseHTMLStartTagComment(line)
	case "HTMLEndTagComment":
		return parseHTMLEndTagComment(line)
	case "GCCAsmStmt":
		return parseGCCAsmStmt(line)
	case "GotoStmt":
		return parseGotoStmt(line)
	case "IBActionAttr":
		return parseIBActionAttr(line)
	case "IBOutletAttr":
		return parseIBOutletAttr(line)
	case "IfStmt":
		return parseIfStmt(line)
	case "ImplicitCastExpr":
		return parseImplicitCastExpr(line)
	case "ImplicitValueInitExpr":
		return parseImplicitValueInitExpr(line)
	case "IncompleteArrayType":
		return parseIncompleteArrayType(line)
	case "IndirectFieldDecl":
		return parseIndirectFieldDecl(line)
	case "InitListExpr":
		return parseInitListExpr(line)
	case "InlineCommandComment":
		return parseInlineCommandComment(line)
	case "IntegerLiteral":
		return parseIntegerLiteral(line)
	case "LabelStmt":
		return parseLabelStmt(line)
	case "MallocAttr":
		return parseMallocAttr(line)
	case "MaxFieldAlignmentAttr":
		return parseMaxFieldAlignmentAttr(line)
	case "MayAliasAttr":
		return parseMayAliasAttr(line)
	case "MemberExpr":
		return parseMemberExpr(line)
	case "MinVectorWidthAttr":
		return parseMinVectorWidthAttr(line)
	case "ModeAttr":
		return parseModeAttr(line)
	case "NoDebugAttr":
		return parseNoDebugAttr(line)
	case "NoEscapeAttr":
		return parseNoEscapeAttr(line)
	case "NoInlineAttr":
		return parseNoInlineAttr(line)
	case "NoThrowAttr":
		return parseNoThrowAttr(line)
	case "NonNullAttr":
		return parseNonNullAttr(line)
	case "NotTailCalledAttr":
		return parseNotTailCalledAttr(line)
	case "NSConsumedAttr":
		return parseNSConsumedAttr(line)
	case "NSConsumesSelfAttr":
		return parseNSConsumesSelfAttr(line)
	case "NSErrorDomainAttr":
		return parseNSErrorDomainAttr(line)
	case "NSReturnsRetainedAttr":
		return parseNSReturnsRetainedAttr(line)
	case "ObjCBoolLiteralExpr":
		return parseObjCBoolLiteralExpr(line)
	case "ObjCBoxableAttr":
		return parseObjCBoxableAttr(line)
	case "ObjCBridgeAttr":
		return parseObjCBridgeAttr(line)
	case "ObjCBridgeRelatedAttr":
		return parseObjCBridgeRelatedAttr(line)
	case "ObjCBridgeMutableAttr":
		return parseObjCBridgeMutableAttr(line)
	case "ObjCCategoryDecl":
		return parseObjCCategoryDecl(line)
	case "ObjCDesignatedInitializerAttr":
		return parseObjCDesignatedInitializerAttr(line)
	case "ObjCExceptionAttr":
		return parseObjCExceptionAttr(line)
	case "ObjCExplicitProtocolImplAttr":
		return parseObjCExplicitProtocolImplAttr(line)
	case "ObjCIndependentClassAttr":
		return parseObjCIndependentClassAttr(line)
	case "ObjCInterface":
		return parseObjCInterface(line,false)
	case "super ObjCInterface":
		return parseObjCInterface(line,true)
	case "ObjCInterfaceDecl":
		return parseObjCInterfaceDecl(line)
	case "ObjCInterfaceType":
		return parseObjCInterfaceType(line)
	case "ObjCIvarDecl":
		return parseObjCIvarDecl(line)
	case "getter ObjCMethod":
		return parseObjCMethod(line)
	case "ObjCMethod":
		return parseObjCMethod(line)
	case "ObjCMessageExpr":
		return parseObjCMessageExpr(line)
	case "ObjCMethodDecl":
		return parseObjCMethodDecl(line)
	case "ObjCObjectType":
		return parseObjCObjectType(line)
	case "ObjCObjectPointerType":
		return parseObjCObjectPointerType(line)
	case "ObjCProtocol":
		return parseObjCProtocol(line)
	case "ObjCReturnsInnerPointerAttr":
		return parseObjCReturnsInnerPointerAttr(line)
	case "ObjCRequiresSuperAttr":
		return parseObjCRequiresSuperAttr(line)
	case "ObjCProtocolDecl":
		return parseObjCProtocolDecl(line)
	case "ObjCRootClassAttr":
		return parseObjCRootClassAttr(line)
	case "ObjCPropertyDecl":
		return parseObjCPropertyDecl(line)
	case "ObjCTypeParamDecl":
		return parseObjCTypeParamDecl(line)
	case "OffsetOfExpr":
		return parseOffsetOfExpr(line)
	case "PackedAttr":
		return parsePackedAttr(line)
	case "ParagraphComment":
		return parseParagraphComment(line)
	case "ParamCommandComment":
		return parseParamCommandComment(line)
	case "ParenExpr":
		return parseParenExpr(line)
	case "ParenType":
		return parseParenType(line)
	case "ParmVarDecl":
		return parseParmVarDecl(line)
	case "PointerType":
		return parsePointerType(line)
	case "DecayedType":
		return parseDecayedType(line)
	case "PredefinedExpr":
		return parsePredefinedExpr(line)
	case "PureAttr":
		return parsePureAttr(line)
	case "QualType":
		return parseQualType(line)
	case "Record":
		return parseRecord(line)
	case "RecordDecl":
		return parseRecordDecl(line)
	case "RecordType":
		return parseRecordType(line)
	case "RestrictAttr":
		return parseRestrictAttr(line)
	case "ReturnStmt":
		return parseReturnStmt(line)
	case "ReturnsTwiceAttr":
		return parseReturnsTwiceAttr(line)
	case "SentinelAttr":
		return parseSentinelAttr(line)
	case "ShuffleVectorExpr":
		return parseShuffleVectorExpr(line)
	case "StmtExpr":
		return parseStmtExpr(line)
	case "StringLiteral":
		return parseStringLiteral(line)
	case "SwiftBridgedTypedefAttr":
		return parseSwiftBridgedTypedefAttr(line)
	case "SwiftErrorAttr":
		return parseSwiftErrorAttr(line)
	case "SwiftNameAttr":
		return parseSwiftNameAttr(line)
	case "SwiftNewtypeAttr":
		return parseSwiftNewtypeAttr(line)
	case "SwiftPrivateAttr":
		return parseSwiftPrivateAttr(line)
	case "SwitchStmt":
		return parseSwitchStmt(line)
	case "TargetAttr":
		return parseTargetAttr(line)
	case "TextComment":
		return parseTextComment(line)
	case "TranslationUnitDecl":
		return parseTranslationUnitDecl(line)
	case "TransparentUnionAttr":
		return parseTransparentUnionAttr(line)
	case "Typedef":
		return parseTypedef(line)
	case "TypedefDecl":
		return parseTypedefDecl(line)
	case "TypedefType":
		return parseTypedefType(line)
	case "UnaryExprOrTypeTraitExpr":
		return parseUnaryExprOrTypeTraitExpr(line)
	case "UnaryOperator":
		return parseUnaryOperator(line)
	case "UnavailableAttr":
		return parseUnavailableAttr(line)
	case "UsedAttr":
		return parseUsedAttr(line)
	case "UnusedAttr":
		return parseUnusedAttr(line)
	case "VAArgExpr":
		return parseVAArgExpr(line)
	case "VarDecl":
		return parseVarDecl(line)
	case "VectorType":
		return parseVectorType(line)
	case "VerbatimBlockComment":
		return parseVerbatimBlockComment(line)
	case "VerbatimBlockLineComment":
		return parseVerbatimBlockLineComment(line)
	case "VerbatimLineComment":
		return parseVerbatimLineComment(line)
	case "VisibilityAttr":
		return parseVisibilityAttr(line)
	case "WarnUnusedResultAttr":
		return parseWarnUnusedResultAttr(line)
	case "WeakAttr":
		return parseWeakAttr(line)
	case "WeakImportAttr":
		return parseWeakImportAttr(line)
	case "WhileStmt":
		return parseWhileStmt(line)
	case "...":
		return parseVariadic(line)
	case "NullStmt":
		return nil
	default:
		return parseUnknown(nodeName,line)
	}
}

func groupsFromRegex(rx, line string) map[string]string {
	// We remove tabs and newlines from the regex. This is purely cosmetic,
	// as the regex input can be quite long and it's nice for the caller to
	// be able to format it in a more readable way.
	fullRegexp := "^(?P<address>[0-9a-fx]+) " +
		strings.Replace(strings.Replace(rx, "\n", "", -1), "\t", "", -1)
	rx = fullRegexp + "[\\s]*$"

	re := util.GetRegex(rx)
	match := re.FindStringSubmatch(line)
	if len(match) == 0 {
		fmt.Printf("AST parser: could not match regexp with string. Regexp:\n" + rx + "\nOriginal line:\n" + line + "\n\n")
		return nil
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	return result
}
