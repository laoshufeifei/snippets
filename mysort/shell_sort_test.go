package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellSortV1(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 3, 4, 9, 7, 6}
	shellSortV1(ints)
	test.Equal(ints, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestShellSortV2(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 3, 4, 9, 7, 6}
	shellSort(ints)
	test.Equal(ints, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestShellSort(t *testing.T) {
	test := assert.New(t)
	ints := []int{10, 5, 2, 7, 8, 16, 15, 4, 3, 19, 1, 20, 17, 9, 6, 18, 11, 14}
	shellSortV1(ints)
	test.Equal(ints, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 14, 15, 16, 17, 18, 19, 20})
}

func TestShellSort2(t *testing.T) {
	test := assert.New(t)
	ints := []int{10, 5, 2, 7, 8, 16, 15, 4, 3, 19, 1, 20, 17, 9, 6, 18, 11, 14}
	shellSort(ints)
	test.Equal(ints, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 14, 15, 16, 17, 18, 19, 20})
}
