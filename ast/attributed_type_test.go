package ast

import (
	"testing"
)

func TestAttributedType(t *testing.T) {
	nodes := map[string]testNode{
		`0x10c0d6770 'CVDisplayLinkRef _Nonnull' sugar`: testNode{&AttributedType{
			Addr:       0x10c0d6770,
			Type:       `CVDisplayLinkRef _Nonnull`,
			Sugar:      true,
			ChildNodes: []Node{},
		},
			0x10c0d6770,
			NewPositionFromString(""),
			[]Node{},
		},
		`0x10c0d68b0 'const CVTimeStamp * _Nonnull' sugar`: testNode{&AttributedType{
			Addr:       0x10c0d68b0,
			Type:       `const CVTimeStamp * _Nonnull`,
			Sugar:      true,
			ChildNodes: []Node{},
		},
			0x10c0d68b0,
			NewPositionFromString(""),
			[]Node{},
		},
		`0x10c0d6ab0 'CVOptionFlags * _Nonnull' sugar`: testNode{&AttributedType{
			Addr:       0x10c0d6ab0,
			Type:       `CVOptionFlags * _Nonnull`,
			Sugar:      true,
			ChildNodes: []Node{},
		},
			0x10c0d6ab0,
			NewPositionFromString(""),
			[]Node{},
		},
		`0x10c0fc7d0 'CVPixelBufferRef _Nonnull' sugar`: testNode{&AttributedType{
			Addr:       0x10c0fc7d0,
			Type:       `CVPixelBufferRef _Nonnull`,
			Sugar:      true,
			ChildNodes: []Node{},
		},
			0x10c0fc7d0,
			NewPositionFromString(""),
			[]Node{},
		},
		`0x7faa1906d680 'NSError * _Nullable' sugar`: testNode{&AttributedType{
			Addr:       0x7faa1906d680,
			Type:       `NSError * _Nullable`,
			Sugar:      true,
			ChildNodes: []Node{},
		},
			0x7faa1906d680,
			NewPositionFromString(""),
			[]Node{},
		},
		`0x7faa19085760 'id<NSSecureCoding> _Nullable' sugar`: testNode{&AttributedType{
			Addr:       0x7faa19085760,
			Type:       `id<NSSecureCoding> _Nullable`,
			Sugar:      true,
			ChildNodes: []Node{},
		},
			0x7faa19085760,
			NewPositionFromString(""),
			[]Node{},
		},
		`0x7faa19085840 'NSError * _Null_unspecified' sugar`: testNode{&AttributedType{
			Addr:       0x7faa19085840,
			Type:       `NSError * _Null_unspecified`,
			Sugar:      true,
			ChildNodes: []Node{},
		},
			0x7faa19085840,
			NewPositionFromString(""),
			[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
