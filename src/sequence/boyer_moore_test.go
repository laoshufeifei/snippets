package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadCharacter(t *testing.T) {
	test := assert.New(t)

	bcs := initBadCharacters("example")
	test.Equal(bcs['e'], 6)
	test.Equal(bcs['l'], 1)
	test.Equal(bcs['p'], 2)
	test.Equal(bcs['m'], 3)
	test.Equal(bcs['a'], 4)
	test.Equal(bcs['x'], 5)
	test.Equal(bcs['e'], 6)

	bcs2 := initBadCharacters("gcagagag")
	test.Equal(bcs2['g'], 2)
	test.Equal(bcs2['c'], 6)
	test.Equal(bcs2['a'], 1)

	bcs3 := initBadCharacters("bcababab")
	test.Equal(bcs3['a'], 1)
	test.Equal(bcs3['b'], 2)
	test.Equal(bcs3['c'], 6)
}

func TestGoodSuffixes(t *testing.T) {
	test := assert.New(t)

	test.Equal(initGoodSuffixes("bcababab"), []int{7, 7, 7, 2, 7, 4, 7, 1})
	test.Equal(initGoodSuffixes("abcd"), []int{4, 4, 4, 1})
	test.Equal(initGoodSuffixes("aaaa"), []int{1, 2, 3, 4})
}

func TestSuffixes(t *testing.T) {
	test := assert.New(t)
	test.Equal(initSuffixes("bcababab"), []int{1, 0, 0, 2, 0, 4, 0, 8})

	ss := []string{
		"aaaaaaa",
		"aaaaaab",
		"elemele",
		"bcababab",
		"aaaaabaaaa",
	}
	for _, s := range ss {
		test.Equal(initSuffixes(s), initSuffixes2(s))
	}
}

func TestBooyerMoore(t *testing.T) {
	test := assert.New(t)

	test.Equal(boyerMoore("xxxxbabcdab", "babcdab"), 4)
	test.Equal(boyerMoore("DABCDABCFACBA", "ABCDABCE"), -1)
	test.Equal(boyerMoore("DABCDABCFACBABCDABCE", "ABCDABCE"), 12)

	test.Equal(boyerMoore("0123abc", "abcd"), -1)
	test.Equal(boyerMoore("0123abc", "1234"), -1)
	test.Equal(boyerMoore("0123abc", "01234"), -1)

	test.Equal(boyerMoore("0123abc", "0"), 0)
	test.Equal(boyerMoore("0123abc", "01"), 0)
	test.Equal(boyerMoore("0123abc", "012"), 0)
	test.Equal(boyerMoore("0123abc", "0123"), 0)

	test.Equal(boyerMoore("0123abc", "c"), 6)
	test.Equal(boyerMoore("0123abc", "bc"), 5)
	test.Equal(boyerMoore("0123abc", "abc"), 4)

	test.Equal(boyerMoore("0123abc", "12"), 1)
	test.Equal(boyerMoore("0123abc", "23"), 2)
	test.Equal(boyerMoore("0123abc", "3a"), 3)

	test.Equal(boyerMoore("123", "1234"), -1)
	test.Equal(boyerMoore("aaabaaaab", "aaaab"), 4)
	test.Equal(boyerMoore("baaabaaabbaa", "baaabb"), 4)

	test.Equal(boyerMoore("abcbbacbabacab", "babac"), 7)
	test.Equal(boyerMoore("HERE IS A SIMPLE EXAMPLE", "EXAMPLE"), 17)
}
