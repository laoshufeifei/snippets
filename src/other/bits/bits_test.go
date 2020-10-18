package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerTwo(t *testing.T) {
	test := assert.New(t)

	test.Equal(cellToPowerOfTwo(0), uint64(1))
	test.Equal(cellToPowerOfTwo(1), uint64(1))
	test.Equal(cellToPowerOfTwo(2), uint64(2))
	test.Equal(cellToPowerOfTwo(3), uint64(4))
	test.Equal(cellToPowerOfTwo(4), uint64(4))
	test.Equal(cellToPowerOfTwo(5), uint64(8))
	test.Equal(cellToPowerOfTwo(100), uint64(128))
	test.Equal(cellToPowerOfTwo(129), uint64(256))

	// 0x1fff: 0000 0001  1111 1111  ....
	// 0x2000: 0000 0010  0000 0000  ....
	test.Equal(cellToPowerOfTwo(0x1fff), uint64(0x2000))
}
