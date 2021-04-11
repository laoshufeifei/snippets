package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerTwo(t *testing.T) {
	test := assert.New(t)

	test.False(isPowerOfTwo(0))
	test.True(isPowerOfTwo(1))
	test.True(isPowerOfTwo(2))
	test.False(isPowerOfTwo(3))
	test.True(isPowerOfTwo(4))
	test.True(isPowerOfTwo(128))
}

func TestCellToPowerOfTwo(t *testing.T) {
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

func TestCellToPowerOfTwo2(t *testing.T) {
	test := assert.New(t)

	test.Equal(cellToPowerOfTwo2(0), uint64(1))
	test.Equal(cellToPowerOfTwo2(1), uint64(1))
	test.Equal(cellToPowerOfTwo2(2), uint64(2))
	test.Equal(cellToPowerOfTwo2(3), uint64(4))
	test.Equal(cellToPowerOfTwo2(4), uint64(4))
	test.Equal(cellToPowerOfTwo2(5), uint64(8))
	test.Equal(cellToPowerOfTwo2(100), uint64(128))
	test.Equal(cellToPowerOfTwo2(129), uint64(256))

	// 0x1fff: 0000 0001  1111 1111  ....
	// 0x2000: 0000 0010  0000 0000  ....
	test.Equal(cellToPowerOfTwo2(0x1fff), uint64(0x2000))
}

func TestCellToPowerOfTwo3(t *testing.T) {
	test := assert.New(t)

	test.Equal(cellToPowerOfTwo3(0), uint64(1))
	test.Equal(cellToPowerOfTwo3(1), uint64(1))
	test.Equal(cellToPowerOfTwo3(2), uint64(2))
	test.Equal(cellToPowerOfTwo3(3), uint64(4))
	test.Equal(cellToPowerOfTwo3(4), uint64(4))
	test.Equal(cellToPowerOfTwo3(5), uint64(8))
	test.Equal(cellToPowerOfTwo3(100), uint64(128))
	test.Equal(cellToPowerOfTwo3(129), uint64(256))

	// 0x1fff: 0000 0001  1111 1111  ....
	// 0x2000: 0000 0010  0000 0000  ....
	test.Equal(cellToPowerOfTwo3(0x1fff), uint64(0x2000))
}

func TestCellToPowerOfTwo4(t *testing.T) {
	test := assert.New(t)

	test.Equal(cellToPowerOfTwo4(0), uint64(1))
	test.Equal(cellToPowerOfTwo4(1), uint64(1))
	test.Equal(cellToPowerOfTwo4(2), uint64(2))
	test.Equal(cellToPowerOfTwo4(3), uint64(4))
	test.Equal(cellToPowerOfTwo4(4), uint64(4))
	test.Equal(cellToPowerOfTwo4(5), uint64(8))
	test.Equal(cellToPowerOfTwo4(100), uint64(128))
	test.Equal(cellToPowerOfTwo4(129), uint64(256))

	// 0x1fff: 0000 0001  1111 1111  ....
	// 0x2000: 0000 0010  0000 0000  ....
	test.Equal(cellToPowerOfTwo4(0x1fff), uint64(0x2000))
}

func TestNumberOfLeadingZeros(t *testing.T) {
	test := assert.New(t)

	test.Equal(numberOfLeadingZeros(0), uint64(64))
	test.Equal(numberOfLeadingZeros(1), uint64(63))
	test.Equal(numberOfLeadingZeros(3), uint64(62))
	test.Equal(numberOfLeadingZeros(0x8), uint64(60))
	test.Equal(numberOfLeadingZeros(0xf), uint64(60))

	test.Equal(numberOfLeadingZeros(0xffffff), uint64(40))
	test.Equal(numberOfLeadingZeros(0xffffffff), uint64(32))
	test.Equal(numberOfLeadingZeros(0xffffffffffff), uint64(16))
	test.Equal(numberOfLeadingZeros(0xffffffffffffff), uint64(8))
	test.Equal(numberOfLeadingZeros(0xfffffffffffffff), uint64(4))
	test.Equal(numberOfLeadingZeros(0xffffffffffffffff), uint64(0))
}
