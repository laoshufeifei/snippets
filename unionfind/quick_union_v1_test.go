package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickUnionV1(t *testing.T) {
	test := assert.New(t)

	union := newQuickUnionV1(5)
	test.Equal(union.Find(0), 0)
	test.Equal(union.Find(1), 1)
	test.False(union.IsSameUnion(0, 1))

	// 0    2 ...
	// |
	// 1
	union.Union(1, 0)
	test.Equal(union.Find(1), 0)
	test.True(union.IsSameUnion(0, 1))
	test.False(union.IsSameUnion(1, 2))

	//   2  3...
	// / |
	// 0 1
	union.Union(1, 2)
	test.Equal(union.Find(0), 2)
	test.Equal(union.Find(1), 2)
	test.True(union.IsSameUnion(1, 2))
	test.True(union.IsSameUnion(0, 2))

	//   3   4
	// / | \
	// 0 1  2
	union.Union(0, 3)
	test.True(union.IsSameUnion(0, 3))
	test.True(union.IsSameUnion(1, 3))
	test.True(union.IsSameUnion(2, 3))
	test.False(union.IsSameUnion(2, 4))

	union.Union(3, 4)
	test.True(union.IsSameUnion(2, 4))
}
