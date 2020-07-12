package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 3, 4, 9, 0, 7, 6}

	v1 := newBubbleSorterV1()
	v1.sortImple(ints)
	test.Equal(ints, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestBubbleSort2(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 3, 4, 9, 0, 7, 6}

	v2 := newBubbleSorterV2()
	v2.sortImple(ints)
	test.Equal(ints, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestBubbleSort3(t *testing.T) {
	test := assert.New(t)
	ints1 := []int{3, 1, 2, 4, 5, 6, 7, 8, 9}
	v1 := newBubbleSorterV1()
	v1.sortImple(ints1)

	ints2 := []int{3, 1, 2, 4, 5, 6, 7, 8, 9}
	v2 := newBubbleSorterV2()
	v2.sortImple(ints2)
	test.Less(v2.cmpCounter, v1.cmpCounter)
}
