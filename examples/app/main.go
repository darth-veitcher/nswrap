package main
//go:generate nswrap

import (
	"fmt"
	"runtime"
	"gitlab.wow.st/gmp/nswrap/examples/app/ns"
)

func didFinishLaunching(n *ns.NSNotification) {
	fmt.Println("Go: did finish launching")
}

func shouldTerminate(s *ns.NSApplication) ns.NSApplicationTerminateReply {
	fmt.Println("Go: should terminate")
	return ns.NSTerminateNow
}

func shouldTerminateAfterLastWindowClosed(s *ns.NSApplication) ns.BOOL {
	return 1
}

func willTerminate(n *ns.NSNotification) {
	fmt.Println("Go: will terminate")
}

func didBecomeActive(n *ns.NSNotification) {
	fmt.Println("Go: did become active")
}

func app() {
	//Lock OS thread because Cocoa uses thread-local storage
	runtime.LockOSThread()
	a := ns.NSApplicationSharedApplication()
	a.SetActivationPolicy(ns.NSApplicationActivationPolicyRegular)

	//Set up an AppDelegate
	del := ns.AppDelegateAlloc()
	del.ApplicationDidFinishLaunchingCallback(didFinishLaunching)
	del.ApplicationShouldTerminateCallback(shouldTerminate)
	del.ApplicationShouldTerminateAfterLastWindowClosedCallback(shouldTerminateAfterLastWindowClosed)
	del.ApplicationWillTerminateCallback(willTerminate)
	del.ApplicationDidBecomeActiveCallback(didBecomeActive)
	a.SetDelegate(del)

	//Set up an NSWindow
	w := ns.NSWindowAlloc().InitWithContentRect(
		ns.NSMakeRect(200,200,600,600),
		ns.NSWindowStyleMaskTitled | ns.NSWindowStyleMaskClosable,
		ns.NSBackingStoreBuffered,
		0,
		nil,
	)
	nst := ns.NSStringWithGoString
	w.SetTitle(nst("Hi World"))
	w.MakeKeyAndOrderFront(w)
	w.SetAlphaValue(0.85)

	//Build a basic menu
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

	//Add a random button that does nothing
	b1 := ns.NSButtonWithTitle(nst("push"),s,nil)
	b1.Id.NSView().SetFrame(ns.NSMakeRect(0,550,100,50))
	w.ContentView().AddSubview(&b1.NSView,ns.NSWindowAbove,nil)

	//Run the app
	a.Run()
}

func main() {
	//Run our app in an autorelease pool just for fun
	go ns.Autoreleasepool(app)
	select { }
}

