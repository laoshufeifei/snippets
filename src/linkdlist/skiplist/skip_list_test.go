package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipList(t *testing.T) {
	test := assert.New(t)

	l := New()
	test.True(l != nil)
	test.Equal(l.Size(), 0)
	test.Equal(l.Level(), 0)

	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	test.Equal(l.Size(), 4)
	test.True(l.Level() >= 1)

	l.Add(1)
	l.Add(2)
	test.Equal(l.Size(), 6)

	node0, node100 := l.Get(0), l.Get(100)
	test.True(node0 == nil)
	test.True(node100 == nil)

	node1 := l.Get(1)
	test.True(node1 != nil)
	test.True(node1.data == 1)

	node3, node100 := l.Remove(3), l.Remove(100)
	test.True(node3.data == 3)
	test.True(node100 == nil)

	for i := 0; i < 100; i++ {
		l.Remove(i)
	}
	test.Equal(l.Size(), 0)
	test.Equal(l.Level(), 0)
}
