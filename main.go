package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strings"

	"git.wow.st/gmp/nswrap/ast"
	"git.wow.st/gmp/nswrap/types"
	"git.wow.st/gmp/nswrap/wrap"

	"github.com/a8m/envsubst"
	"gopkg.in/yaml.v2"
)

var Debug = false
var Profile = false

//automatically add interfaces if they are found in the input interface
//declarations
var autoadd = []string{
	"NSObject",
}

type conf struct {
	Positions      bool
	Package        string
	Inputfiles     []string
	Astfile        string
	Debugast       bool
	Classes        []string
	Functions      []string
	FunctionIgnore []string
	Enums          []string
	Delegates      map[string]map[string][]string
	Subclasses     map[string]map[string][]string
	Frameworks     []string
	Libraries      []string
	Frameworkdirs  []string
	Imports        []string
	Sysimports     []string
	Pragma         []string
	Typesubs       map[string]string
	Vaargs         int
	Clang          string
	//Arc flag for debugging only, builds will break
	Arc         bool
	Autorelease bool
	Nogc        bool
}

var Config conf

func readAST(data []byte) []string {
	return strings.Split(string(data), "\n")
}

type treeNode struct {
	indent int
	node   ast.Node
}

func printLinesWithContext(lines []string, i int) {
	b := i - 3
	if b < 0 {
		b = 0
	}
	var flag string
	for x := b; (x < b+6) && (x < len(lines)); x++ {
		if x == i {
			flag = "--> "
		} else {
			flag = "    "
		}
		fmt.Printf("%s%s\n", flag, lines[x])
	}
}

func convertLinesToNodes(lines []string) []treeNode {
	nodes := make([]treeNode, len(lines))
	var counter int
	unknowns := 0
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		// It is tempting to discard null AST nodes, but these may
		// have semantic importance: for example, they represent omitted
		// for-loop conditions, as in for(;;).
		line = strings.Replace(line, "<<<NULL>>>", "NullStmt", 1)
		trimmed := strings.TrimLeft(line, "|\\- `")
		node := ast.Parse(trimmed)
		indentLevel := (len(line) - len(trimmed)) / 2
		nodes[counter] = treeNode{indentLevel, node}
		counter++
		if !Config.Debugast {
			continue
		}
		switch node.(type) {
		case *ast.Unknown:
			fmt.Printf("Unrecognized node:\n")
			printLinesWithContext(lines, i)
			fmt.Printf("\n")
			unknowns++
			if unknowns > 5 {
				fmt.Printf("nswrap failed due to unrecognized nodes.\n")
				os.Exit(-1)
			}
		}
	}
	nodes = nodes[0:counter]

	if Config.Debugast && unknowns > 0 {
		fmt.Printf("\nswrap failed due to unrecognized nodes.\n")
		os.Exit(-1)
	}
	return nodes
}

func convertLinesToNodesParallel(lines []string) []treeNode {
	// function f separate full list on 2 parts and
	// then each part can recursive run function f
	var f func([]string, int) []treeNode

	f = func(lines []string, deep int) []treeNode {
		deep = deep - 2
		part := len(lines) / 2

		var tr1 = make(chan []treeNode)
		var tr2 = make(chan []treeNode)

		go func(lines []string, deep int) {
			if deep <= 0 || len(lines) < deep {
				tr1 <- convertLinesToNodes(lines)
				return
			}
			tr1 <- f(lines, deep)
		}(lines[0:part], deep)

		go func(lines []string, deep int) {
			if deep <= 0 || len(lines) < deep {
				tr2 <- convertLinesToNodes(lines)
				return
			}
			tr2 <- f(lines, deep)
		}(lines[part:], deep)

		defer close(tr1)
		defer close(tr2)

		return append(<-tr1, <-tr2...)
	}

	// Parameter of deep - can be any, but effective to use
	// same amount of CPU
	return f(lines, runtime.NumCPU())
}

// buildTree converts an array of nodes, each prefixed with a depth into a tree.
func buildTree(nodes []treeNode, depth int) []ast.Node {
	if len(nodes) == 0 {
		return []ast.Node{}
	}

	// Split the list into sections, treat each section as a tree with its own
	// root.
	sections := [][]treeNode{}
	for _, node := range nodes {
		if node.indent == depth {
			sections = append(sections, []treeNode{node})
		} else {
			sections[len(sections)-1] = append(sections[len(sections)-1], node)
		}
	}

	results := []ast.Node{}
	for _, section := range sections {
		slice := []treeNode{}
		for _, n := range section {
			if n.indent > depth {
				slice = append(slice, n)
			}
		}

		children := buildTree(slice, depth+1)
		for _, child := range children {
			section[0].node.AddChild(child)
		}
		results = append(results, section[0].node)
	}

	return results
}

func matches(x string, rs []string) bool {
	for _, r := range rs {
		if m, _ := regexp.MatchString("^"+r+"$", x); m {
			return true
		}
	}
	return false
}

// Start begins transpiling an input file.
func Start() (err error) {
	astPP := []byte{}
	if Config.Astfile != "" {
		fmt.Printf("Reading ast file %s\n", Config.Astfile)
		_, err = os.Stat(Config.Astfile)
		if err != nil {
			return fmt.Errorf("Input AST file %s not found", Config.Astfile)
		}
		astPP, err = ioutil.ReadFile(Config.Astfile)
		if err != nil {
			return err
		}
	} else {
		for _, in := range Config.Inputfiles {
			_, err = os.Stat(in)
			if err != nil {
				return fmt.Errorf("Input file %s is not found", in)
			}
		}

		// Generate AST
		cargs := []string{"-xobjective-c", "-Xclang", "-ast-dump",
			"-fsyntax-only", "-fno-color-diagnostics"}
		if Config.Arc {
			cargs = append(cargs, "-fobjc-arc")
		}
		for _, f := range Config.Frameworkdirs {
			cargs = append(cargs, "-F"+f)
		}
		cargs = append(cargs, Config.Inputfiles...)
		fmt.Printf("Generating AST\n")
		clang := "clang"
		if Config.Clang != "" {
			clang = Config.Clang
		}
		fmt.Printf("%s %s\n", clang, strings.Join(cargs, " "))
		astPP, err = exec.Command(clang, cargs...).Output()
		if err != nil {
			// If clang fails it still prints out the AST, so we have to run it
			// again to get the real error.
			//errBody, _ := exec.Command("clang", cargs...).CombinedOutput()
			var txt string
			switch x := err.(type) {
			case *exec.ExitError:
				txt = string(x.Stderr)
			default:
				txt = err.Error()
			}
			fmt.Printf("clang failed:\n%s\n", txt)
			os.Exit(-1)
		}
	}

	lines := readAST(astPP)

	// Converting to nodes
	fmt.Printf("Building nodes\n")
	if Config.Positions {
		ast.TrackPositions = true
	}
	wrap.Gogc = true

	if Config.Arc {
		wrap.Arc = true
		wrap.Gogc = false
	}
	if Config.Autorelease {
		wrap.Autorelease = true
		wrap.Gogc = false
	}
	if Config.Nogc {
		wrap.Gogc = false
	}

	//NOTE: converting in parallel is slower on my system
	//nodes := convertLinesToNodesParallel(lines)
	nodes := convertLinesToNodes(lines)

	// build tree
	fmt.Printf("Assembling tree\n")
	tree := buildTree(nodes, 0)
	w := wrap.NewWrapper(Debug)
	w.Package = Config.Package
	w.Frameworks = Config.Frameworks
	w.Libraries = Config.Libraries
	w.Frameworkdirs = Config.Frameworkdirs
	w.Import(Config.Imports)
	w.SysImport(Config.Sysimports)
	w.Pragmas = Config.Pragma
	w.Delegate(Config.Delegates)
	w.Subclass(Config.Subclasses)
	types.Typesubs = Config.Typesubs
	if Config.Vaargs == 0 {
		Config.Vaargs = 16
	}
	w.Vaargs = Config.Vaargs
	for _, u := range tree {
		fmt.Printf("--processing translation unit\n")
		for _, n := range u.Children() {
			switch x := n.(type) {
			case *ast.ObjCInterfaceDecl:
				w.AddInterface(x)
				for _, ss := range Config.Subclasses {
					if sc, ok := ss["superclass"]; ok {
						if matches(x.Name, sc) {
							Config.Classes = append(Config.Classes, x.Name)
						}
					}
				}
				if matches(x.Name, autoadd) {
					Config.Classes = append(Config.Classes, x.Name)
				}
			case *ast.ObjCCategoryDecl:
				w.AddCategory(x)
			case *ast.TypedefDecl:
				if !x.IsImplicit {
					w.AddTypedef(x.Name, x.Type)
				}
			case *ast.FunctionDecl:
				if matches(x.Name, Config.Functions) &&
					!matches(x.Name, Config.FunctionIgnore) {
					w.AddFunction(x)
				}
			case *ast.ObjCProtocolDecl:
				w.AddProtocol(x)
			case *ast.EnumDecl:
				w.AddEnum(x, Config.Enums)
			}
		}
	}
	w.Wrap(Config.Classes)
	return nil
}

func main() {
	if Profile {
		f1, err := os.Create("cpuprofile.pprof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f1.Close()
		if err := pprof.StartCPUProfile(f1); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	conffile := "nswrap.yaml"
	if len(os.Args) > 1 {
		conffile = os.Args[1]
	}
	confbytes, err := ioutil.ReadFile(conffile)
	if err != nil {
		fmt.Printf("%s\n\nFATAL ERROR: Configuration file not found (default: nswrap.yaml)\n", err)
		os.Exit(-1)
	}
	confbytes, err = envsubst.Bytes(confbytes)
	if err != nil {
		fmt.Printf("FATAL ERROR: Shell string variable substitution faled: %s\n", err)
		os.Exit(-1)
	}
	if err = yaml.UnmarshalStrict(confbytes, &Config); err != nil {
		fmt.Printf("Cannot decode config file nswrap.yaml. %s\n", err)
		os.Exit(-1)
	}
	ast.Debug = Config.Debugast
	if err := Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}
	if Profile {
		f2, err := os.Create("memprofile.pprof")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f2.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f2); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
