package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntComparator(t *testing.T) {
	test := assert.New(t)

	test.Equal(IntComparator(0, 0), 0)
	test.Equal(IntComparator(1, 1), 0)
	test.Equal(IntComparator(1, 2), -1)
	test.Equal(IntComparator(2, 1), 1)
	test.Equal(IntComparator(-2, -1), -1)
}

func TestStringComparator(t *testing.T) {
	test := assert.New(t)

	test.Equal(StringComparator("abc", "abc"), 0)
	test.Equal(StringComparator("abc", "ab"), 1)
	test.Equal(StringComparator("ab", "abc"), -1)

	test.Equal(StringComparator("ad", "abc"), 1)
	test.Equal(StringComparator("aa", "ab"), -1)

	test.Equal(StringComparator("aaa", "aa"), 1)
}
