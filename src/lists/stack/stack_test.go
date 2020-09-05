package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	test := assert.New(t)

	s := New()
	test.True(s.IsEmpty())

	s.Push(1)
	test.False(s.IsEmpty())
	test.Equal(s.Size(), 1)

	s.Push(2)
	test.Equal(s.Size(), 2)

	s.Push(3)
	test.Equal(s.Size(), 3)
	test.Equal(s.String(), "1, 2, 3")

	s.Pop()
	test.Equal(s.Size(), 2)
	test.Equal(s.String(), "1, 2")

	s.Pop()
	test.Equal(s.Size(), 1)
	test.Equal(s.String(), "1")

	s.Pop()
	test.Equal(s.Size(), 0)
}
