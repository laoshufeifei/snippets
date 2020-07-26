package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickUnion(t *testing.T) {
	test := assert.New(t)
	test.True(true)

	union := NewQuickUnion(6)
	union.Push("a", "b", "c", "d", "e", "f")
	test.False(union.IsSameUnion("a", "b"))
	test.False(union.IsSameUnion("a", "c"))
	test.False(union.IsSameUnion("c", "d"))

	union.Union("a", "b")
	test.True(union.IsSameUnion("a", "b"))

	union.Union("b", "c")
	test.True(union.IsSameUnion("a", "c"))
}
