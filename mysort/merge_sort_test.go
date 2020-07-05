package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 9, 4, 3, 7, 6, 0}
	mergeSort(ints)
	test.Equal(ints, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}
