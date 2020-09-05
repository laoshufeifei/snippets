package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueuek(t *testing.T) {
	test := assert.New(t)

	q := New()
	test.True(q.IsEmpty())

	q.Push(1)
	test.False(q.IsEmpty())
	test.Equal(q.Size(), 1)
	test.Equal(q.String(), "1")

	head, ok := q.Head()
	test.Equal(head, 1)
	test.True(ok)

	tail, _ := q.Tail()
	test.Equal(tail, 1)
	test.True(ok)

	q.Push(2)
	test.Equal(q.Size(), 2)
	test.Equal(q.String(), "1, 2")

	head, ok = q.Head()
	test.Equal(head, 1)
	test.True(ok)

	tail, _ = q.Tail()
	test.Equal(tail, 2)

	q.Push(3)
	test.Equal(q.Size(), 3)
	test.Equal(q.String(), "1, 2, 3")

	tail, _ = q.Tail()
	test.Equal(tail, 3)

	// 1, 2, 3 ===> 2, 3
	poll, ok := q.Poll()
	test.Equal(poll, 1)
	test.Equal(q.String(), "2, 3")

	head, _ = q.Head()
	test.Equal(head, 2)

	// 2, 3 ===> 3
	poll, ok = q.Poll()
	test.Equal(poll, 2)
	test.Equal(q.String(), "3")

	// 3 ===> []
	poll, ok = q.Poll()
	test.Equal(poll, 3)
	test.Equal(q.Size(), 0)

	poll, ok = q.Poll()
	test.False(ok)
	test.Equal(poll, nil)

	head, ok = q.Head()
	test.False(ok)
	test.Equal(head, nil)
}
