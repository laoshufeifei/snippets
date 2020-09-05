package stack

import (
	"fmt"
	"snippets/src/lists/arraylist"
	"strings"
)

// Stack holds elements in an array-list
type Stack struct {
	list *arraylist.ArrayList
}

// New instantiates a new empty stack
func New() *Stack {
	return &Stack{list: arraylist.New()}
}

// Push adds a value onto the top of the stack
func (s *Stack) Push(value interface{}) {
	s.list.Push(value)
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (s *Stack) Pop() (value interface{}, ok bool) {
	value, ok = s.list.Get(s.list.Size() - 1)
	s.list.Remove(s.list.Size() - 1)
	return
}

// Peek returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (s *Stack) Peek() (value interface{}, ok bool) {
	return s.list.Get(s.list.Size() - 1)
}

// IsEmpty returns true if stack does not contain any elements.
func (s *Stack) IsEmpty() bool {
	return s.list.Size() == 0
}

// Size returns number of elements within the stack.
func (s *Stack) Size() int {
	return s.list.Size()
}

// Clear removes all elements from the stack.
func (s *Stack) Clear() {
	s.list.Clear()
}

// Values returns all elements in the stack (LIFO order).
func (s *Stack) Values() []interface{} {
	size := s.list.Size()
	elements := make([]interface{}, size, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = s.list.Get(i - 1) // in reverse (LIFO)
	}
	return elements
}

// String returns a string representation of container
func (s *Stack) String() string {
	str := ""
	values := []string{}
	for _, value := range s.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
