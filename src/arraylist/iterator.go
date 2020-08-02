package arraylist

// Iterator ...
type Iterator struct {
	list  *ArrayList
	index int
}

// Ensure Iterator has implement the PubicMethods interface;
func assertIteratorImplementation() {
	var _ iteratorWithIndex = (*Iterator)(nil)
}

// Iterator ...
func (l *ArrayList) Iterator() Iterator {
	return Iterator{list: l, index: -1}
}

// Index ...
func (i *Iterator) Index() int {
	return i.index
}

// Value ...
func (i *Iterator) Value() interface{} {
	return i.list.elements[i.index]
}

// Next ...
func (i *Iterator) Next() bool {
	if i.index < i.list.Size() {
		i.index++
		return i.index < i.list.Size()
	}

	return false
}

// Prev ...
func (i *Iterator) Prev() bool {
	if i.index >= 0 {
		i.index--
		return i.index >= 0
	}

	return false
}

// SeekToStart ...
func (i *Iterator) SeekToStart() {
	i.index = -1
}

// SeekToEnd ...
func (i *Iterator) SeekToEnd() {
	i.index = i.list.Size()
}
