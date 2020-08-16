package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrustForce1(t *testing.T) {
	test := assert.New(t)

	test.Equal(brustForceSearch1("0123abc", "abcd"), -1)
	test.Equal(brustForceSearch1("0123abc", "1234"), -1)
	test.Equal(brustForceSearch1("0123abc", "01234"), -1)

	test.Equal(brustForceSearch1("0123abc", "0"), 0)
	test.Equal(brustForceSearch1("0123abc", "01"), 0)
	test.Equal(brustForceSearch1("0123abc", "012"), 0)
	test.Equal(brustForceSearch1("0123abc", "0123"), 0)

	test.Equal(brustForceSearch1("0123abc", "c"), 6)
	test.Equal(brustForceSearch1("0123abc", "bc"), 5)
	test.Equal(brustForceSearch1("0123abc", "abc"), 4)

	test.Equal(brustForceSearch1("0123abc", "12"), 1)
	test.Equal(brustForceSearch1("0123abc", "23"), 2)
	test.Equal(brustForceSearch1("0123abc", "3a"), 3)

	test.Equal(brustForceSearch1("123", "1234"), -1)
	test.Equal(brustForceSearch1("DABCDABCFACBA", "ABCDABCE"), -1)
	test.Equal(brustForceSearch1("DABCDABCFACBABCDABCE", "ABCDABCE"), 12)
}

func TestBrustForce2(t *testing.T) {
	test := assert.New(t)

	test.Equal(brustForceSearch2("0123abc", "abcd"), -1)
	test.Equal(brustForceSearch2("0123abc", "1234"), -1)
	test.Equal(brustForceSearch2("0123abc", "01234"), -1)

	test.Equal(brustForceSearch2("0123abc", "0"), 0)
	test.Equal(brustForceSearch2("0123abc", "01"), 0)
	test.Equal(brustForceSearch2("0123abc", "012"), 0)
	test.Equal(brustForceSearch2("0123abc", "0123"), 0)

	test.Equal(brustForceSearch2("0123abc", "c"), 6)
	test.Equal(brustForceSearch2("0123abc", "bc"), 5)
	test.Equal(brustForceSearch2("0123abc", "abc"), 4)

	test.Equal(brustForceSearch2("0123abc", "12"), 1)
	test.Equal(brustForceSearch2("0123abc", "23"), 2)
	test.Equal(brustForceSearch2("0123abc", "3a"), 3)

	test.Equal(brustForceSearch2("123", "1234"), -1)
	test.Equal(brustForceSearch2("DABCDABCFACBA", "ABCDABCE"), -1)
	test.Equal(brustForceSearch2("DABCDABCFACBABCDABCE", "ABCDABCE"), 12)
}
