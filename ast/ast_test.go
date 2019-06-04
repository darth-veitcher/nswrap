package ast

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"git.wow.st/gmp/nswrap/util"
)

type testNode struct {
	n Node
	addr Address
	pos Position
	children []Node
}

func init() {
	TrackPositions = true
}

func formatMultiLine(o interface{}) string {
	s := fmt.Sprintf("%#v", o)
	s = strings.Replace(s, "{", "{\n", -1)
	s = strings.Replace(s, ", ", "\n", -1)

	return s
}

func runNodeTest(t *testing.T, actual Node, expected testNode, i *int) {
	testName := fmt.Sprintf("Example%d", *i)
	t.Run(testName + "a", func(t *testing.T) {
		if !reflect.DeepEqual(expected.n, actual) {
			t.Errorf("%s", util.ShowDiff(formatMultiLine(expected.n),
				formatMultiLine(actual)))
		}
	})
	t.Run(testName+"b", func(t *testing.T) {
		if !reflect.DeepEqual(actual.Address(),expected.addr) {
			t.Errorf("Address mismatch")
		}
	})
	t.Run(testName+"c", func(t *testing.T) {
		if !reflect.DeepEqual(actual.Position(),expected.pos) {
			t.Errorf("Position mismatch")
		}
	})
	t.Run(testName+"d", func(t *testing.T) {
		if !reflect.DeepEqual(actual.Children(),expected.children) {
			t.Errorf("Children mismatch")
		}
	})
	t.Run(testName+"e", func(t *testing.T) {
		cs := expected.children
		node := &Unknown{}
		actual.AddChild(node)
		cs = append(cs,node)
		if !reflect.DeepEqual(actual.Children(),cs) {
			t.Errorf("Children mismatch")
		}
	})
	(*i)++
}

func runNodeTests(t *testing.T, tests map[string]testNode) {
	i := 1
	for line, expected := range tests {
		// Append the name of the struct onto the front. This would
		// make the complete line it would normally be parsing.
		name := reflect.TypeOf(expected.n).Elem().Name()
		actual := Parse(name + " " + line)

		runNodeTest(t,actual,expected,&i)
	}
}

func TestPrint(t *testing.T) {
	cond := &ConditionalOperator{}
	cond.AddChild(&ImplicitCastExpr{})
	cond.AddChild(&ImplicitCastExpr{})
	s := Atos(cond)
	if len(s) == 0 {
		t.Fatalf("Cannot convert AST tree : %#v", cond)
	}
	lines := strings.Split(s, "\n")
	var amount int
	for _, l := range lines {
		if strings.Contains(l, "ImplicitCastExpr") {
			amount++
		}
	}
	if amount != 2 {
		t.Error("Not correct design of output")
	}
}

var lines = []string{
// c2go ast sqlite3.c | head -5000 | sed 's/^[ |`-]*//' | sed 's/<<<NULL>>>/NullStmt/g' | gawk 'length > 0 {print "`" $0 "`,"}'
}

func BenchmarkParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, line := range lines {
			Parse(line)
		}
	}
}

