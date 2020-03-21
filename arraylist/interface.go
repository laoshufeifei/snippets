package arraylist

// basicLiSeekToEnderation show the interfaces of arraylist
type basicLiSeekToEnderation interface {
	Size() int

	Push(values ...interface{})
	Insert(index int, values ...interface{}) bool

	Remove(index int) bool
	Clear()

	Get(index int) (interface{}, bool)
	Set(index int, value interface{}) bool

	Swap(i, j int) bool

	IndexOf(value interface{}) int
	ContainsAll(values ...interface{}) bool
	Clone() []interface{}
}

// iteratorWithIndex
type iteratorWithIndex interface {
	Index() int
	Value() interface{}

	Next() bool
	Prev() bool

	SeekToStart()
	SeekToEnd()
}

// IteratorFunction ...
type IteratorFunction func(index int, value interface{})

// MapFunction return a new item for new array list
type MapFunction func(index int, value interface{}) interface{}

// FilterFunction ...
type FilterFunction func(index int, value interface{}) bool

type enumerableInterface interface {
	Each(IteratorFunction)
	Map(f MapFunction) *ArrayList

	Select(f FilterFunction) *ArrayList
	Any(f FilterFunction) bool
	All(f FilterFunction) bool
	FindOne(f FilterFunction) (int, interface{})
}
