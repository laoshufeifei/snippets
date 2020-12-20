package segmenttree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixSum1(t *testing.T) {
	test := assert.New(t)

	st := newSegmentTree([]int{1, 3, 5, 7, 9, 11})
	test.Equal(st.tree[0], 36)

	st.update(4, 6)
	test.Equal(st.tree[0], 33)

	test.Equal(st.query(2, 5), 29)
}
