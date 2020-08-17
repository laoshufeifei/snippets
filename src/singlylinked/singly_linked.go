package singlylinked

import (
	"fmt"
	"strings"
)

// Copy form gods
// https://github.com/emirpasic/gods/blob/master/lists/singlylinkedlist/singlylinkedlist.go

// List ...
type List struct {
	header *element
	tail   *element
	size   int
}

type element struct {
	value interface{}
	next  *element
}

// New ...
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Append(values...)
	}
	return list
}

// Append ...
func (l *List) Append(values ...interface{}) {
	for _, value := range values {
		newElements := &element{value: value}
		if l.size == 0 {
			l.header = newElements
			l.tail = newElements
		} else {
			l.tail.next = newElements
			l.tail = newElements
		}
		l.size++
	}
}

// Get get value
func (l *List) Get(index int) (interface{}, bool) {
	item := l.getElementWithIndex(index)
	if item == nil {
		return nil, false
	}

	return item.value, true
}

// Remove ...
func (l *List) Remove(index int) bool {
	if l.outofRange(index) {
		return false
	}

	// 已经检查过了，肯定存在
	var prev *element
	current := l.header
	for i := 0; i != index; i++ {
		prev = current
		current = current.next
	}

	if current == l.header {
		l.header = l.header.next
	}
	if current == l.tail {
		l.tail = prev
	}

	if prev != nil {
		prev.next = current.next
	}

	current = nil
	l.size--
	return true
}

// Insert ...
func (l *List) Insert(index int, values ...interface{}) bool {
	if l.outofRange(index) {
		if index == l.size {
			l.Append(values...)
			return true
		}
		return false
	}

	if index == 0 {
		l.Prepend(values...)
		return true
	}

	var prev *element
	current := l.header
	for i := 0; i != index; i++ {
		prev = current
		current = current.next
	}

	for _, value := range values {
		newElements := &element{value: value}
		prev.next = newElements
		prev = newElements
		l.size++
	}
	prev.next = current

	return true
}

// Prepend ...
// ["c","d"] -> Prepend(["a","b"]) -> ["a","b","c",d"]
func (l *List) Prepend(values ...interface{}) {
	for v := len(values) - 1; v >= 0; v-- {
		newElements := &element{value: values[v], next: l.header}
		l.header = newElements
		if l.size == 0 {
			l.tail = newElements
		}
		l.size++
	}
}

// Reverse ...
func (l *List) Reverse() {
	if l.size <= 1 {
		return
	}

	element := l.header
	// 第一个(为了减少在循环中做判断，至少有两个元素)
	l.tail = element
	oldNext := element.next
	element.next = nil

	element = oldNext
	for element != nil {
		oldNext = element.next
		element.next = l.header
		l.header = element

		element = oldNext
	}
}

//IndexOf returns index of provided element
func (l *List) IndexOf(value interface{}) int {
	if l.size == 0 {
		return -1
	}

	idx := 0
	for item := l.header; item != nil; item = item.next {
		if item.value == value {
			return idx
		}
		idx++
	}
	return -1
}

// Set ...
func (l *List) Set(index int, value interface{}) bool {
	if l.outofRange(index) {
		return true
	}

	foundElement := l.getElementWithIndex(index)
	foundElement.value = value
	return true
}

// Swap ...
func (l *List) Swap(i, j int) bool {
	if l.outofRange(i) || l.outofRange(j) {
		return false
	}

	if i == j {
		return true
	}

	element1, element2 := l.getElementWithIndex(i), l.getElementWithIndex(j)
	element1.value, element2.value = element2.value, element1.value
	return true
}

// Size returns number of elements within the list.
func (l *List) Size() int {
	return l.size
}

// String ...
func (l *List) String() string {
	values := []string{}
	for element := l.header; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	return strings.Join(values, ",")
}

// getElementWithIndex
func (l *List) getElementWithIndex(index int) *element {
	if l.outofRange(index) {
		return nil
	}

	foundElement := l.header
	for i := 0; i < index; i++ {
		foundElement = foundElement.next
	}

	return foundElement
}

func (l *List) outofRange(index int) bool {
	return index < 0 || index >= l.size
}
