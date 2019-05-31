package ast

import (
	"testing"
)

func TestVectorType(t *testing.T) {
	nodes := map[string]Node{
		`0x7faa18f2d520 '__attribute__((__vector_size__(1 * sizeof(long long)))) long long' 1`:
		&VectorType{
			Addr:         0x7faa18f2d520,
			Type:         `__attribute__((__vector_size__(1 * sizeof(long long)))) long long`,
			Length:       1,
			ChildNodes:   []Node{},
		},
		`0x7fca42b88f30 '__attribute__((__vector_size__(16 * sizeof(signed char)))) signed char' 16`:
		&VectorType{
			Addr:         0x7fca42b88f30,
			Type:         `__attribute__((__vector_size__(16 * sizeof(signed char)))) signed char`,
			Length:       16,
			ChildNodes:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}
