# NSWrap

Create Go language bindings for Objective-C.

Using NSWrap, you can work with MacOS interfaces, subclasses,
library functions, protocols and delegates entirely in Go.

# Getting Started

## Installation

NSWrap runs on MacOS and requires `clang` (from the XCode command line
tools) and the MacOS system header files.

```sh
go get git.wow.st/gmp/nswrap
```

The `nswrap` command line tool should now be installed in your `go/bin` path.

Since NSWrap uses `clang` to generate an AST from Objective-C input files, you
will need to install XCode and its associated command line tools. Enter
`clang --version` from your terminal prompt to see if you have it installed.
You will also need the Objective-C header files for the
various frameworks you want to use. Look for them in
`/System/Library/Frameworks/*/Headers`.

## Try Out An Example

NSWrap is designed to be easy to use. To get started with an example, visit
your Go source directory in a terminal and enter:

```sh
cd git.wow.st/gmp/nswrap/examples/app
go generate
go build
./app
```

# Basic Usage

## YAML configuration file

NSWrap takes no command line arguments. All configuration directives are
included in a file named `nswrap.yaml`, which must be found in the directory
from which NSWrap is invoked.

```yaml
# nswrap.yaml example

package: MyWrapper
inputfiles:
  - /System/Library/Frameworks/Foundation.framework/Headers/Foundation.h

classes:
  - NSString
  - NSArray

frameworks: [ Foundation ]
pragma [ clang diagnostic ignored "-Wformat-security" ]
```

Regular expressions are permitted in the names of classes, functions,
protocols and protocol methods, overridden superclass methods, and enums.
Since the `NSObject` class is necessary for memory management, NSWrap will
automatically include it if it is encountered in an input header file.

When invoked, NSWrap creates a subdirectory with the name of the package
as specified in `nswrap.yaml` or, by default, `ns` if a package name is not
specified.
In the output directory, a `main.go` file and, if required, `exports.go`,
will be created or overwritten.

To automatically invoke NSWrap, put a `//go:generate nswrap` comment at the
top of your go source file and use `go generate` to create your Objective-C
bindings.

NSWrap will look for Objective-C header files where directed under
`inputfiles` in your configuration file. CGo will also automatically
compile and link any Objective-C implementation (`.m`) files found in
this output directory, so put them in there if you are going to be
hand-crafting any Objective-C implementations that need to go in the same
package as your automatically generated bindings.

## Class and Instance Methods

NSWrap will create bindings for all classes identified in the `classes`
directive of the configuration file. All of the class and instance methods
are bound to Go and all types identified in the process are wrapped
in Go types (as described below), except for methods that contain unsupported
return types or paramater types such as blocks and function pointers.

```go
s1 := ns.NSStringAlloc()        // allocate an instance of NSString
s2 := ns.NSStringWithSting(s1)  // call a class method of NSString
class := ns.NSStringClass()     // class method returning the class of NSString
fmt.Println(s2.UTF8String())    // call UTF8String, an NSString instance method
```

As seen above, generated class methods will have the same name as their
Objective-C method name, converted to the Go TitleCase convention, prefixed
with the class name, and, if necessary, disambiguated for overloaded
Objective-C methods. Any redundant initial
characters are elided (e.g. the Objective-C
`[NSString stringWithString:aString]` is shortened in Go to
`ns.NSStringWithString(aString)`). Instance methods are converted to
TitleCase and disambiguated for method overloading as described below.

Note that while return types and parameter types needed for the binding will
be defined and wrapped for you in Go types,
you will not get any of their methods 
unless those types also appear in your NSWrap configuration file.
For example, the `[NSDictionary WithObjects: forKeys:]` constructor takes two
`NSArray` parameters, so if you want to use it from Go you will probably want
to have `NSArray` in your configuration file in addition to `NSDictionary`.

## Overloaded Methods

Because Go does not allow overloaded functions, NSWrap automatically
disambiguates overloaded method names as required.
This is done by successively adding parameter names onto the end of the Go
function name until a unique name is created.

For example, `NSString` provides the folowing `compare` methods:

```objective-c
- compare:
- compare:options:
- compare:options:range:
- compare:options:range:locale:
```

These are translated into Go as:

```go
func (o *NSString) Compare(string *NSString) NSComparisonResult { }

func (o *NSString) CompareOptions(string *NSString, mask NSStringCompareOptions) NSComparisonResult { }

func (o *NSString) CompareOptionsRange(string *NSString, mask NSStringCompareOptions,
	rangeOfReceiverToCompare NSRange) NSComparisonResult { }

func (o *NSString) CompareOptionsRangeLocale(string *NSString, mask NSStringCompareOptions,
	rangeOfReceiverToCompare NSRange, locale NSObject) NSComparisonResult { }
```

## NSString Helpers

When NSWrap sees a class or instance method ending in `...WithString` (taking
an Objective-C `NSString` as a parameter), it will automatically create an
additional helper method ending in `WithGoString` that takes a Go string.

```go
str := ns.NSStringWithGoString("** your string goes here **")
fmt.Printf("%s\n",str)
```

NSWrap creates a `Char` Go type that is equivalent to a C `char`. A pointer to
`Char` in Go code can therefore be used with Objective-C functions and methods
that take a `char*` parameter.

When NSWrap binds a method that returns `*Char` (and is in garbage collected mode,
the default), it first calls `strdup`
on the output of the underlying Objective-C method. Therefore, the returned
pointer is manually allocated and will need to be freed later from Go. NSWrap
creates a
`(*Char).Free()` method for use when these pointers are no longer needed.
This copying is necessary because the Objective-C runtime will sometimes
return pointers to internal objects that are impossible to manage from the
Go side. NSWrap aims to cause any internal objects to be deallocated as soon
as possible so they do not cause memory leaks. This means that any returned
C strings need to be copied and memory managed manually from the Go side.

NSWrap provides the helper functions `CharWithGoString` and `CharWithBytes`
that take, respectively, Go strings and Go byte arrays (`[]byte`) and return
`*Char` in Go. As demonstrated above, NSWrap also provides `String()`
methods so that the `*Char` and `*NSString` types implement the `Stringer`
Go interface and therefore can be sent directly to functions like `fmt.Printf`.
The `String()` method on `*NSString` creates a temporary `*Char` internally
but frees it for you before returning. Since methods returning
`*Char` return a pointer that needs to be manually freed, it is important
to use these properly in order to avoid leaks:

```go
nst := ns.NSStringWithGoString("one string")

// NO: the next line leaks a *Char (UTF8String)
//mygostring := nst.UTF8String().String()

// OK: NSWrap creates a temporary *Char and frees it for you:
mygostring := nst.String()

// ALSO OK: manually free your own temporary *Char:
mytmpchar := nst.UTF8String()
mygostring = mytmpchar.String()
mytmpchar.Free()
```
In most cases it will be more convenient to convert directly to Go strings instead
of `*Char`.

## Working With NSObject and its Descendants

Objective-C objects are represented in Go by a type and an interface as
follows:

```go
type Id struct {
	ptr unsafe.Pointer
}
func (o *Id) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }

type NSObject interface {
	Ptr() unsafe.Pointer
}
```
Other object types in Go are structs that directly or indirectly embed `Id`
and therefore contain an `unsafe.Pointer` to an Objective-C object, and
implement `NSObject` by inheriting the `Ptr()` method.

Because of this implementation, you will note that every Objective-C object
is represented by at least two pointers -- an underlying pointer to the 
Objective-C object
in CGo memory (allocated by the Objective-C runtime), as well as a pointer
allocated by the Go runtime to an `Id` type, or to another type that directly
or indirectly embeds `Id`. This "dual pointer" approach is necessary to ensure
that memory management can be made to work correctly (see below for details).

* The NSObject Interface

The `Id` type in Go represents the Objective-C type `id`, which is a pointer
to an Objective-C object. Because `cgo` does not understand this type,
NSWrap will always translate it to a `void*` on the C side.
The `NSObject` interface in Go allows any type that directly or indirectly
embeds `Id` to be used with generic Objective-C functions. For example:

```go
o1 := ns.NSStringWithGoString("my string")
s1 := ns.NSSetWithObjects(o1)
a := ns.NSMutableArrayWithObjects(o1,s1)
```
Since `NSString` and `NSSet` in Go both implement the `NSObject` interface,
they can both be used as parameters to the `NSMutableArray` constructor.

This will help you, too, when working with delegates
(see below). Classes that accept delegates will generally accept any
`NSObject` in their `initWithDelegate()` or `setDelegate()` methods, and
may or may not test at runtime if the provided object actually
implements the required delegate protocol.

* Inheritance

Objective-C allows single inheritance. NSWrap automatically adds
inherited methods to classes that are includled in your binding.

Types created by NSWrap also "embed" their parent class. For example, top
level objects that inherit from `NSObject` in Objective-C
embed the Go type `Id` and therefore implement the `NSObject` Go interface.
Other objects embed their direct superclass. For example:

```go
type NSArray struct { Id }
func (o *NSArray) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSArray() *NSArray {
	return (*NSArray)(unsafe.Pointer(o))
}

type NSMutableArray struct { NSArray }
func (o *NSMutableArray) Ptr() unsafe.Pointer { if o == nil { return nil }; return o.ptr }
func (o *Id) NSMutableArray() *NSMutableArray {
        return (*NSMutableArray)(unsafe.Pointer(o))
}
```

Observe:
```go
b := ns.NSButtonAlloc()        // NSButton > NSControl > NSView > NSResponder > NSObject
b.InitWithFrame(ns.NSMakeRect(100,100,200,200))
b.SetTitle(nst("PUSH"))
vw := win.ContentView()
vw.AddSubview(&b.NSView)	// Pass the button's embedded NSView
```
In Go, `NSButtonAlloc` returns a Go object of type `ns.NSButton`. However,
the `initWithFrame` method is defined in AppKit for `NSView`. NSWrap will find this
method and add it to the Go `NSButton` type when creating your wrapper because
`NSButton` inherits from `NSControl` which inherits from `NSView`.

As of this
writing, on MacOS 10.13.6, NSWrap binds 115 instance methods for `NSObject`,
so things like `Hash()`, `IsEqualTo()`, `ClassName()`, `RespondsToSelector()`
and many many others are available and can be called on any object directly
from Go.

All objects implement the `NSObject` interface, but from time to time you
will encounter a method that takes a parameter of a different type that may
not exactly match the type you have. For example, if you want to pass your
`NSButton` as a parameter to a method that accepts an `NSView` type, you need
to explicitly pass its embedded `NSView` (`&b.NSView` in the example above).
This approach is safer than "converting" the button to an `NSView` (see below)
because it will only work on objects that directly or indirectly embed an
`NSView` Go type.

NSWrap creates a method for `Id` allowing objects to be converted
at run-time to any other class. You will need this for Enumerators and
functions like `NSArray`'s `GetObjects`, for example,
which always return `*Id`. Make
sure you know (or test) what type your objects are before converting them.
You can implement a version of a Go type switch this way:

```go
switch {
case o.IsKindOfClass(ns.NSStringClass()):
        // do something with o.NSString()
case o.IsKindOfClass(ns.NSSetClass()):
        // do something with o.NSSet()
default:
        ...
}
```

Because `Id` can be converted to any type, and every object in the Foundation
classes inherits from `Id`, it is possible to send any message to any
object, if you are feeling lucky. If you are not lucky you will get an
exception from the Objective-C runtime. You are going to have to explicitly
convert your object to the wrong type before the compiler will let you do this.

```go
a := ns.NSArrayWithObjects(o1,o2)      // NSArray embeds Id
fmt.Println(a.NSString().UTF8String()) // DON'T!
//  |         |          \-method of NSString, returns *Char, a "Stringer"
//  |         \-method of Id returning NSString
//  \-calls "String()" on its parameters
```

The above code will compile, but you will get an exception at runtime:

```sh
*** Terminating app due to uncaught exception 'NSInvalidArgumentException', reason:
'-[__NSArrayM UTF8String]: unrecognized selector sent to instance 0x4608940'
```


## Variadic Functions

As seen above with the `NSMutableArrayWithObjects()` constructor example,
NSWrap supports variadic
functions. Because of the limitations of `cgo`, there is a numerical limit
to the number of parameters in a variadic function call, which defaults to
16 but can be set with the `vaargs` configuration directive. NSWrap will
automatically include a `nil` sentinel when calling any Objective-C
methods with variadic parameter lists. The direct types `va_list` and
`va_list_tag` are not currently supported.

## Pointers to Pointers

When NSWrap encounters a pointer to a pointer to an Objective-C object, it
treats it as an array of objects and translates it into a pointer to a
Go slice. If you are passing empty slices into these functions, be sure to
pre-allocate them to a sufficient capacity. Ssee below for an
example. These Go slices can be used for input and output of methods and
functions.

Pointers to pointers are sometimes passed to Objective-C methods or functions
as a way of receiving output from those functions, especially because
Objective-C does not allow for multiple return values. In those cases, after
the CGo call, the method parameter will be treated as an array of
object pointers that may have been modified by the Objective-C function or
method. NSWrap will copy the object pointers back into the input Go slice, up
to its capacity (which will never be changed). The input Go slice is then
truncated to the appropriate length. If there is no output, the length will
be set to 0.

An example in Core Foundation is the `getObjects:andKeys:count` method for
`NSDictionary`:

```go
	nst := ns.NSStringWithGoString
        dict := ns.NSDictionaryWithObjectsForKeys(
                ns.NSArrayWithObjects(nst("obj1"),nst("obj2")),
                ns.NSArrayWithObjects(nst("key1"),nst("key2")),
        )
        va,ka := make([]*ns.Id,0,5), make([]*ns.Id,0,5)  // length 0, capacity 5 slices
        dict.GetObjects(&va,&ka,5)
	// last parameter to GetObjects is the count, must be less than or equal to the input slice capacity
        fmt.Printf("Length of va is now %d\n",len(va)) // va and ka slices are now length = 2
        for i,k := range ka {
                fmt.Printf("-- %s -> %s\n",k.NSString(),va[i].NSString())
        }
```

NSWrap will not check the "count" parameter, so the user will always need
to make sure it is less than or equal to the capacity of the input
Go slices.

Using pointers to pointers is necessary in many Core Foundation situations
where you need to get an error message out of a function or method, or in other
cases where an Objective-C method wants to provide multiple return values.
Here is an example using `[NSString stringWithContentsOfURL...]`:

```go
        err := make([]*ns.NSError,1)
        n1 = ns.NSStringWithContentsOfURLEncoding(ns.NSURLWithGoString("htttypo://example.com"), 0, &err)
	if len(err) > 0 {
        	fmt.Printf("err: %s\n",err[0].LocalizedDescription())
//err: The file couldn’t be opened because URL type htttypo isn’t supported.
	}
```

## Selectors

You can specify selectors using a Go string. The `Selector()` function
returns a Go type `SEL` which corresponds to a pointer to
`struct objc_selector` in C.
Among other things, this lets you set actions on `NSControls` and `NSMenuItems`:

```go
appMenu.AddItemWithTitle(
		ns.NSStringWithGoString("Quit"),
		ns.Selector("terminate:"),
		ns.NSStringWithGoString("q"))
```

## Enumerators

NSWrap provides a `ForIn` method for the `NSEnumerator` type. Call it with a
`func(*ns.Id) bool` parameter that returns `true` to continue and `false` to
stop the enumeration.

```go
a := ns.NSArrayWithObjects(o1,o2,o3)
i := 0
a.ObjectEnumerator().ForIn(func (o *ns.Id) bool {
	switch {
	case o.IsKindOfClass(ns.NSStringClass()):
		fmt.Printf("%d: %s\n", i, o.NSString())
		i++
		return true  // continue enumeration
	default:
		fmt.Println("Unknown class")
		return false  // terminate enumeration
	}
})
```

As seen above, you can do the usual Objective-C thing for runtime type
identification.

## Enum Definitions

NSWrap translates C `enum` values into Go constants. The enums you want are
specified in `nswrap.yaml` by regular expression, which, in the case of named
enums, must match the name of the `enum` itself, or in the case of anonymous
enums, must match the name of the constant(s) you are looking for as declared
within the `enum`.
The generated constants receive Go types associated with their underlying C
types, which are automatically declared by NSWrap as needed.

The following configuration:

```yaml
# nswrap.yaml
inputfiles: [/System/Library/Frameworks/AppKit.framework/Headers/AppKit.h]
enums:
    - _CLOCK.*             # match constants in an anonymous enum
    - NSWindowOrdering.*   # match a named enum
```

results in:

```go
//ns/main.go
...
type NSWindowOrderingMode C.enum_NSWindowOrderingMode
const NSWindowAbove NSWindowOrderingMode = C.NSWindowAbove
const NSWindowBelow NSWindowOrderingMode = C.NSWindowBelow
const NSWindowOut NSWindowOrderingMode = C.NSWindowOut

const _CLOCK_REALTIME  = C._CLOCK_REALTIME
const _CLOCK_MONOTONIC  = C._CLOCK_MONOTONIC
const _CLOCK_MONOTONIC_RAW  = C._CLOCK_MONOTONIC_RAW
...
```

## Delegates

The `delegates` directive in `nswrap.yaml` creates a new Objective-C
class and associated Go wrapper functions. For example, the following
configuration file creates a class called `CBDelegate` that implements
the `CBCentralManagerDelegate` and `CBPeripheralDelegate`
protocols from Core Bluetooth, along with the Go code you need to allocate
and use instances of the new class.

```yaml
# nswrap.yaml
inputfiles:
    - /System/Library/Frameworks/CoreBluetooth.framework/Headers/CoreBluetooth.h

classes:
    - CBCentralManager

delegates:
    CBDelegate:                                  # a name for your delegate class
        CBCentralManagerDelegate:                # a protocol to implement
            - centralManagerDidUpdateState       # messages you want to respond to
            - centralManagerDidDiscoverPeripheral
            - centralManagerDidConnectPeripheral
        CBPeripheralDelegate:                    # another protocol to implement
            - peripheralDidDiscoverServices
            - peripheralDidDiscoverCharacteristicsForService
            - peripheralDidUpdateValueForCharacteristic
...
```

The generated delegate inherits from `NSObject` and, in its interface
declaration, is advertised as implementing the protocols specified in
`nswrap.yaml`.

When a delegate is activated and one of the callback methods named in the
configuration file is called, the delegate will call back into a Go 
function exported by NSWrap. If a user-defined callback function has been
registered,
it will be called with all of its parameters converted to their Go type
equivalents. User-defined callbacks are registered by calling a function
with the method name in TitleCase + `Callback`, so in the example above,
if your delegate was named `del`, you would call
`del.CentralManagerDidUpdateStateCallback(...)` with the name of
your callback function to register to receive notifications when your central
manager updates its state.

The example in `examples/bluetooth` implements a working Bluetooth Low-Energy
heart rate monitor entirely in Go.

The following Go code instantiates a `CBDelegate` object,
registers a callback for `centralManagerDidUpdateState`, allocates
a `CBCentralManager` object, and installs our delegate:

```go
func cb(c ns.CBCentralManager) {
	...
}

var (
	del *ns.CBDelegate // use global variables so these don't get garbage collected
	cm *ns.CBCentralManager
)

func main() {
	...
	del = ns.CBDelegateAlloc()
	del.CentralManagerDidUpdateStateCallback(cb)
	cm = ns.CBCentralManagerAlloc().InitWithDelegateQueue(del,queue)
```

When you provide user-defined callback functions, you will need to specify
them with exactly the right type,
matching NSWrap's generated Go wrapper types for the callback function and
the Go types for all of its parameters. If `go build` fails, the error
messages will point you in the right direction.

```
$ go build
./main.go:127:43: cannot use didFinishLaunching (type func(*ns.NSNotification, bool)) as type
func(*ns.NSNotification) in argument to del.ApplicationDidFinishLaunchingCallback
```
In the above example, the build failed because an extra `bool` parameter was
included in the callback function. The compiler is telling you that the right
type for the callback is `func(*ns.NSNotification)` with no return value.

## Working with AppKit

You can wrap the AppKit framework classes and create an `NSApplication`
Delegate. This allows you to build a Cocoa application entirely in Go.

Because AppKit uses thread local storage, you will need to make sure all
calls into it are done from the main OS thread. This can be a challenge in
Go and you will want to make use of `runtime.LockOSThread()`.

This is actually a full working Cocoa application:

```yaml
# nswrap.yaml
inputfiles:
    - /System/Library/Frameworks/AppKit.framework/Headers/AppKit.h

classes:
    - NSApplication
    - NSWindow
    - NSString
    - NSMenu

enums:
    - NSApplication.*
    - NSBackingStore.*
    - NSWindowStyleMask.*

functions:
    - NSMakeRect

delegates:
    AppDelegate:
      NSApplicationDelegate:
        - applicationDidFinishLaunching
        - applicationShouldTerminateAfterLastWindowClosed
frameworks: [ Foundation, AppKit, CoreGraphics ]
```

```go
//go:generate nswrap
package main

import (
	"fmt"
	"runtime"
	"git.wow.st/gmp/nswrap/examples/app/ns" // point to your own NSWrap output directory
)

func didFinishLaunching(n *ns.NSNotification) {
	fmt.Println("Go: did finish launching!")
}

func shouldTerminate(s *ns.NSApplication) ns.BOOL {
	return 1
}

var (
	a *ns.NSApplication // global vars so these are not garbage collected
	del *ns.AppDelegate
	win *ns.NSWindow
)

func main() {
	runtime.LockOSThread()
	a = ns.NSApplicationSharedApplication()
	a.SetActivationPolicy(ns.NSApplicationActivationPolicyRegular)
	del = ns.AppDelegateAlloc()
	del.ApplicationDidFinishLaunchingCallback(didFinishLaunching)
	del.ApplicationShouldTerminateAfterLastWindowClosedCallback(shouldTerminate)
	a.SetDelegate(del)

	win = ns.NSWindowAlloc().InitWithContentRectStyleMask(
		ns.NSMakeRect(200,200,600,600),
		ns.NSWindowStyleMaskTitled | ns.NSWindowStyleMaskClosable,
		ns.NSBackingStoreBuffered,
		0,
	)
	win.SetTitle(ns.NSStringWithGoString("Hi World"))
	win.MakeKeyAndOrderFront(win)
	a.Run()
}
```

Pretty simple right? Not really, NSWrap just generated over 39,000 lines of
code. See `examples/app` for a slightly more complex example with working
menus, visual format-based auto layout, and a custom button class.

## Subclasses

NSWrap includes functionality to generate subclasses as specified in
`nswrap.yaml`.

You can override existing methods or create new methods with any type
signature you specify using Objective-C method signature syntax.

```yaml
# nswrap.yaml
...
subclasses:
  myClass:                      # the name of the new class
    yourClass:                  # the superclass to inherit from
      - init.*                  # what methods to override
      - -(void)hi_there:(int)x  # Objective-C prototype of your new method(s)
#       \--the initial hyphen indicates that this is an instance method
```

In the example above, your new class will be named `myClass` in Objective-C
and `MyClass` in Go. It will override any `init` methods found in `yourClass`
(which must be defined in one of the header files included in the
`inputfiles` directive of `nswrap.yaml`). In addition, because the second
entry under `yourClass` starts with a `-`, it will be treated as a new
instance method definition for `myClass`. The remainder of the line will
be parsed as an Objective-C method prototype in order to determine the method
name, its return type, and the names and types of its parameters if any.

Since multiple inheritance is not permitted in Objective-C, it is not possible
to specify more than one superclass in a `subclasses` entry.

Go callbacks for overridden methods are passed a special struct 
as their first parameter. This struct is filled with superclass methods, which
allows you to do things like this:

```go
func methodCallback(super ns.MyClassSupermethods, param NSString) {
	...
	super.Method(param)
}
```

You can use subclasses to define new AppKit controls with configurable
callbacks. For example, let's make an `NSButton` that calls back into Go when
you press it:

```yaml
# nswrap.yaml
...
subclasses:
    GButton:
        NSButton:
            - -(void)pressed
...
```

```go
func pressed() {
	fmt.Println("Button pressed!")
}
...
func didFinishLaunching(n ns.NSNotification) {
	...
	button := ns.GButtonAlloc()
	button.Init()
	button.PressedCallback(pressed)		# register user-defined callback
	button.SetAction(ns.Selector("pressed"))
	button.SetTarget(button)
	button.SetTitle(ns.NSStringWithGoString("PUSH"))
	...
}
```

Later on you can add your new button to a view and tell Cocoa where to lay
it out. It's all a little verbose, but that's because for some reason you
decided to write Objective-C code in Go.

## Memory management

As mentioned above, NSWrap is designed for there to be at least one Go pointer
associated with each underlying Objective-C object pointer.
Since Objective-C memory
is always allocated by the Objective-C runtime, it is not possible for the Go
runtime to have visibility into these memory regions or to directly manage memory
used by the CGo code. However, Go will keep track of the associated Go pointer
that was created the first time the corresponding Objective-C object was passed
over to the Go side and an `Id` or other NSWrap struct type was allocated.
Because of this,
it is possible to hook into the Go garbage collection system in an attempt to
manage Objective-C memory strictly from the Go side. When there are no remaining Go
pointers to an NSWrap `Id` struct, it will be deallocated by the Go garbage collector
and a finalizer will be called that `release`es the corresponding Objective-C
object.

The memory management rules work as follows:

* Objects in Go are represented by pointers to types that implement the `NSObject`
interface
* NSObject has one method, `Ptr()`, which returns an `unsafe.Pointer` to an 
Objective-C object.
* All methods that return objects to Go call `retain` except for `new`, `init`,
`alloc`, `copy` and `mutableCopy`, which already return retained objects from the
Objective-C runtime.
* Go wrappers for Objective-C methods call `runtime.SetFinalizer()`, which calls
`release` when the associated Go struct is garbage collected.
* All Objective-C methods are run inside an `@autoreleasepool {}` block to prevent
internal memory leaks within the Objective-C libraries and frameworks.
* Objects sent to you in callback functions are not memory managed by Go
and must be manually
managed using `Retain()` and `Release()` methods if you need to take ownership of them.
A rule of thumb is that if you assign such an object to a persistent Go variable for
use outside of the callback, call `Retain()`.

Because of the linkage with the Go garbage collector described above, there should be
no need for any memory management code to be written from the Go side, except in the
case mentioned above where your Go delegate receives objects that need to be kept
around outside of the callback.

Since everything in Objective C inherits methods from `NSObject`, you can call
`Retain()`, `Release()` and `Autorelease()` on any object. You can technically bind
the `NSAutoreleasePool` class and create and drain instances of it from the Go side,
but this is not recommended in the default, garbage collected mode and can run into
problems because the Go runtime is inherently multithreaded. See `examples/memory`
for an example of manual memory management, which should be possible to do reliably
but I'm not sure why you would go through the trouble.

NSWrap is doing a number of things behind the scenes to make garbage collection work.
As mentioned,
all Objective-C methods are called within an `@autorelease {}` block. This is
necessary because some foundation classes (notably `NSString`) create internal
objects that are `autoreleased` but never returned to the caller. These objects can
never be deallocated unless the method in question was called within an autorelease
pool.

NSWrap assumes you
are going to take ownership of every Objective-C object returned by a method, either
directly as a return value or through a pointer to a pointer given as a parameter.
Therefore, NSWrap calls `retain` on all of these objects before going back to the
Go side, unless the object is either `nil` or equivalent to the input object.
NSWrap also will not call `retain` on the return values of `init`, `new`, `copy`,
`mutableCopy` or `alloc` methods. If you do not want ownership of the object,
simply assign it to a local varable and the garbage collector will take care of
releasing it.

In
order for this to work on a pointer to a pointer parameter, NSWrap treats the
input
parameter as an array with a length specified by either a `range` parameter (of
type `NSRange`) or a `count` parameter of an integer type. If there is neither a
`range` or `count` parameter, NSWrap assumes the array is length 1.

As an example, in Objective-C, if you were to take an object out of an `NSArray`
and the array was later deallocated, there is no guarantee that the object you
obtained is still around unless you called `retain` on it. This is not necessary
with NSWrap, which automatically retains objects returned by methods like
`objectAtIndex:` and `getObjects:range` and manages them with the Go garbage
collector.

The methods described above work for methods that return Objective-C
objects, which can be `retain`ed, but not with methods that return other types of
pointers such as C strings. NSWrap has a special case for C strings (`*Char`
in Go), calling
`strdup` on the return value within the `@autoreleasepool` block. This
ensures that the string is preserved even if it points to a termporary
autoreleased
object. Since this behavior results in newly allocated memory, these pointers will
need to be freed from Go later on. Since these are pointers to C memory,
it is not possible to set a finalizer on these pointers for garbage collection
by Go.

Note that the Go garbage collector is lazy and will not activate unless your
application is running low on heap space. That means in practice that Objective-C
objects are going to stick around a lot longer than they might in a pure
Objective-C application. If this is an issue, simply run the Go GC manually with
`runtime.GC()`.


# Limitations

## Blocks and Function Pointers

NSWrap does not support methods or functions that take C functions or blocks
as parameters or return values.

# Why?

Um, I was trying to make a nice modern Go binding for CoreBluetooth on MacOS
and got carried away.

# Acknowledgements

This work was inspired by Maxim Kupriianov's excellent
[c-for-go](https://github.com/xlab/c-for-go). Much of the
infrastructure was lifted from Elliot Chance's equally excellent
[c2go](https://github.com/elliotchance/c2go). Kiyoshi Murata's
post on [coderwall.com](https://coderwall.com/p/l9jr5a/accessing-cocoa-objective-c-from-go-with-cgo)
was an essential piece of inspiration.

The combinatorial Objective-C type parsers are mine as are the 
Objective-C and Go code generators, so this is where you will find
all of the bugs.
