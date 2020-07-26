package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickUnionV2(t *testing.T) {
	test := assert.New(t)

	union := newQuickUnionV2(6)
	test.Equal(union.Find(0), 0)
	test.Equal(union.Find(1), 1)
	test.False(union.IsSameUnion(0, 1))
	test.Equal(union.rank[0], 1)

	// 0    2 3 4 5
	// |
	// 1
	union.Union(1, 0)
	test.Equal(union.Find(1), 0)
	test.True(union.IsSameUnion(0, 1))
	test.False(union.IsSameUnion(1, 2))
	test.Equal(union.rank[0], 2)

	// 0    3 4 5
	// | \
	// 1  2
	union.Union(1, 2)
	test.Equal(union.Find(0), 0)
	test.Equal(union.Find(1), 0)
	test.True(union.IsSameUnion(1, 2))
	test.True(union.IsSameUnion(0, 2))
	test.Equal(union.rank[0], 2)
	test.Equal(union.rank[1], 1)
	test.Equal(union.rank[2], 1)

	// 0          5
	// | \      / |
	// 1  2    3  4
	union.Union(3, 5)
	union.Union(4, 3)
	test.Equal(union.Find(3), 5)
	test.Equal(union.Find(4), 5)
	test.Equal(union.rank[5], 2)

	//            5
	//          / | \
	//         0  3  4
	//        / \
	//       1   2
	union.Union(2, 3)
	test.Equal(union.rank[5], 3)
	test.True(union.IsSameUnion(2, 4))
}
