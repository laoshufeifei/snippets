package binaryheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryHeapPushOne(t *testing.T) {
	test := assert.New(t)

	h := NewWithIntComparator()
	h.Push(3) // 3
	test.Equal(h.Size(), 1)

	h.Push(2) // 2, 3
	test.Equal(h.Size(), 2)

	h.Push(1) // 1, 3, 2
	test.Equal(h.Size(), 3)

	test.Equal(h.String(), "1, 3, 2")
}

func TestBinaryHeapPushMore(t *testing.T) {
	test := assert.New(t)

	h := NewWithIntComparator()
	h.Push(15, 20, 3, 1, 2)
	// swap 1 3
	// swap 0 1
	// swap 1 4
	test.Equal(h.String(), "1, 2, 3, 20, 15")
	test.Equal(h.Size(), 5)
}

func TestBinaryHeapPop(t *testing.T) {
	test := assert.New(t)

	h := NewWithIntComparator()
	h.Push(15, 20, 3, 1, 2)
	test.Equal(h.String(), "1, 2, 3, 20, 15")
	test.Equal(h.Size(), 5)

	v, ok := h.Pop()
	test.True(ok)
	test.Equal(v, 1)
	test.Equal(h.String(), "2, 15, 3, 20")
	test.Equal(h.Size(), 4)

	v, _ = h.Pop()
	test.Equal(v, 2)
	test.Equal(h.Size(), 3)

	v, _ = h.Pop()
	test.Equal(v, 3)
	test.Equal(h.Size(), 2)

	v, _ = h.Pop()
	test.Equal(v, 15)
	test.Equal(h.Size(), 1)

	v, _ = h.Pop()
	test.Equal(v, 20)
	test.Equal(h.Size(), 0)

	// heap is empty
	v, ok = h.Pop()
	test.False(ok)
	test.Equal(v, nil)
	test.Equal(h.Size(), 0)
}

func TestBinaryHeapPeek(t *testing.T) {
	test := assert.New(t)

	h := NewWithIntComparator()
	h.Push(15, 20, 3, 1, 2)
	top, ok := h.Peek()
	test.True(ok)
	test.Equal(top, 1)
}
