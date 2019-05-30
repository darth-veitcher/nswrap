package ast

import (
	"reflect"
	"testing"

	"git.wow.st/gmp/nswrap/util"
)

func TestArrayFiller(t *testing.T) {
	expected := &ArrayFiller{
		ChildNodes: []Node{},
	}
	actual := Parse(`array filler`)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%s", util.ShowDiff(formatMultiLine(expected),
			formatMultiLine(actual)))
	}
}
