package main

//go:generate nswrap

import (
	"fmt"
	"git.wow.st/gmp/nswrap/examples/app/ns"
	"runtime"
	"time"
)

//Shortcut for literal NSStrings
var nst = ns.NSStringWithGoString

func pb1() {
	fmt.Println("Pushed button 1")
}

func pb2() {
	fmt.Println("Pushed button 2")
	a.Terminate(a)
}

func db() {
	fmt.Println("button deallocated")
}

func didFinishLaunching(n *ns.NSNotification) {
	fmt.Println("Go: did finish launching")
	fmt.Printf("Notification: %s\n", n.Name().UTF8String())
	//Set up an NSWindow
	win = ns.NSWindowAlloc().InitWithContentRectStyleMask(
		ns.NSMakeRect(200, 200, 600, 600),
		ns.NSWindowStyleMaskTitled|ns.NSWindowStyleMaskClosable|
			ns.NSWindowStyleMaskResizable,
		ns.NSBackingStoreBuffered,
		0,
	)
	// We do not need to retain this because we are in garbage collection mode
	// and have assigned it to a global variable.
	//win.Retain()

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
	appMenu.AddItemWithTitle(nst("Quit"), ns.Selector("terminate:"), nst("q"))
	a.SetMainMenu(m1)
	fileMenu.AddItemWithTitle(nst("Open"), nil, nst(""))
	fileMenu.AddItemWithTitle(nst("New"), nil, nst(""))

	a.SetMainMenu(m1)

	//add some custom buttons

	b1 := ns.GButtonAlloc()
	b2 := ns.GButtonAlloc()

	b1.Init()
	b1.PressedCallback(pb1)
	b1.DeallocCallback(db)
	b1.SetAction(ns.Selector("pressed"))
	b1.SetTarget(b1)
	b1.SetTitle(nst("PUSH"))

	b2.Init()
	b2.PressedCallback(pb2)
	b2.DeallocCallback(db)
	b2.SetTarget(b2)
	b2.SetAction(ns.Selector("pressed"))
	b2.SetTitle(nst("QUIT"))

	//add some layout constraints

	b1.SetTranslatesAutoresizingMaskIntoConstraints(0)
	b2.SetTranslatesAutoresizingMaskIntoConstraints(0)

	cv := win.ContentView()

	cv.AddSubview(&b1.NSView)
	cv.AddSubview(&b2.NSView)

	viewmap := ns.NSDictionaryWithObjectsForKeys(
		ns.NSArrayWithObjects(b1, b2),
		ns.NSArrayWithObjects(nst("b1"), nst("b2")))

	cv.AddConstraints(ns.NSLayoutConstraintsWithVisualFormat(
		nst("V:|-[b1]"), 0, nil, viewmap))
	cv.AddConstraints(ns.NSLayoutConstraintsWithVisualFormat(
		nst("H:|-[b1]"), 0, nil, viewmap))
	cv.AddConstraints(ns.NSLayoutConstraintsWithVisualFormat(
		nst("H:[b1]-[b2]"), ns.NSLayoutFormatAlignAllBaseline, nil, viewmap))

	a.ActivateIgnoringOtherApps(1)
}

func shouldTerminateAfterLastWindowClosed(s *ns.NSApplication) ns.BOOL {
	fmt.Println("Go: should terminate after last window closed")
	return 1
}

func willTerminate(n *ns.NSNotification) {
	fmt.Println("Go: will terminate")
}

func didBecomeActive(n *ns.NSNotification) {
	fmt.Println("Go: did become active")
	fmt.Printf("Notification: %s\n", n.Name().UTF8String())
}

var (
	a   *ns.NSApplication
	del *ns.AppDelegate
	win *ns.NSWindow
)

func app() {
	// Lock OS thread because Cocoa uses thread-local storage
	runtime.LockOSThread()
	a = ns.NSApplicationSharedApplication()
	a.SetActivationPolicy(ns.NSApplicationActivationPolicyRegular)

	// Set up an AppDelegate
	// assign it to a global variable so it doesn't get garbage collected
	del = ns.AppDelegateAlloc()
	del.ApplicationDidFinishLaunchingCallback(didFinishLaunching)
	del.ApplicationShouldTerminateAfterLastWindowClosedCallback(shouldTerminateAfterLastWindowClosed)
	del.ApplicationWillTerminateCallback(willTerminate)
	del.ApplicationDidBecomeActiveCallback(didBecomeActive)

	a.SetDelegate(del)

	// Run the app
	a.Run()
}

func main() {
	// Run GC every second to ensure pointers are not being prematurely released.
	go func() {
		for {
			runtime.GC()
			time.Sleep(time.Second)
		}
	}()

	// Run our app in an autorelease pool just for fun
	go app()
	select {}
}
