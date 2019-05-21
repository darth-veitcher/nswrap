package main
//go:generate nswrap

import (
	"fmt"
	"runtime"
	"gitlab.wow.st/gmp/nswrap/examples/app/ns"
)

var nst = ns.NSStringWithGoString

func didFinishLaunching(n *ns.NSNotification) {
	fmt.Println("Go: did finish launching")
}

func shouldTerminateAfterLastWindowClosed(s *ns.NSApplication) ns.BOOL {
	return 1
}

func willTerminate(n *ns.NSNotification) {
	fmt.Println("Go: will terminate")
}

func didBecomeActive(n *ns.NSNotification) {
	fmt.Println("Go: did become active")
	vc := win.ContentViewController()
	if vc == nil {
		fmt.Println("vc == nil")
	} else {
		fmt.Println("vc is not nil")
	}
}

var (
	win *ns.NSWindow
)

func app() {
	//Lock OS thread because Cocoa uses thread-local storage
	runtime.LockOSThread()
	a := ns.NSApplicationSharedApplication()
	a.SetActivationPolicy(ns.NSApplicationActivationPolicyRegular)

	//Set up an AppDelegate
	del := ns.AppDelegateAlloc()
	del.ApplicationDidFinishLaunchingCallback(didFinishLaunching)
	del.ApplicationShouldTerminateAfterLastWindowClosedCallback(shouldTerminateAfterLastWindowClosed)
	del.ApplicationWillTerminateCallback(willTerminate)
	del.ApplicationDidBecomeActiveCallback(didBecomeActive)
	a.SetDelegate(del)

	//Set up an NSWindow
	win = ns.NSWindowAlloc().InitWithContentRect(
		ns.NSMakeRect(200,200,600,600),
		ns.NSWindowStyleMaskTitled | ns.NSWindowStyleMaskClosable |
		  ns.NSWindowStyleMaskResizable,
		ns.NSBackingStoreBuffered,
		0,
		nil,
	)
	win.SetTitle(nst("Hi World"))
	win.MakeKeyAndOrderFront(win)
	win.SetAlphaValue(0.85)

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

	appMenu.AddItemWithTitle(nst("About"), nil, nst(""))
	appMenu.AddItemWithTitle(nst("Preferences"), nil, nst(""))
	appMenu.AddItemWithTitle(nst("Quit"),ns.Selector("terminate:"), nst("q"))
	a.SetMainMenu(m1)
	fileMenu.AddItemWithTitle(nst("Open"), nil, nst(""))
	fileMenu.AddItemWithTitle(nst("New"), nil, nst(""))

	a.SetMainMenu(m1)

	//Run the app
	a.Run()
}

func main() {
	//Run our app in an autorelease pool just for fun
	go ns.Autoreleasepool(app)
	select { }
}

