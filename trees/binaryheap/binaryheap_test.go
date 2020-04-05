package binaryheap

import (
	"math/rand"
	"testing"
	"time"

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

// IntComparator compare two int
func _intComparatorForMaxHeap(a, b interface{}) int {
	aValue := a.(int)
	bValue := b.(int)
	switch {
	case aValue == bValue:
		return 0
	case aValue > bValue:
		return -1
	default:
		return 1
	}
}

func TestBinaryHeapMaxHeap(t *testing.T) {
	test := assert.New(t)

	h := NewWithComparator(_intComparatorForMaxHeap)
	h.Push(15, 20, 3, 1, 2)
	test.Equal(h.String(), "20, 15, 3, 1, 2")

	v, _ := h.Pop()
	test.Equal(v, 20)
	test.Equal(h.String(), "15, 2, 3, 1")

	v, _ = h.Pop()
	test.Equal(v, 15)
	test.Equal(h.String(), "3, 2, 1")

	v, _ = h.Pop()
	test.Equal(v, 3)
	test.Equal(h.String(), "2, 1")
}

func TestBinarayHeadRangom(t *testing.T) {
	test := assert.New(t)
	rand.Seed(time.Now().Unix())

	h := NewWithIntComparator()
	for i := 0; i < 10000; i++ {
		h.Push(rand.Intn(10000))
	}

	prev, _ := h.Pop()
	for h.Size() > 0 {
		value, _ := h.Pop()
		test.LessOrEqual(prev.(int), value.(int))
	}
}
