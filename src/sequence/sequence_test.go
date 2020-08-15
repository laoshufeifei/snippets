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

func TestKMPV1(t *testing.T) {
	test := assert.New(t)

	test.Equal(kmpSearchV1("DABCDABCFACBA", "ABCDABCE"), -1)
	test.Equal(kmpSearchV1("DABCDABCFACBABCDABCE", "ABCDABCE"), 12)

	test.Equal(kmpSearchV1("0123abc", "abcd"), -1)
	test.Equal(kmpSearchV1("0123abc", "1234"), -1)
	test.Equal(kmpSearchV1("0123abc", "01234"), -1)

	test.Equal(kmpSearchV1("0123abc", "0"), 0)
	test.Equal(kmpSearchV1("0123abc", "01"), 0)
	test.Equal(kmpSearchV1("0123abc", "012"), 0)
	test.Equal(kmpSearchV1("0123abc", "0123"), 0)

	test.Equal(kmpSearchV1("0123abc", "c"), 6)
	test.Equal(kmpSearchV1("0123abc", "bc"), 5)
	test.Equal(kmpSearchV1("0123abc", "abc"), 4)

	test.Equal(kmpSearchV1("0123abc", "12"), 1)
	test.Equal(kmpSearchV1("0123abc", "23"), 2)
	test.Equal(kmpSearchV1("0123abc", "3a"), 3)

	test.Equal(kmpSearchV1("123", "1234"), -1)
	test.Equal(kmpSearchV2("aaabaaaab", "aaaab"), 4)
}

func TestKMPV1Next(t *testing.T) {
	test := assert.New(t)

	test.Equal(initNextV1("abacababd"), []int{-1, 0, 0, 1, 0, 1, 2, 3, 2})
	test.Equal(initNextV1("abcdabce"), []int{-1, 0, 0, 0, 0, 1, 2, 3})
	test.Equal(initNextV1("abcdabd"), []int{-1, 0, 0, 0, 0, 1, 2})
	test.Equal(initNextV1("aaaab"), []int{-1, 0, 1, 2, 3})
}

func TestKMPV2(t *testing.T) {
	test := assert.New(t)

	test.Equal(kmpSearchV2("DABCDABCFACBA", "ABCDABCE"), -1)
	test.Equal(kmpSearchV2("DABCDABCFACBABCDABCE", "ABCDABCE"), 12)

	test.Equal(kmpSearchV2("0123abc", "abcd"), -1)
	test.Equal(kmpSearchV2("0123abc", "1234"), -1)
	test.Equal(kmpSearchV2("0123abc", "01234"), -1)

	test.Equal(kmpSearchV2("0123abc", "0"), 0)
	test.Equal(kmpSearchV2("0123abc", "01"), 0)
	test.Equal(kmpSearchV2("0123abc", "012"), 0)
	test.Equal(kmpSearchV2("0123abc", "0123"), 0)

	test.Equal(kmpSearchV2("0123abc", "c"), 6)
	test.Equal(kmpSearchV2("0123abc", "bc"), 5)
	test.Equal(kmpSearchV2("0123abc", "abc"), 4)

	test.Equal(kmpSearchV2("0123abc", "12"), 1)
	test.Equal(kmpSearchV2("0123abc", "23"), 2)
	test.Equal(kmpSearchV2("0123abc", "3a"), 3)

	test.Equal(kmpSearchV2("123", "1234"), -1)
	test.Equal(kmpSearchV2("aaabaaaab", "aaaab"), 4)
}

func TestKMPV2Next(t *testing.T) {
	test := assert.New(t)

	test.Equal(initNextV2("abacababd"), []int{-1, 0, -1, 1, -1, 0, -1, 3, 2})
	test.Equal(initNextV2("abcdabce"), []int{-1, 0, 0, 0, -1, 0, 0, 3})
	test.Equal(initNextV2("abcdabd"), []int{-1, 0, 0, 0, -1, 0, 2})
	test.Equal(initNextV2("aaaab"), []int{-1, -1, -1, -1, 3})
}
