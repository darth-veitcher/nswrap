package ast

import (
	"testing"
)

func TestRemoveQuotes(t *testing.T) {
	t.Run("RemoveQuotes1", func(t *testing.T) {
		if removeQuotes(`''`) != "" {
			t.Errorf("Empty single quoted string test failed\n")
		}
	})
	t.Run("RemoveQuotes2", func(t *testing.T) {
		if removeQuotes(`"hi there"`) != `hi there` {
			t.Errorf("Double quoted string test failed: %s -> %s\n", `"hi there"`, removeQuotes(`"hi there"`))
		}
	})
	t.Run("TypesTree", func(t *testing.T) {
		if typesTree(nil, 0) != "" {
			t.Errorf(`typesTree(nil,0) did not return ""`)
		}
	})
}
