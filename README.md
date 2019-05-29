# NSWrap

Create Go language bindings for Objective-C.

Using NSWrap, you can easily work with many MacOS interfaces, subclasses,
library functions, protocols and delegates entirely in Go.

# Getting Started

## Installation

NSWrap runs on MacOS and requires `clang` (from the XCode command line
tools) and the MacOS system header files.

```sh
go get git.wow.st/gmp/nswrap/...
```

From your `go` source directory, type:

```sh
cd git.wow.st/gmp/nswrap
go install
```

Since NSWrap uses `clang` to generate an AST from Objective-C input files, you
will need to install XCode and its associated command line tools. Enter
`clang --version` from your terminal prompt to see if you have it installed
already. You will also need to have the Objective-C header files for the
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
in Go types (as described below), except for methods that contain prohibited
return types or paramater types (such as blocks and function pointers).

```go
s1 := ns.NSStringAlloc()        // allocate and autorelease an instance of NSString
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
`ns.NSStringWithString(aString)`). Instance methods are carried over
as-is but in TitleCase, and disambiguated for method overloading as described
below.

Note that while return types and parameter types needed for the binding will
be defined and wrapped for you in Go types,
you will not get any of their methods 
unless those types also appear in your NSWrap configuration file.
For example, the `NSDictionaryWithObjects(...)` constructor takes two `NSArray`
parameters, so if you want to use it you will probably want
to have `NSArray` in your configuration file in addition to `NSDictionary`.

## Overloaded Methods

Because Go does not allow overloaded function definitions, NSWrap automatically
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
func (o NSString) Compare(string NSString) NSComparisonResult { }

func (o NSString) CompareOptions(string NSString, mask NSStringCompareOptions) NSComparisonResult { }

func (o NSString) CompareOptionsRange(string NSString, mask NSStringCompareOptions,
	rangeOfReceiverToCompare NSRange) NSComparisonResult { }

func (o NSString) CompareOptionsRangeLocale(string NSString, mask NSStringCompareOptions,
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
NSWrap provides the helper functions `CharWithGoString` and `CharWithBytes`
that take, respectively, Go strings and Go byte arrays (`[]byte`) and return
`*Char` in Go. As demonstrated above, NSWrap also provides a `String()`
methods so that the `*Char` and `NSString` types implement the `Stringer`
Go interface.

## Working With NSObject and its Descendants

Objective-C Objects are represented in Go by a type and an interface as
follows:

```go
type Id struct {
	ptr unsafe.Pointer
}
func (o Id) Ptr() unsafe.Pointer { return o.ptr }

type NSObject interface {
	Ptr() unsafe.Pointer
}
```
Other object types in Go are structs that directly or indirectly embed `Id`
and therefore implement `NSObject`.

* The NSObject Interface

The `Id` type in Go represents the Objective-C type `id`, which is a pointer
to an Objective-C object. Because `cgo` does not understand this type,
NSWrap will always translate it to a `void*` on the C side.
The `NSObject` interface in Go allows any `NS` type to be used with
generic Objective-C functions. For example:

```go
o1 := ns.NSStringWithGoString("my string")
s1 := ns.NSSetWithOBjects(o1)
a := ns.NSMutableArrayWithObjects(o1,s1)
```
Since `NSString` and `NSSet` in Go both implement the `NSObject` interface,
they can both be used as parameters to the `NSMutableArray` constructor.

This will help you, too, with delegates
(see below). Classes that accept delegates will generally accept any
`NSObject` in ther `initWithDelegate()` or `setDelegate()` methods, and
may or may not test at runtime if the provided object actually
implements the required delegate protocol.

* Inheritance

Objective-C permits single inheritance. In Go, this is modeled using
embedding. Top level objects that inherit from `NSObject` in Objective-C
embed the Go type `Id` and therefore implement the `NSObject` Go interface.
Other objects embed their superclass. For example:

```go
type NSArray struct { Id }
func (o NSArray) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSArray() NSArray {
	ret := NSArray{}
	ret.ptr = o.ptr
	return ret
}

type NSMutableArray struct { NSArray }
func (o NSMutableArray) Ptr() unsafe.Pointer { return o.ptr }
func (o Id) NSMutableArray() NSMutableArray {...}
```

Observe:
```go
b := ns.NSButtonAlloc()        // NSButton > NSControl > NSView > NSResponder > NSObject
b.InitWithFrame(ns.NSMakeRect(100,100,200,200)) // Method of NSView
b.SetTitle(nst("PUSH"))                         // Method of NSButton
vw := win.ContentView()
vw.AddSubview(b.NSView)				// Pass the button's embedded NSView
```
In Go, `NSButtonAlloc` returns a Go object of type `ns.NSButton`. However,
there is no `InitWithFrame` method for receivers of this type. This is
not necessary because `NSButton` embeds `NSControl` which in turn embeds
`NSView`. The `InitWithFrame` method only needs to be implemented for `NSView`
receivers. Go will automatically find the indirectly embedded `NSView` and
call the right method.

Go's type inference appears to be slightly broken (as of 1.12.1) because
the following does not work. Look out for this if you are getting type
errors:

```go
//DO NOT DO THIS
b := ns.NSButtonAlloc().InitWithFrame(ns.MakeRect(100,100,200,200))
//For some reason Go thinks b has type ns.NSView, because InitWithFrame is defined for ns.NSView, even though
//NSButtonAlloc() returns an ns.NSButton.
```

Go has no trouble finding embedded methods for your `NSButton` and will
happily search up the chain through `NSControl`, `NSView`, `NSResponder` and
`NSObject` and all of their associated protocols and categories. As of this
writing, on MacOS 10.13.6, NSWrap binds 90 instance methods for `NSObject`,
so things like `Hash()`, `IsEqualTo()`, `ClassName()` and many many
others are available and can be called on any object directly from Go.

Go does not perform the same type
magic when you use variables as function or method parameters.
If you want to pass your `NSButton` as a parameter to a method that accepts
an `NSView` type, you need to explicitly pass the embedded `NSView`
(`b.NSView` in the example above).

NSWrap creates a method for `Id` allowing objects to be converted
at run-time to any other class. You will need this for Enumerators, which
always return `Id`. See below under Enumerators for an example, but make
sure you know (or test) what type your objects are before converting them,
or else you will get an exception from the Objective-C runtime.

Because `Id` can be converted to any type, and every object in the Foundation
classes inherits from `Id`, it is possible to send any message to any
object, if you are feeling lucky. You are going to have to explicitly
convert your object to the wrong type before the compiler will let you do this.

```go
a := ns.NSArrayWithObjects(o1,o2)      // NSArray embeds Id
fmt.Println(a.NSString().UTF8String()) // DON'T!
//  |         |          \-method of NSString, returns *Char, a "Stringer" type
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
16 but can be set with the `vaargs` configuration directive.

## Pointers to Pointers

When NSWrap encounters a pointer to a pointer to an Objective-C object, it
treats it as an array of objects and translates it into a pointer to a
Go slice. If you are passing empty slices into these functions, be sure to
pre-allocate them to a sufficient size and capacity (see below for an
example). These Go slices can be used for input and output of methods and
functions.

Pointers to pointers are sometimes passed to Objective-C methods or functions
as a way of receiving output from those functions. In those cases, after the
CGo call, the method parameter is treated as a nil-terminated array of object
pointers. The object pointers are copied into the input Go slice, which is
then  truncated to the appropriate length.

An example in Core Foundation is the `getObjects:andKeys:count` method for
`NSDictionary`:

```go
	nst := ns.NSStringWithGoString
        dict := ns.NSDictionaryWithObjectsForKeys(
                ns.NSArrayWithObjects(nst("obj1"),nst("obj2")),
                ns.NSArrayWithObjects(nst("key1"),nst("key2")),
        )
        os,ks := make([]ns.Id,0,5), make([]ns.Id,0,5)  // length 0, capacity 5 slices
        dict.GetObjects(&os,&ks,5) // count = 5, must be the same size or smaller than the input slice capacity
        fmt.Printf("Length of os is now %d\n",len(os)) // os and ks slices are now length = 2
        for i,k := range ks {
                fmt.Printf("-- %s -> %s\n",k.NSString(),os[i].NSString())
        }
```

Using pointers to pointers is necessary in many Core Foundation situations
where you need to get an error message out of a function or method, for example
in `[NSString stringWithContentsOfURL...]`:

```go
        err := make([]ns.NSError,1)
        n1 = ns.NSStringWithContentsOfURLEncoding(ns.NSURLWithGoString("htttypo://example.com"),0,&err)
        fmt.Printf("err: %s\n",err[0].LocalizedDescription())
//err: The file couldn’t be opened because URL type htttypo isn’t supported.
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
`func(ns.Id) bool` parameter that returns `true` to continue and `false` to
stop the enumeration.

```go
a := ns.NSArrayWithObjects(o1,o2,o3)
a.ObjectEnumerator().ForIn(func (o ns.Id) bool {
	switch {
	case o.IsKindOfClass(ns.NSStringClass()):
		fmt.Println(o.NSString().UTF8String())
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

NSWrap translates C `enum` values into Go constants. The enums you need are
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


## Memory management

You can call `Retain()`, `Release()` and `Autorelease()` on any object.

All allocation functions generated by NSWrap call `autorelease` before they
return an object. If you are not working in an environment (such as an
Application Delegate callback) that provides an autorelease pool, you can
create your own:

* Work directly with NSAutoreleasePool objects

```go
swamp := ns.NSAutoreleasePoolAlloc().Init()
del := ns.AppDelegateAlloc()
menu := ns.NSMenuAlloc().InitWithTitle(nst("Main"))
str := ns.NSStringWithGoString("these objects will be automatically deallocated when swamp is drained.")
...
swamp.Drain()
```

* ...or use the AutoreleasePool() helper function

NSWrap provides a helper function that can be passed a `func()` with no
parameters or return value. It is conventient to give it an anonymous function
and write your code in line, just like you would if you were using an
`@autoreleasepool { }` block.

```go
ns.AutoreleasePool(func() {
	a := MyObjectAlloc().Init()
	b := MyOtherObjectAlloc().Init()
	...
})
```

You will need to make sure `NSAutoreleasePool` is included in the `classes`
directive in your configuration file before working with
`NSAutoreleasePool` objects or the `AutoreleasePool` helper function.

Memory management seems to work but there ought to be a comprehensive
tests before anyone should feel confident with it.

## Delegates

The `delegates` directive in `nswrap.yaml` creates a new Objective-C
class and associated Go wrapper functions. For example, the following
configuration file creates a class called `CBDelegate` that implements
the Objective-C protocols `CBCentralManagerDelegate` and
`CBPeripheralDelegate`, along with the Go code you need to allocate
and use instances of the class.

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

The generated delegate inherits from NSObject and is identified as implementing
the protocols specified in `nswrap.yaml`.

When a delegate is activated and one of the callback methods named in the
configuration file is called, the delegate will call back into an exported Go 
function. If a user-defined callback function has been specified,
it will be called with all of its parameters converted to their Go type
equivalents. User-defined callbacks are registered by calling a function
with the method name in TitleCase + `Callback`, so in the example above,
call `ns.CentralManagerDidUpdateStateCallback(...)` with the name of your
callback function to register to receive notifications when your central
manager updates its state.

The code in `examples/bluetooth` implements a working Bluetooth Low Energy
heart rate monitor entirely in Go.

The following Go code creates a CBDelegate object in Go,
registers a callback for `centralManagerDidUpdateState`, allocates
a CBCentralManager object, and installs our delegate:

```go
func cb(c ns.CBCentralManager) {
	...
}

func main() {
	...
	del := ns.CBDelegateAlloc()
	del.CentralManagerDidUpdateStateCallback(cb)
	cm := ns.CBCentralManagerAlloc().InitWithDelegateQueue(del,queue)
```

When you provide user-defined callback functions, you will need to specify
them with exactly the right type,
matching NSWrap's generated Go wrapper types for the callback function and
the Go types for all of its parameters. If `go build` fails, the error
messages will point you in the right direction.

```
$ go build
./main.go:127:43: cannot use didFinishLaunching (type func(ns.NSNotification, bool)) as type
func(ns.NSNotification) in argument to del.ApplicationDidFinishLaunchingCallback
```
In the above example, build failed because an extra `bool` parameter was
included in the callback function. The compiler is telling you that the right
type for the callback is `func(ns.NSNotification)` with no return value.

## Working with AppKit

You can wrap the AppKit framework classes and create an NSApplication
Delegate. This allows you to build a Cocoa app entirely in Go.

Because AppKit uses thread local storage, you will need to make sure all
calls into it are done from the main OS thread. This can be a challenge in
Go even though runtime.LockOSThread() is supposed to provide
this functionality. Good luck with that!

This is actually a full working example:

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
	"git.wow.st/gmp/nswrap/examples/app/ns" // point to your NSWrap output directory
)

func didFinishLaunching(n ns.NSNotification) {
	fmt.Println("Go: did finish launching!")
}

func shouldTerminate(s ns.NSApplication) ns.BOOL {
	return 1
}

func main() {
	runtime.LockOSThread()
	a := ns.NSApplicationSharedApplication()
	a.SetActivationPolicy(ns.NSApplicationActivationPolicyRegular)
	del := ns.AppDelegateAlloc()
	del.ApplicationDidFinishLaunchingCallback(didFinishLaunching)
	del.ApplicationShouldTerminateAfterLastWindowClosedCallback(shouldTerminate)
	a.SetDelegate(del)

	win := ns.NSWindowAlloc().InitWithContentRectStyleMask(
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

Pretty simple right? Not really, NSWrap just generated almost 15,000 lines of
code. See `examples/app` for a slightly more complex example with working
menus and visual format-based auto layout.

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
#       |--note the hyphen indicating that this is an instance method
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

You can use subclasses to define new AppKit controls with configurable
callbacks. For example, lets make an `NSButton` that calls back into Go when
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

Later on you can add the your new button to a view and tell Cocoa where to lay
it out. It's all a little verbose, but that's because for some reason you
decided to write Objective-C code in Go.

# Limitations

## Blocks

NSWrap does not support methods or functions that take C functions or blocks
as parameters or return values.

# Why?

Um, I was trying to make a nice modern Go binding for CoreBluetooth on MacOS
and got carried away.

# Acknowledgements

This work was inspired by Maxim's excellent
[c-for-go](https://github.com/xlab/c-for-go). Much of the
infrastructure was lifted from Elliot Chance's equally excellent
[c2go](https://github.com/elliotchance/c2go). Kiyoshi Murata's
post on [coderwall.com](https://coderwall.com/p/l9jr5a/accessing-cocoa-objective-c-from-go-with-cgo)
was an essential piece of inspiration.

The combinatorial Objective-C type parsers are mine as are the 
Objective-C and Go code generators, so this is where you will find
all of the bugs.
