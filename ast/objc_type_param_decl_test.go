package ast

import (
	"testing"
)

func TestObjCTypeParamDecl(t *testing.T) {
	nodes := map[string]Node{
		`0x10c2d1a78 <col:27> col:27 AnchorType 'id':'id'`:
		&ObjCTypeParamDecl{
			Addr:         0x10c2d1a78,
			Pos:          NewPositionFromString("col:27"),
			Position2:    "",
			Name:         "AnchorType",
			Type:         "id",
			Type2:        "",
			IsReferenced: false,
			IsCovariant:  false,
			IsBounded:    false,
			ChildNodes:   []Node{},
		},
		`0x7faa181df328 <col:16> col:16 ObjectType covariant 'id':'id'`:
		&ObjCTypeParamDecl{
			Addr:         0x7faa181df328,
			Pos:          NewPositionFromString("col:16"),
			Position2:    "",
			Name:         "ObjectType",
			Type:         "id",
			Type2:        "",
			IsReferenced: false,
			IsCovariant:  true,
			IsBounded:    false,
			ChildNodes:   []Node{},
		},
		`0x7faa18216cf0 <col:26, col:43> col:26 referenced UnitType bounded 'NSUnit *'`:
		&ObjCTypeParamDecl{
			Addr:         0x7faa18216cf0,
			Pos:          NewPositionFromString("col:26, col:43"),
			Position2:    "",
			Name:         "UnitType",
			Type:         "NSUnit *",
			Type2:        "",
			IsReferenced: true,
			IsCovariant:  false,
			IsBounded:    true,
			ChildNodes:   []Node{},
		},
		`0x7faa18ba2ba8 <col:25> col:25 referenced K covariant 'id':'id'`:
		&ObjCTypeParamDecl{
			Addr:         0x7faa18ba2ba8,
			Pos:          NewPositionFromString("col:25"),
			Position2:    "",
			Name:         "K",
			Type:         "id",
			Type2:        "",
			IsReferenced: true,
			IsCovariant:  true,
			IsBounded:    false,
			ChildNodes:   []Node{},
		},
		`0x7faa18ba2c18 <col:28> col:28 V covariant 'id':'id'`:
		&ObjCTypeParamDecl{
			Addr:         0x7faa18ba2c18,
			Pos:          NewPositionFromString("col:28"),
			Position2:    "",
			Name:         "V",
			Type:         "id",
			Type2:        "",
			IsReferenced: false,
			IsCovariant:  true,
			IsBounded:    false,
			ChildNodes:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}
