/*
Implement array list,
This package is not thread safe.
*/

package arraylist

// Ensure ArrayList has implement the PubicMethods interface;
// If not will has error when build
func assertInterfaceImplementation() {
	var _ PublicListInterface = (*ArrayList)(nil)
}

// ArrayList holds the elements in a slice
type ArrayList struct {
	elements []interface{}
	size     int
}

// New create a new array list
func New(values ...interface{}) *ArrayList {
	l := &ArrayList{}
	if len(values) != 0 {
		l.Push(values...)
	}

	return l
}

// Size return the size of elements
func (l *ArrayList) Size() int {
	return l.size
}

// Push push one or more values to the end
func (l *ArrayList) Push(values ...interface{}) {
	l.expand(len(values))
	for _, value := range values {
		l.elements[l.size] = value
		l.size++
	}
}

// Insert insert item(s) at the index.
// If index == l.Size() this function is equal Push();
// If index > l.Size() return false
func (l *ArrayList) Insert(index int, values ...interface{}) bool {
	if l.outOfRange(index) {
		if index == l.size {
			l.Push(values...)
			return true
		}
		return false
	}

	delta := len(values)
	l.expand(delta)
	copy(l.elements[index+delta:], l.elements[index:l.size])
	copy(l.elements[index:], values)

	l.size += delta
	return true
}

// Remove remove a item with index.
// If index is out of range, return false
func (l *ArrayList) Remove(index int) bool {
	if l.outOfRange(index) {
		return false
	}

	l.elements[index] = nil
	copy(l.elements[index:], l.elements[index+1:l.size])
	l.size--

	return true
}

// Clear clear all elements
func (l *ArrayList) Clear() {
	l.elements = []interface{}{}
	l.size = 0
}

// Get return the item at index.
// If index is out of range return (nil, false)
func (l *ArrayList) Get(index int) (interface{}, bool) {
	if l.outOfRange(index) {
		return nil, false
	}

	return l.elements[index], true
}

// Set modify the item at index
func (l *ArrayList) Set(index int, value interface{}) bool {
	if l.outOfRange(index) {
		if index == l.size {
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
	if l.size == 0 {
		return -1
	}

	for index, item := range l.elements {
		if item == value {
			return index
		}
	}

	return -1
}

// ContainsAll If every items of values is in the list, will return true.
func (l *ArrayList) ContainsAll(values ...interface{}) bool {
	for _, value := range values {
		index := l.IndexOf(value)
		if index == -1 {
			return false
		}
	}
	return true
}

// Clone return a copy of l.elements, not the references
func (l *ArrayList) Clone() []interface{} {
	newElements := make([]interface{}, l.size, l.size)
	copy(newElements, l.elements)
	return newElements
}

///////////////////////////////////////// Private method

func (l *ArrayList) resize(cap int) {
	newElments := make([]interface{}, cap, cap)
	copy(newElments, l.elements)
	l.elements = newElments
}

func (l *ArrayList) expand(delta int) {
	capacity := cap(l.elements)
	if l.Size()+delta >= capacity {
		capacity = (capacity + delta) * 2
		l.resize(capacity)
	}
}

func (l *ArrayList) outOfRange(index int) bool {
	return index < 0 || index >= l.size
}
