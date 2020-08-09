package bloomfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitSet(t *testing.T) {
	test := assert.New(t)

	s := newBitSet(130)
	test.True(s.length == 130)
	test.True(s.arrayLen == 3)

	test.False(s.get(0))
	test.False(s.get(63))
	test.False(s.get(64))
	test.False(s.get(129))

	test.True(s.set(0))
	test.True(s.get(0))

	test.True(s.set(8))
	test.True(s.get(8))
	test.False(s.set(8))

	test.True(s.set(128))
	test.True(s.get(128))

	test.False(s.set(1280))
	test.False(s.get(1280))
}
