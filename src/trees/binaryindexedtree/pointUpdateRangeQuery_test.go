package binaryindexedtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixSum1(t *testing.T) {
	test := assert.New(t)

	tree := newPointUpdateRangeQueryTree([]int{1, 5, 8, 6, 9, 2, 3, 7})
	test.Equal(tree.pointQuery(1), 1)
	test.Equal(tree.pointQuery(2), 5)
	test.Equal(tree.pointQuery(3), 8)
	test.Equal(tree.pointQuery(4), 6)
	test.Equal(tree.pointQuery(5), 9)
	test.Equal(tree.pointQuery(6), 2)
	test.Equal(tree.pointQuery(7), 3)
	test.Equal(tree.pointQuery(8), 7)

	test.Equal(tree.prefixSum(0), 0)
	test.Equal(tree.prefixSum(1), 1)
	test.Equal(tree.prefixSum(2), 6)
	test.Equal(tree.prefixSum(3), 14)
	test.Equal(tree.prefixSum(4), 20)
	test.Equal(tree.prefixSum(5), 29)
	test.Equal(tree.prefixSum(6), 31)
	test.Equal(tree.prefixSum(7), 34)
	test.Equal(tree.prefixSum(8), 41)
	test.Equal(tree.prefixSum(100), 41)
}

func TestRangeSum1(t *testing.T) {
	test := assert.New(t)

	tree := newPointUpdateRangeQueryTree([]int{1, 5, 8, 6, 9, 2, 3, 7})
	test.Equal(tree.rangeSum(0, 0), 0)

	test.Equal(tree.rangeSum(0, 1), 1)
	test.Equal(tree.rangeSum(1, 1), 1)
	test.Equal(tree.rangeSum(2, 2), 5)
	test.Equal(tree.rangeSum(3, 3), 8)

	test.Equal(tree.rangeSum(0, 3), 14)
	test.Equal(tree.rangeSum(1, 3), 14)
	test.Equal(tree.rangeSum(2, 3), 13)
	test.Equal(tree.rangeSum(2, 30), 40)
}

func TestUpdate1(t *testing.T) {
	test := assert.New(t)

	tree := newPointUpdateRangeQueryTree([]int{1, 5, 8, 6, 9, 2, 3, 7})
	test.Equal(tree.prefixSum(1), 1)
	test.Equal(tree.prefixSum(2), 6)
	test.Equal(tree.prefixSum(3), 14)
	test.Equal(tree.prefixSum(4), 20)
	test.Equal(tree.prefixSum(7), 34)

	tree.pointUpdate(3, 10)

	test.Equal(tree.prefixSum(1), 1)
	test.Equal(tree.prefixSum(2), 6)
	test.Equal(tree.prefixSum(4), 30)
	test.Equal(tree.prefixSum(7), 44)
	test.Equal(tree.prefixSum(8), 51)
}
