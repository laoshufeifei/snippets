package arraylist

func assetEnumerableImplement() {
	var _ enumerableInterface = (*ArrayList)(nil)
}

// Each run the function(f) for each item
func (list *ArrayList) Each(f IteratorFunction) {
	iter := list.Iterator()
	for iter.Next() {
		f(iter.Index(), iter.Value())
	}
}

// Map run function f for each item, than return a new array list
func (list *ArrayList) Map(f MapFunction) *ArrayList {
	newList := New()
	newList.elements = make([]interface{}, list.Size())

	iter := list.Iterator()
	for iter.Next() {
		newValue := f(iter.Index(), iter.Value())
		newList.Set(iter.Index(), newValue)
	}

	return newList
}

// Select ...
func (list *ArrayList) Select(f FilterFunction) *ArrayList {
	newList := New()

	iter := list.Iterator()
	for iter.Next() {
		if f(iter.Index(), iter.Value()) {
			newList.Push(iter.Value())
		}
	}

	return newList
}

// Any ...
func (list *ArrayList) Any(f FilterFunction) bool {
	iter := list.Iterator()
	for iter.Next() {
		if f(iter.Index(), iter.Value()) {
			return true
		}
	}

	return false
}

// All ...
func (list *ArrayList) All(f FilterFunction) bool {
	iter := list.Iterator()
	for iter.Next() {
		if !f(iter.Index(), iter.Value()) {
			return false
		}
	}

	return false
}

// FindOne ...
func (list *ArrayList) FindOne(f FilterFunction) (int, interface{}) {
	iter := list.Iterator()
	for iter.Next() {
		if f(iter.Index(), iter.Value()) {
			return iter.Index(), iter.Value()
		}
	}

	return -1, nil
}
