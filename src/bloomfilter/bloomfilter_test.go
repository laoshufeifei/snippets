package bloomfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomFilter(t *testing.T) {
	test := assert.New(t)

	test.Equal(1, 1)

	b := New(0.01, 1000)
	b.Put("abc")
	test.True(b.Contains("abc"))
}
