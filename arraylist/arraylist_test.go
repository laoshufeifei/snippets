package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListNew(t *testing.T) {
	test := assert.New(t)

	l1 := New()
	test.NotEqual(l1, nil)
	test.Equal(l1.Size(), 0)
	test.Equal(l1.Clone(), []interface{}{})

	l1.Push(nil)
	v, ok := l1.Get(0)
	test.True(ok)
	test.Equal(v, nil)
	test.Equal(l1.Clone(), []interface{}{nil})

	l2 := New("a", "b", "c")
	test.NotEqual(l2, nil)
	test.Equal(l2.Size(), 3)
	test.Equal(l2.Clone(), []interface{}{"a", "b", "c"})
}

func TestListPush(t *testing.T) {
	test := assert.New(t)

	l := New()
	l.Push("a")
	test.Equal(l.Size(), 1)
	test.Equal(l.Clone(), []interface{}{"a"})

	l.Push("b", "c")
	test.Equal(l.Size(), 3)
	test.Equal(l.Clone(), []interface{}{"a", "b", "c"})

	ok := l.Insert(3, "d", "e")
	test.Equal(l.Size(), 5)
	test.True(ok)
	test.Equal(l.Clone(), []interface{}{"a", "b", "c", "d", "e"})

	ok = l.Insert(10, "d", "e")
	test.Equal(l.Size(), 5)
	test.False(ok)
}

func TestListDelete(t *testing.T) {
	test := assert.New(t)

	l := New("a", "b", "c")
	test.Equal(l.Size(), 3)

	ok := l.Remove(1)
	test.True(ok)
	test.Equal(l.Size(), 2)
	test.Equal(l.Clone(), []interface{}{"a", "c"})

	ok = l.Remove(10)
	test.False(ok)

	l.Clear()
	test.Equal(l.Size(), 0)
	test.Equal(l.Clone(), []interface{}{})
}

func TestListUpdate(t *testing.T) {
	test := assert.New(t)

	l := New("a", "b", "c")
	test.Equal(l.Clone(), []interface{}{"a", "b", "c"})

	ok := l.Swap(0, 1)
	test.True(ok)
	test.Equal(l.Clone(), []interface{}{"b", "a", "c"})

	ok = l.Set(0, "d")
	test.True(ok)
	test.Equal(l.Clone(), []interface{}{"d", "a", "c"})

	ok = l.Set(10, "d")
	test.False(ok)

	v, ok := l.Get(2)
	test.True(ok)
	test.Equal(v, "c")

	_, ok = l.Get(10)
	test.False(ok)
}

func TestListSearch(t *testing.T) {
	test := assert.New(t)
	l := New("a", "b", "c")

	test.Equal(l.IndexOf("a"), 0)
	test.Equal(l.IndexOf("b"), 1)
	test.Equal(l.IndexOf("c"), 2)
	test.Equal(l.IndexOf("d"), -1)

	test.True(l.ContainsAll("a"))
	test.True(l.ContainsAll("b", "c"))
	test.False(l.ContainsAll("b", "c", "d"))
}

func TestListClone(t *testing.T) {
	test := assert.New(t)

	l := New("a", "b", "c")
	items := l.Clone()
	test.Equal(items, []interface{}{"a", "b", "c"})

	// change items
	items[1] = "d"
	test.Equal(l.Clone(), []interface{}{"a", "b", "c"})
	test.NotEqual(l.Clone(), items)
}
