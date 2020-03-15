package arraylist

// PublicListInterface show the interfaces of arraylist
type PublicListInterface interface {
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
