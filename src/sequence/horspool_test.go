package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHorspool(t *testing.T) {
	test := assert.New(t)

	test.Equal(horspool("xxxxbabcdab", "babcdab"), 4)
	test.Equal(horspool("", ""), 0)
	test.Equal(horspool("", "a"), -1)
	test.Equal(horspool("a", ""), 0)

	test.Equal(horspool("HERE IS A SIMPLE EXAMPLE", "EXAMPLE"), 17)
	test.Equal(horspool("abcbbacbabacab", "babac"), 7)

	test.Equal(horspool("DABCDABCFACBA", "ABCDABCE"), -1)
	test.Equal(horspool("DABCDABCFACBABCDABCE", "ABCDABCE"), 12)

	test.Equal(horspool("0123abc", "abcd"), -1)
	test.Equal(horspool("0123abc", "1234"), -1)
	test.Equal(horspool("0123abc", "01234"), -1)

	test.Equal(horspool("0123abc", "0"), 0)
	test.Equal(horspool("0123abc", "01"), 0)
	test.Equal(horspool("0123abc", "012"), 0)
	test.Equal(horspool("0123abc", "0123"), 0)

	test.Equal(horspool("0123abc", "c"), 6)
	test.Equal(horspool("0123abc", "bc"), 5)
	test.Equal(horspool("0123abc", "abc"), 4)

	test.Equal(horspool("0123abc", "12"), 1)
	test.Equal(horspool("0123abc", "23"), 2)
	test.Equal(horspool("0123abc", "3a"), 3)

	test.Equal(horspool("123", "1234"), -1)
	test.Equal(horspool("aaabaaaab", "aaaab"), 4)
	test.Equal(horspool("baaabaaabbaa", "baaabb"), 4)
}
