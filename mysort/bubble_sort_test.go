package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 3, 4, 9, 0, 7, 6}
	bubbleSortV1(ints)
	test.Equal(ints, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestBubbleSort2(t *testing.T) {
	test := assert.New(t)
	ints := []int{3, 1, 2, 4, 5, 6, 7, 8, 9}
	bubbleSort(ints)
	test.Equal(ints, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}
