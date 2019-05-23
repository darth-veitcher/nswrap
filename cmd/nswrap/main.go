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

	"gopkg.in/yaml.v2"
	"gitlab.wow.st/gmp/nswrap/ast"
	"gitlab.wow.st/gmp/nswrap/wrap"
)

var Debug = false
var Profile = false

type conf struct {
	Positions bool
	Package string
	Inputfiles []string
	Classes []string
	Functions []string
	Enums []string
	Delegates map[string]map[string][]string
	Subclasses map[string]map[string][]string
	Frameworks []string
	Imports []string
	Sysimports []string
	Pragma []string
	Vaargs int
}

var Config conf

func readAST(data []byte) []string {
	return strings.Split(string(data), "\n")
}

type treeNode struct {
	indent int
	node   ast.Node
}

func convertLinesToNodes(lines []string) []treeNode {
	nodes := make([]treeNode, len(lines))
	var counter int
	for _, line := range lines {
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
	}
	nodes = nodes[0:counter]

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
	for _,r := range rs {
		if m, _ := regexp.MatchString("^" + r + "$",x); m {
			return true
		}
	}
	return false
}

// Start begins transpiling an input file.
func Start() (err error) {
	for _, in := range Config.Inputfiles {
		_, err := os.Stat(in)
		if err != nil {
			return fmt.Errorf("Input file %s is not found", in)
		}
	}

	// Generate AST
	cargs := []string{"-xobjective-c", "-Xclang", "-ast-dump",
			"-fsyntax-only","-fno-color-diagnostics"}
	cargs = append(cargs,Config.Inputfiles...)
	fmt.Printf("Generating AST\n")
	astPP, err := exec.Command("clang",cargs...).Output()
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

	lines := readAST(astPP)

	// Converting to nodes
	fmt.Printf("Building nodes\n")
	if Config.Positions {
		ast.TrackPositions = true
	}
	//NOTE: converting in parallel is slower on my system
	//nodes := convertLinesToNodesParallel(lines)
	nodes := convertLinesToNodes(lines)

	// build tree
	fmt.Printf("Assembling tree\n")
	tree := buildTree(nodes, 0)
	w := wrap.NewWrapper(Debug)
	w.Package = Config.Package
	w.Frameworks(Config.Frameworks)
	w.Import(Config.Imports)
	w.SysImport(Config.Sysimports)
	w.Pragma(Config.Pragma)
	w.Delegate(Config.Delegates)
	w.Subclass(Config.Subclasses)
	if Config.Vaargs == 0 {
		Config.Vaargs = 16
	}
	w.Vaargs = Config.Vaargs
	for _, u := range tree {
		fmt.Printf("--processing translation unit\n")
		for _, n := range(u.Children()) {
			switch x := n.(type) {
			case *ast.ObjCInterfaceDecl:
				w.AddInterface(x)
				for _,ss := range Config.Subclasses {
					for ps,_ := range ss {
						if matches(x.Name,[]string{ps}) {
							Config.Classes = append(Config.Classes,x.Name)
						}
					}
				}
			case *ast.ObjCCategoryDecl:
				w.AddCategory(x)
			case *ast.TypedefDecl:
				w.AddTypedef(x.Name,x.Type)
			case *ast.FunctionDecl:
				if matches(x.Name,Config.Functions) {
					w.AddFunction(x)
				}
			case *ast.ObjCProtocolDecl:
				w.AddProtocol(x)
			case *ast.EnumDecl:
				w.AddEnum(x,Config.Enums)
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

	confbytes, err := ioutil.ReadFile("nswrap.yaml")
	if err != nil {
		fmt.Printf("Cannot open config file nswrap.yaml. %s\n",err)
		os.Exit(-1)
	}
	if err = yaml.Unmarshal(confbytes,&Config); err != nil {
		fmt.Printf("Cannot decode config file nswrap.yaml. %s\n",err)
		os.Exit(-1)
	}
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
