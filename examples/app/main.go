package main
//go:generate nswrap

import (
	"runtime"
	"gitlab.wow.st/gmp/nswrap/examples/app/ns"
)

func app() {
	//Lock OS thread because Cocoa uses thread-local storage
	runtime.LockOSThread()
	a := ns.NSApplicationSharedApplication()
	a.SetActivationPolicy(ns.NSApplicationActivationPolicyRegular)
	w := ns.NSWindowAlloc().InitWithContentRect(
		ns.NSMakeRect(200,200,600,600),
		ns.NSWindowStyleMaskTitled,
		ns.NSBackingStoreBuffered,
		0,
		nil,
	)
	nst := ns.NSStringWithGoString
	w.SetTitle(nst("Hi World"))
	w.MakeKeyAndOrderFront(w)
	w.SetAlphaValue(0.85)
	m1 := ns.NSMenuAlloc().InitWithTitle(nst("Main"))
	appItem := ns.NSMenuItemAlloc()
	fileItem := ns.NSMenuItemAlloc()
	m1.AddItem(appItem)
	m1.AddItem(fileItem)

	appMenu := ns.NSMenuAlloc().InitWithTitle(nst("App"))
	fileMenu := ns.NSMenuAlloc().InitWithTitle(nst("File"))
	m1.SetSubmenu(appMenu, appItem)
	m1.SetSubmenu(fileMenu, fileItem)

	s := ns.NSStringWithGoString("")
	appMenu.AddItemWithTitle(nst("About"), nil, s)
	appMenu.AddItemWithTitle(nst("Preferences"), nil, s)
	appMenu.AddItemWithTitle(nst("Quit"),ns.Selector("terminate:"), nst("q"))
	a.SetMainMenu(m1)
	fileMenu.AddItemWithTitle(nst("Open"), nil, s)
	fileMenu.AddItemWithTitle(nst("New"), nil, s)

	a.SetMainMenu(m1)

	b1 := ns.NSButtonWithTitle(nst("push"),s,nil)
	b1.SetFrame(ns.NSMakeRect(0,550,100,50))
	w.ContentView().AddSubview(&b1.NSView,ns.NSWindowAbove,nil)
	a.Run()
}

func main() {
	go ns.Autorelease(app)
	select { }
}

