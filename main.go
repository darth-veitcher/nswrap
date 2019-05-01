package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
	"gitlab.wow.st/gmp/nswrap/ast"
	"gitlab.wow.st/gmp/nswrap/types"
	"gitlab.wow.st/gmp/nswrap/wrap"
)

var Debug = false

type conf struct {
	Package string
	InputFiles []string
	Classes []string
	Functions []string
	Imports []string
	SysImports []string
	Pragma []string
	VaArgs int
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
	// 1. Compile it first (checking for errors)
	for _, in := range Config.InputFiles {
		_, err := os.Stat(in)
		if err != nil {
			return fmt.Errorf("Input file %s is not found", in)
		}
	}

	// 2. Preprocess NOT DONE

	// 3. Generate JSON from AST
	cargs := []string{"-xobjective-c", "-Xclang", "-ast-dump",
			"-fsyntax-only","-fno-color-diagnostics"}
	cargs = append(cargs,Config.InputFiles...)
	astPP, err := exec.Command("clang",cargs...).Output()
	if err != nil {
		// If clang fails it still prints out the AST, so we have to run it
		// again to get the real error.
//		errBody, _ := exec.Command("clang", cargs...).CombinedOutput()

		panic("clang failed: " + err.Error() + ":\n\n")
	}

	lines := readAST(astPP)

	// Converting to nodes
	nodes := convertLinesToNodesParallel(lines)

	// build tree
	tree := buildTree(nodes, 0)
	//unit := tree[0]
	w := wrap.NewWrapper(Debug)
	w.Package = Config.Package
	w.Import(Config.Imports)
	w.SysImport(Config.SysImports)
	w.Pragma(Config.Pragma)
	if Config.VaArgs == 0 {
		Config.VaArgs = 16
	}
	w.VaArgs = Config.VaArgs
	for _, u := range tree {
		for _, n := range(u.Children()) {
			switch x := n.(type) {
			case *ast.ObjCInterfaceDecl:
				w.AddInterface(x)
			case *ast.ObjCCategoryDecl:
				w.AddCategory(x)
			case *ast.TypedefDecl:
				types.AddTypedef(x.Name,x.Type)
			case *ast.FunctionDecl:
				if matches(x.Name,Config.Functions) {
					w.AddFunction(x)
				}
			}
		}
	}
	w.Wrap(Config.Classes)
	return nil
}

func main() {
	if _, err := toml.DecodeFile("nswrap.toml",&Config); err != nil {
		fmt.Printf("Cannot open config file nswrap.toml.\n")
		os.Exit(-1)
	}
	if err := Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}
}
