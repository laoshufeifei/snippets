package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRabinKarp(t *testing.T) {
	test := assert.New(t)

	test.Equal(indexRabinKarp("xabcd", "abc"), 1)

	test.Equal(indexRabinKarp("xxxxbabcdab", "babcdab"), 4)
	test.Equal(indexRabinKarp("DABCDABCFACBA", "ABCDABCE"), -1)
	test.Equal(indexRabinKarp("DABCDABCFACBABCDABCE", "ABCDABCE"), 12)

	test.Equal(indexRabinKarp("0123abc", "abcd"), -1)
	test.Equal(indexRabinKarp("0123abc", "1234"), -1)
	test.Equal(indexRabinKarp("0123abc", "01234"), -1)

	test.Equal(indexRabinKarp("0123abc", "0"), 0)
	test.Equal(indexRabinKarp("0123abc", "01"), 0)
	test.Equal(indexRabinKarp("0123abc", "012"), 0)
	test.Equal(indexRabinKarp("0123abc", "0123"), 0)

	test.Equal(indexRabinKarp("0123abc", "c"), 6)
	test.Equal(indexRabinKarp("0123abc", "bc"), 5)
	test.Equal(indexRabinKarp("0123abc", "abc"), 4)

	test.Equal(indexRabinKarp("0123abc", "12"), 1)
	test.Equal(indexRabinKarp("0123abc", "23"), 2)
	test.Equal(indexRabinKarp("0123abc", "3a"), 3)

	test.Equal(indexRabinKarp("123", "1234"), -1)
	test.Equal(indexRabinKarp("aaabaaaab", "aaaab"), 4)
	test.Equal(indexRabinKarp("baaabaaabbaa", "baaabb"), 4)

	test.Equal(indexRabinKarp("abcbbacbabacab", "babac"), 7)
	test.Equal(indexRabinKarp("HERE IS A SIMPLE EXAMPLE", "EXAMPLE"), 17)
}
