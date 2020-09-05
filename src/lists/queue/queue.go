package queue

import (
	"fmt"
	"snippets/src/lists/arraylist"
	"strings"
)

// Queue holds elements in an array-list
type Queue struct {
	list *arraylist.ArrayList
}

// New instantiates a new empty queue
func New() *Queue {
	return &Queue{list: arraylist.New()}
}

// Push adds a value onto the top of the queue
func (s *Queue) Push(value interface{}) {
	s.list.Push(value)
}

// Poll ...
func (s *Queue) Poll() (value interface{}, ok bool) {
	value, ok = s.list.Get(0)
	s.list.Remove(0)
	return
}

// Head ...
func (s *Queue) Head() (value interface{}, ok bool) {
	return s.list.Get(0)
}

// Tail ...
func (s *Queue) Tail() (value interface{}, ok bool) {
	return s.list.Get(s.list.Size() - 1)
}

// IsEmpty ...
func (s *Queue) IsEmpty() bool {
	return s.list.Size() == 0
}

// Size returns number of elements within the queue.
func (s *Queue) Size() int {
	return s.list.Size()
}

// Clear removes all elements from the queue.
func (s *Queue) Clear() {
	s.list.Clear()
}

// Values returns all elements in the queue
func (s *Queue) Values() []interface{} {
	size := s.list.Size()
	elements := make([]interface{}, size, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = s.list.Get(i - 1)
	}
	return elements
}

// String returns a string representation of container
func (s *Queue) String() string {
	str := ""
	values := []string{}
	for _, value := range s.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
