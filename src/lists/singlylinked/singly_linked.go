package singlylinked

import (
	"fmt"
	"strings"
)

// 参考 gods, 但是 gods 中没有使用虚拟节点
// 启用虚拟节点的好处是在经常头插、尾插的时候方便，但是费一点内存
// https://github.com/emirpasic/gods/blob/master/lists/singlylinkedlist/singlylinkedlist.go

// List ...
type List struct {
	header *element // 虚拟头节点
	tail   *element // 虚拟尾节点
	size   int
}

type element struct {
	value interface{}
	next  *element
}

// New ...
func New(values ...interface{}) *List {
	l := &List{
		header: &element{},
		tail:   &element{},
	}
	l.tail.next = l.header

	if len(values) > 0 {
		l.Append(values...)
	}
	return l
}

// Append ...
func (l *List) Append(values ...interface{}) {
	for _, value := range values {
		newElements := &element{value: value}
		l.tail.next.next = newElements
		l.tail.next = newElements
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
	if l.outOfRange(index) {
		return false
	}

	// 已经检查过了，肯定存在
	var prev *element
	current := l.header.next
	for i := 0; i != index; i++ {
		prev = current
		current = current.next
	}

	if l.header.next == current {
		l.header.next = current.next
	}
	if l.tail.next == current {
		l.tail.next = prev
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
	if l.outOfRange(index) {
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
	current := l.header.next
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
		newElements := &element{value: values[v], next: l.header.next}
		l.header.next = newElements
		l.size++
	}
}

// Reverse ...
func (l *List) Reverse() {
	if l.size <= 1 {
		return
	}

	current := l.header.next
	l.tail.next = current

	var prev, next *element
	for current != nil {
		next = current.next

		current.next = prev
		prev = current

		current = next
	}
	l.header.next = prev
}

//IndexOf returns index of provided element
func (l *List) IndexOf(value interface{}) int {
	if l.size == 0 {
		return -1
	}

	idx := 0
	for item := l.header.next; item != nil; item = item.next {
		if item.value == value {
			return idx
		}
		idx++
	}
	return -1
}

// Set ...
func (l *List) Set(index int, value interface{}) bool {
	if l.outOfRange(index) {
		return true
	}

	foundElement := l.getElementWithIndex(index)
	foundElement.value = value
	return true
}

// Swap ...
func (l *List) Swap(i, j int) bool {
	if l.outOfRange(i) || l.outOfRange(j) {
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

// first ...
func (l *List) firstElement() *element {
	return l.header.next
}

// last ...
func (l *List) lastElement() *element {
	return l.tail.next
}

// String ...
func (l *List) String() string {
	values := []string{}
	for element := l.header.next; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	return strings.Join(values, ",")
}

// getElementWithIndex
func (l *List) getElementWithIndex(index int) *element {
	if l.outOfRange(index) {
		return nil
	}

	foundElement := l.header.next
	for i := 0; i < index; i++ {
		foundElement = foundElement.next
	}

	return foundElement
}

func (l *List) outOfRange(index int) bool {
	return index < 0 || index >= l.size
}
