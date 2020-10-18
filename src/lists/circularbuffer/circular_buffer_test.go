package circularbuffer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	test := assert.New(t)

	ret := 0
	src := []byte("01234567")
	dst := make([]byte, len(src))

	ret = copy(dst, src)
	test.Equal(ret, 8)
	test.Equal(string(dst), "01234567")

	ret = copy(dst, src[4:])
	test.Equal(ret, 4)
	test.Equal(string(dst), "45674567")

	ret = copy(dst[4:], src[4:4])
	test.Equal(ret, 0)
	test.Equal(string(dst), "45674567")

	ret = copy(dst[4:5], src)
	test.Equal(ret, 1)
	test.Equal(string(dst), "45670567")

	ret = copy(dst[4:5], src[1:3])
	test.Equal(ret, 1)
	test.Equal(string(dst), "45671567")

	copy(dst, src)
	test.Equal(string(dst), "01234567")

	ret = copy(dst[2:4], src[5:])
	test.Equal(ret, 2)
	test.Equal(string(dst), "01564567")
}

func TestCapacity(t *testing.T) {
	test := assert.New(t)

	r := New(1)
	test.Equal(r.Capacity(), uint32(4))

	r = New(3)
	test.Equal(r.Capacity(), uint32(4))

	r = New(4)
	test.Equal(r.Capacity(), uint32(4))

	r = New(12)
	test.True(r == nil)
}

func TestFreeSize(t *testing.T) {
	test := assert.New(t)

	r := New(8)
	test.Equal(r.freeSize(), uint32(8))

	r.writeCount = uint32(2)
	r.readCount = uint32(math.MaxUint32) - uint32(2)
	test.Equal(r.freeSize(), uint32(3))

	r.writeCount = uint32(3)
	r.readCount = uint32(math.MaxUint32) - uint32(2)
	test.Equal(r.freeSize(), uint32(2))
}

func TestReadWriteOnce(t *testing.T) {
	test := assert.New(t)

	r := New(6)
	test.Equal(r.Capacity(), uint32(8))
	test.True(r.IsEmpty())
	test.False(r.IsFull())
	test.Equal(r.bufferSize(), uint32(0))
	test.Equal(r.freeSize(), uint32(8))

	writedSize := r.writeOnce([]byte("0123"))
	test.Equal(writedSize, uint32(4))
	test.Equal(r.writeOffset(), uint32(4))
	test.False(r.IsEmpty())
	test.False(r.IsFull())
	test.Equal(r.bufferSize(), uint32(4))
	test.Equal(r.freeSize(), uint32(4))

	readBytes := r.readOnce(2)
	test.Equal(len(readBytes), 2)
	test.Equal(string(readBytes), "01")
	test.Equal(r.readOffset(), uint32(2))
	test.False(r.IsEmpty())
	test.False(r.IsFull())
	test.Equal(r.bufferSize(), uint32(2))
	test.Equal(r.freeSize(), uint32(6))

	writedSize = r.writeOnce([]byte("4567890"))
	test.Equal(writedSize, uint32(6))
	test.Equal(r.writeOffset(), uint32(2))
	test.False(r.IsEmpty())
	test.True(r.IsFull())
	test.Equal(r.freeSize(), uint32(0))
	test.Equal(r.bufferSize(), uint32(8))
	test.Equal(string(r.buffer), "89234567")

	writedSize = r.writeOnce([]byte("4567890"))
	test.Equal(writedSize, uint32(0))

	readBytes = r.readOnce(20)
	test.Equal(len(readBytes), 8)
	test.Equal(string(readBytes), "23456789")

	readBytes = r.readOnce(20)
	test.Equal(len(readBytes), 0)
}

func TestReadWrite(t *testing.T) {
	test := assert.New(t)

	r := New(6)
	writedSize := r.Write([]byte("0123"))
	test.Equal(writedSize, uint32(4))

	readBytes := r.Read(2)
	test.Equal(len(readBytes), 2)
	test.Equal(string(readBytes), "01")

	writedSize = r.Write([]byte("456789"))
	test.Equal(writedSize, uint32(6))
	r.SetEOF()

	readBytes = r.Read(20)
	test.Equal(len(readBytes), 8)
	test.Equal(string(readBytes), "23456789")

	readBytes = r.Read(20)
	test.Equal(len(readBytes), 0)

	writedSize = r.Write([]byte("456789"))
	test.Equal(writedSize, uint32(0))
}

func TestExpand1(t *testing.T) {
	test := assert.New(t)

	r := New(3)
	test.Equal(r.Capacity(), uint32(4))
	test.Equal(r.bufferSize(), uint32(0))
	test.Equal(r.writeOffset(), uint32(0))

	r.Expand()
	test.Equal(r.Capacity(), uint32(8))
	test.Equal(r.bufferSize(), uint32(0))
	test.Equal(r.writeOffset(), uint32(0))

	r.Write([]byte("0123"))
	r.Read(1)

	r.Expand()
	test.Equal(r.Capacity(), uint32(16))
	test.Equal(r.bufferSize(), uint32(3))
	test.Equal(r.writeOffset(), uint32(3))
	test.Equal(r.freeSize(), uint32(13))

	readBytes := r.Read(10)
	test.Equal(string(readBytes), "123")
}

func TestExpand2(t *testing.T) {
	test := assert.New(t)

	r := New(8)
	test.Equal(r.Capacity(), uint32(8))
	test.Equal(r.bufferSize(), uint32(0))
	test.Equal(r.writeOffset(), uint32(0))

	r.Write([]byte("01234567"))
	r.Read(1)
	r.Write([]byte("8"))
	test.Equal(string(r.buffer), "81234567")

	r.Expand()
	test.Equal(r.Capacity(), uint32(16))
	test.Equal(r.bufferSize(), uint32(8))
	test.Equal(r.writeOffset(), uint32(8))
	test.Equal(r.freeSize(), uint32(8))

	readBytes := r.Read(8)
	test.Equal(string(readBytes), "12345678")
}
