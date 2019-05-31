package ast

import (
	"testing"
)

func TestObjCMethodDecl(t *testing.T) {
	nodes := map[string]Node{
		`0x7f92a58a7570 <line:17:1, col:27> col:1 - isEqual: 'BOOL':'signed char'`:
		&ObjCMethodDecl{
			Addr:         0x7f92a58a7570,
			Pos:          NewPositionFromString("line:17:1, col:27"),
			Position2:    "",
			Name:         "isEqual",
			Type:         "BOOL",
			Type2:        "",
			ClassMethod:  false,
			Parameters:   []string{"isEqual"},
			Implicit:     false,
			ChildNodes:   []Node{},
		},
		`0x7f92a58a7978 </usr/include/objc/NSObject.h:22:1, col:21> col:1 - self 'instancetype':'id'`:
		&ObjCMethodDecl{
			Addr:         0x7f92a58a7978,
			Pos:          NewPositionFromString("/usr/include/objc/NSObject.h:22:1, col:21"),
			Position2:    "",
			Name:         "self",
			Type:         "instancetype",
			Type2:        "",
			ClassMethod:  false,
			Parameters:   []string{},
			Implicit:     false,
			ChildNodes:   []Node{},
		},
		`0x7f92a58a82b0 <line:34:1, col:42> col:1 - respondsToSelector: 'BOOL':'signed char'`:
		&ObjCMethodDecl{
			Addr:         0x7f92a58a82b0,
			Pos:          NewPositionFromString("line:34:1, col:42"),
			Position2:    "",
			Name:         "respondsToSelector",
			Type:         "BOOL",
			Type2:        "",
			ClassMethod:  false,
			Parameters:   []string{"respondsToSelector"},
			Implicit:     false,
			ChildNodes:   []Node{},
		},
		`0x7f92a58a82b0 <line:34:1, col:42> col:1 + instancesRespondToSelector: 'BOOL':'signed char'`:
		&ObjCMethodDecl{
			Addr:         0x7f92a58a82b0,
			Pos:          NewPositionFromString("line:34:1, col:42"),
			Position2:    "",
			Name:         "instancesRespondToSelector",
			Type:         "BOOL",
			Type2:        "",
			ClassMethod:  true,
			Parameters:   []string{"instancesRespondToSelector"},
			Implicit:     false,
			ChildNodes:   []Node{},
		},
		`0x7f92a58a7cd8 <line:26:1, col:83> col:1 - performSelector:withObject:withObject: 'id':'id'`:
		&ObjCMethodDecl{
			Addr:         0x7f92a58a7cd8,
			Pos:          NewPositionFromString("line:26:1, col:83"),
			Position2:    "",
			Name:         "performSelector",
			Type:         "id",
			Type2:        "",
			ClassMethod:  false,
			Parameters:   []string{"performSelector","withObject","withObject"},
			Implicit:     false,
			ChildNodes:   []Node{},
		},
		`0x7f92a4459318 <line:41:71> col:71 implicit - writableTypeIdentifiersForItemProvider 'NSArray<NSString *> * _Nonnull':'NSArray<NSString *> *'`:
		&ObjCMethodDecl{
			Addr:         0x7f92a4459318,
			Pos:          NewPositionFromString("line:41:71"),
			Position2:    "",
			Name:         "writableTypeIdentifiersForItemProvider",
			Type:         `NSArray<NSString *> * _Nonnull`,
			Type2:        "",
			ClassMethod:  false,
			Parameters:   []string{},
			Implicit:     true,
			ChildNodes:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}
