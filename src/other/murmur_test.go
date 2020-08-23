package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMurMur2_32(t *testing.T) {
	test := assert.New(t)

	test.Equal(murmurHash2_32([]byte("abc"), 1), uint32(1621425345))
	test.Equal(murmurHash2_32([]byte("abcde"), 1), uint32(3469237630))
	test.Equal(murmurHash2_32([]byte("abcdefgh01234"), 31), uint32(1980857320))
}

func TestMurMur3_32(t *testing.T) {
	test := assert.New(t)

	test.Equal(murmurHash3_32([]byte("abc"), 1), uint32(2859854335))
	test.Equal(murmurHash3_32([]byte("abcde"), 1), uint32(4289507611))
	test.Equal(murmurHash3_32([]byte("abcdefgh01234"), 31), uint32(3764697295))
}
