package binaryindexedtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowbits(t *testing.T) {
	test := assert.New(t)

	test.Equal(lowbit(0), 0)
	test.Equal(lowbit(1), 1)
	test.Equal(lowbit(4), 4)
	test.Equal(lowbit(12), 4)
}
