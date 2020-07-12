package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 3, 4, 9, 7, 6}

	sorter := newQuickSorter()
	sorter.sortImple(ints)
	test.Equal(ints, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}
