/*
Implement array list,
This package is not thread safe.
*/

package arraylist

import "fmt"

// Ensure ArrayList has implement the PubicMethods interface;
// If not will has error when build
func assertBasicImplementation() {
	var _ basicListOperation = (*ArrayList)(nil)
}

// ArrayList holds the elements in a slice
type ArrayList struct {
	elements []interface{}
}

// New create a new array list
func New(items ...interface{}) *ArrayList {
	l := &ArrayList{}

	newLength := len(items)
	if newLength != 0 {
		l.elements = make([]interface{}, newLength)
		for index, item := range items {
			l.elements[index] = item
		}
	}

	return l
}

// NewWithLength make(interface{}, length)
func NewWithLength(length int) *ArrayList {
	l := &ArrayList{}
	l.elements = make([]interface{}, length)
	return l
}

// Size return the size of elements
func (l *ArrayList) Size() int {
	return len(l.elements)
}

// Push push one or more values to the end
func (l *ArrayList) Push(items ...interface{}) {
	l.elements = append(l.elements, items...)
	// size += len(items)
}

// Insert insert item(s) at the index.
// If index == l.Size() this function is equal Push();
// If index > l.Size() return false
func (l *ArrayList) Insert(index int, items ...interface{}) bool {
	size := l.Size()
	if l.outOfRange(index) {
		if index == size {
			l.Push(items...)
			return true
		}
		return false
	}

	delta := len(items)
	l.elements = append(l.elements, items...)
	copy(l.elements[index+delta:], l.elements[index:size])
	copy(l.elements[index:], items)
	// size += delta

	return true
}

// Remove remove a item with index.
// If index is out of range, return false
func (l *ArrayList) Remove(index int) bool {
	if l.outOfRange(index) {
		return false
	}

	l.elements[index] = nil
	size := l.Size()
	copy(l.elements[index:], l.elements[index+1:size])
	// size--
	l.elements = l.elements[:size-1]

	return true
}

// Clear clear all elements
func (l *ArrayList) Clear() {
	l.elements = make([]interface{}, 0)
}

// Get return the item at index.
// If index is out of range return (nil, false)
func (l *ArrayList) Get(index int) (interface{}, bool) {
	fmt.Println(l.Size())
	if l.outOfRange(index) {
		return nil, false
	}

	return l.elements[index], true
}

// Set modify the item at index
func (l *ArrayList) Set(index int, value interface{}) bool {
	if l.outOfRange(index) {
		if index == l.Size() {
			l.Push(value)
			return true
		}
		return false
	}

	l.elements[index] = value
	return true
}

// Swap swap item
func (l *ArrayList) Swap(i, j int) bool {
	if l.outOfRange(i) || l.outOfRange(j) {
		return false
	}

	l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
	return true
}

// IndexOf return the first index that equal value;
// If not found, will return -1
func (l *ArrayList) IndexOf(value interface{}) int {
	if l.Size() == 0 {
		return -1
	}

	for index, item := range l.elements {
		if item == value {
			return index
		}
	}

	return -1
}

// ContainsAll If every item is in the list, will return true.
func (l *ArrayList) ContainsAll(items ...interface{}) bool {
	for _, value := range items {
		index := l.IndexOf(value)
		if index == -1 {
			return false
		}
	}
	return true
}

// Clone return a copy of l.elements, not the references
func (l *ArrayList) Clone() []interface{} {
	size := l.Size()
	newElements := make([]interface{}, size, size)
	copy(newElements, l.elements)
	return newElements
}

///////////////////////////////////////// Private method

func (l *ArrayList) outOfRange(index int) bool {
	return index < 0 || index >= l.Size()
}
