package singlylinked

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedBasic(t *testing.T) {
	test := assert.New(t)

	l := New()
	l.Append("c", "d")
	test.Equal(l.String(), "c,d")
	test.Equal(l.header.value, "c")
	test.Equal(l.tail.value, "d")

	l.Prepend("a", "b")
	test.Equal(l.String(), "a,b,c,d")
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")

	element, ok := l.Get(1)
	test.True(ok)
	test.Equal(element, "b")

	test.Equal(l.Size(), 4)

	test.Equal(l.IndexOf("a"), 0)
	test.Equal(l.IndexOf("b"), 1)
	test.Equal(l.IndexOf("c"), 2)
	test.Equal(l.IndexOf("x"), -1)

	element, ok = l.Get(4)
	test.False(ok)
	test.True(element == nil)
}

func TestSinglyLinkedInsert(t *testing.T) {
	test := assert.New(t)

	l := New("c", "d")
	test.Equal(l.String(), "c,d")

	l.Insert(0, "a", "b")
	test.Equal(l.String(), "a,b,c,d")
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")

	l.Insert(1, "m", "n")
	test.Equal(l.String(), "a,m,n,b,c,d")
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")

	size := l.Size()
	test.Equal(size, 6)

	ok := l.Insert(6, "x", "y")
	test.Equal(l.String(), "a,m,n,b,c,d,x,y")
	test.True(ok)
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "y")

	ok = l.Insert(16, "x", "y", "z")
	test.False(ok)
}

func TestSinglyLinkedInsert2(t *testing.T) {
	test := assert.New(t)

	l := New()
	l.Insert(0, "a", "b", "c", "d")
	test.Equal(l.String(), "a,b,c,d")
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")

	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")
}

func TestSinglyLinkedModify(t *testing.T) {
	test := assert.New(t)

	l := New("a", "b", "c", "d")
	test.Equal(l.String(), "a,b,c,d")

	l.Set(1, "e")
	test.Equal(l.String(), "a,e,c,d")
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")

	l.Swap(0, 1)
	test.Equal(l.String(), "e,a,c,d")
	test.Equal(l.header.value, "e")
	test.Equal(l.tail.value, "d")

	ok := l.Swap(0, 100)
	test.False(ok)
}

func TestSinglyLinkedRemove(t *testing.T) {
	test := assert.New(t)

	l := New("a", "b", "c", "d")
	test.Equal(l.String(), "a,b,c,d")
	test.Equal(l.Size(), 4)
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")

	ok := l.Remove(1)
	test.True(ok)
	test.Equal(l.String(), "a,c,d")
	test.Equal(l.Size(), 3)
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")

	ok = l.Remove(11)
	test.False(ok)
	test.Equal(l.String(), "a,c,d")
	test.Equal(l.Size(), 3)
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")

	ok = l.Remove(0)
	test.True(ok)
	test.Equal(l.String(), "c,d")
	test.Equal(l.Size(), 2)
	test.Equal(l.header.value, "c")
	test.Equal(l.tail.value, "d")

	ok = l.Remove(1)
	test.True(ok)
	test.Equal(l.String(), "c")
	test.Equal(l.Size(), 1)
	test.Equal(l.header.value, "c")
	test.Equal(l.tail.value, "c")
}

func TestSinglyLinkedReverse(t *testing.T) {
	test := assert.New(t)

	l := New("a", "b", "c", "d")
	test.Equal(l.String(), "a,b,c,d")
	test.Equal(l.Size(), 4)
	test.Equal(l.header.value, "a")
	test.Equal(l.tail.value, "d")

	l.Reverse()
	test.Equal(l.String(), "d,c,b,a")
	test.Equal(l.Size(), 4)
	test.Equal(l.header.value, "d")
	test.Equal(l.tail.value, "a")

	l2 := New("a")
	test.Equal(l2.String(), "a")
	l2.Reverse()
	test.Equal(l2.String(), "a")
}
