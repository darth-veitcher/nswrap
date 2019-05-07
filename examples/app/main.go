package main
//go:generate nswrap

import (
	"runtime"
	"gitlab.wow.st/gmp/nswrap/examples/app/ns"
)

func nsmgr() {
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
	w.SetTitle(ns.NSStringWithGoString("Hi World"))
	w.MakeKeyAndOrderFront(w)
	w.SetAlphaValue(0.85)
	m1 := ns.NSMenuAlloc().InitWithTitle(ns.NSStringWithGoString("Main"))
	appItem := ns.NSMenuItemAlloc()
	fileItem := ns.NSMenuItemAlloc()
	m1.AddItem(appItem)
	m1.AddItem(fileItem)

	appMenu := ns.NSMenuAlloc().InitWithTitle(ns.NSStringWithGoString("App"))
	fileMenu := ns.NSMenuAlloc().InitWithTitle(ns.NSStringWithGoString("File"))
	m1.SetSubmenu(appMenu, appItem)
	m1.SetSubmenu(fileMenu, fileItem)

	s := ns.NSStringWithGoString("")
	appMenu.AddItemWithTitle(ns.NSStringWithGoString("About"), nil, s)
	appMenu.AddItemWithTitle(ns.NSStringWithGoString("Preferences"), nil, s)
	a.SetMainMenu(m1)
	fileMenu.AddItemWithTitle(ns.NSStringWithGoString("Open"), nil, s)
	fileMenu.AddItemWithTitle(ns.NSStringWithGoString("New"), nil, s)

	a.SetMainMenu(m1)
	a.Run()
}

func main() {
	go nsmgr()
	select { }
}

