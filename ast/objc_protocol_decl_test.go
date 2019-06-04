package ast

import (
	"testing"
)

func TestObjCProtocolDecl(t *testing.T) {
	nodes := map[string]testNode{
		`0x10a9da630 <line:22:1, col:11> col:11 NSBrowserDelegate`:
		testNode{&ObjCProtocolDecl{
			Addr:         0x10a9da630,
			Pos:          NewPositionFromString("line:22:1, col:11"),
			Position2:    "",
			Name:         "NSBrowserDelegate",
			ChildNodes:   []Node{},
		},
		0x10a9da630,
		NewPositionFromString("line:22:1, col:11"),
		[]Node{},
		},
		`0x10c37fc70 </System/Library/Frameworks/AppKit.framework/Headers/NSPrintPanel.h:58:1, line:70:2> line:58:11 NSPrintPanelAccessorizing`:
		testNode{&ObjCProtocolDecl{
			Addr:         0x10c37fc70,
			Pos:          NewPositionFromString("/System/Library/Frameworks/AppKit.framework/Headers/NSPrintPanel.h:58:1, line:70:2"),
			Position2:    "",
			Name:         "NSPrintPanelAccessorizing",
			ChildNodes:   []Node{},
		},
		0x10c37fc70,
		NewPositionFromString("/System/Library/Frameworks/AppKit.framework/Headers/NSPrintPanel.h:58:1, line:70:2"),
		[]Node{},
		},
	}

	runNodeTests(t, nodes)
}
