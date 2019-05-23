package main
//go:generate nswrap

import (
	"fmt"
	"runtime"
	"gitlab.wow.st/gmp/nswrap/examples/app/ns"
)

//Shortcut for literal NSStrings
var nst = ns.NSStringWithGoString

func didFinishLaunching(n ns.NSNotification) {
	fmt.Println("Go: did finish launching")
	fmt.Printf("Notification: %s\n",n.Name().UTF8String())
	//Set up an NSWindow
	win = ns.NSWindowAlloc().InitWithContentRect(
		ns.NSMakeRect(200,200,600,600),
		ns.NSWindowStyleMaskTitled | ns.NSWindowStyleMaskClosable |
		ns.NSWindowStyleMaskResizable,
		ns.NSBackingStoreBuffered,
		0,
		ns.NSScreen{},
	)
	// retain win since we called Alloc and did not add it to a collection
	win.Retain()

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

	//add some buttons and do some layout

	//don't do this:
	//b := ns.NSButtonAlloc().InitWithFrame(ns.NSMakeRect(100,100,100,50))

	b1 := ns.NSButtonWithTitle(nst("PUSH"),ns.Id{},ns.Selector(""))
	b2 := ns.NSButtonWithTitle(nst("QUIT"),ns.Id{},ns.Selector("terminate:"))
	b1.SetTranslatesAutoresizingMaskIntoConstraints(0)
	b2.SetTranslatesAutoresizingMaskIntoConstraints(0)

	cv := win.ContentView()

	cv.AddSubview(b1.NSView,ns.NSWindowAbove,ns.NSView{})
	cv.AddSubview(b2.NSView,ns.NSWindowAbove,ns.NSView{})

	viewmap := ns.NSDictionaryWithObjects(
		ns.NSArrayWithObjects(b1,b2),
		ns.NSArrayWithObjects(nst("b1"),nst("b2")))

	cv.AddConstraints(ns.NSLayoutConstraintsWithVisualFormat(
		nst("V:|-[b1]"),0, ns.NSDictionary{}, viewmap))
	cv.AddConstraints(ns.NSLayoutConstraintsWithVisualFormat(
		nst("H:|-[b1]"),0, ns.NSDictionary{}, viewmap))
	cv.AddConstraints(ns.NSLayoutConstraintsWithVisualFormat(
		nst("H:[b1]-[b2]"),ns.NSLayoutFormatAlignAllBaseline,
		ns.NSDictionary{}, viewmap))
}

func shouldTerminateAfterLastWindowClosed(s ns.NSApplication) ns.BOOL {
	return 1
}

func willTerminate(n ns.NSNotification) {
	fmt.Println("Go: will terminate")
}

func didBecomeActive(n ns.NSNotification) {
	fmt.Println("Go: did become active")
	fmt.Printf("Notification: %s\n",n.Name().UTF8String())
}

var (
	a ns.NSApplication
	win ns.NSWindow
)

func app() {
	//Lock OS thread because Cocoa uses thread-local storage
	runtime.LockOSThread()
	a = ns.NSApplicationSharedApplication()
	a.SetActivationPolicy(ns.NSApplicationActivationPolicyRegular)

	//Set up an AppDelegate
	del := ns.AppDelegateAlloc()
	del.ApplicationDidFinishLaunchingCallback(didFinishLaunching)
	del.ApplicationShouldTerminateAfterLastWindowClosedCallback(shouldTerminateAfterLastWindowClosed)
	del.ApplicationWillTerminateCallback(willTerminate)
	del.ApplicationDidBecomeActiveCallback(didBecomeActive)

	a.SetDelegate(del)

	//Run the app
	a.Run()
}

func main() {
	//Run our app in an autorelease pool just for fun
	go ns.Autoreleasepool(app)
	select { }
}

