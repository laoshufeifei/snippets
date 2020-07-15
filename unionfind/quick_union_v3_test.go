package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickUnionV3(t *testing.T) {
	test := assert.New(t)

	union := newQuickUnionV3(6)
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

func TestQuickUnionPathCompression(t *testing.T) {
	test := assert.New(t)
	test.Equal(1, 1)

	// 6 <- 5 <- 4 <- 3 <- 2 <- 1 <- 0
	union := newQuickUnionV3(7)
	for i := 0; i < 6; i++ {
		union.parents[i] = i + 1
	}

	test.Equal(union.FindWithPathCompression(0), 6)
	for i := 0; i < 7; i++ {
		test.Equal(union.Find(i), 6)
		test.Equal(union.parents[i], 6)
	}
}

func TestQuickUnionPathSplit(t *testing.T) {
	test := assert.New(t)
	test.Equal(1, 1)

	// 6 <- 5 <- 4 <- 3 <- 2 <- 1 <- 0
	union := newQuickUnionV3(7)
	for i := 0; i < 6; i++ {
		union.parents[i] = i + 1
	}

	// +- 4 <- 2 <- 0
	// 6
	// +- 5 <- 3 <- 1
	test.Equal(union.FindWithPathSplit(0), 6)

	test.Equal(union.parents[0], 2)
	test.Equal(union.parents[2], 4)
	test.Equal(union.parents[4], 6)

	test.Equal(union.parents[1], 3)
	test.Equal(union.parents[3], 5)
	test.Equal(union.parents[5], 6)
}

func TestQuickUnionPathHalf(t *testing.T) {
	test := assert.New(t)
	test.Equal(1, 1)

	// 6 <- 5 <- 4 <- 3 <- 2 <- 1 <- 0
	union := newQuickUnionV3(7)
	for i := 0; i < 6; i++ {
		union.parents[i] = i + 1
	}

	// +--- 5
	// 6
	// +    +--- 3
	// +--- 4
	//      +    +--- 1
	//      +--- 2
	//           +--- 0
	test.Equal(union.FindWithPathHalf(0), 6)

	test.Equal(union.parents[0], 2)
	test.Equal(union.parents[1], 2)
	test.Equal(union.parents[2], 4)
	test.Equal(union.parents[3], 4)
	test.Equal(union.parents[4], 6)
	test.Equal(union.parents[5], 6)
}
