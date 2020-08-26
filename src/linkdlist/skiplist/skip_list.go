package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxLevel    = 32
	probability = 0.25
)

// List ...
type List struct {
	size   int
	level  int
	rand   *rand.Rand
	header *Element // 虚拟节点
}

// Element ...
type Element struct {
	data  int // 先假定 data 都是 int
	nexts []*Element
}

// New skip list
func New() *List {
	header := &Element{
		nexts: make([]*Element, maxLevel),
	}

	return &List{
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		header: header,
	}
}

// Add ...
func (l *List) Add(value int) {
	prevs := make([]*Element, l.level)

	for i := l.level - 1; i >= 0; i-- {
		prev := l.header

		node := l.header.nexts[i]
		for node != nil && node.data < value {
			prev = node
			node = node.nexts[i]
		}

		prevs[i] = prev
		if node != nil && node.data == value {
			// 暂时不允许插入相同的元素
			return
		}
	}

	newLevel := l.randomLevel()
	newElement := &Element{
		data:  value,
		nexts: make([]*Element, newLevel),
	}

	for i := 0; i < newLevel; i++ {
		if i >= l.level {
			l.header.nexts[i] = newElement
		} else {
			newElement.nexts[i] = prevs[i].nexts[i]
			prevs[i].nexts[i] = newElement
		}
	}

	l.size++
	if newLevel > l.level {
		l.level = newLevel
	}

	return
}

// Remove ...
func (l *List) Remove(value int) *Element {
	if l.size == 0 {
		return nil
	}

	prevs := make([]*Element, l.level)

	var current *Element
	for i := l.level - 1; i >= 0; i-- {
		prev := l.header

		node := l.header.nexts[i]
		for node != nil && node.data < value {
			prev = node
			node = node.nexts[i]
		}

		prevs[i] = prev
		if node != nil && node.data == value {
			current = node
		}
	}

	if current == nil {
		return nil
	}

	for i := range current.nexts {
		prevs[i].nexts[i] = current.nexts[i]
	}

	newLevel := l.level
	for newLevel > 0 && l.header.nexts[newLevel-1] == nil {
		newLevel--
	}
	l.level = newLevel

	l.size--
	return current
}

// Get ...
func (l *List) Get(value int) *Element {
	if l.size == 0 {
		return nil
	}

	var node *Element
	for i := l.level - 1; i >= 0; i-- {
		node = l.header.nexts[i]
		for node != nil && node.data < value {
			node = node.nexts[i]
		}

		if node != nil && node.data == value {
			return node
		}
	}

	return nil
}

func (l *List) randomLevel() int {
	level := 1
	f := l.rand.Float32()
	// fmt.Println("f is", f)

	for f < probability && level < maxLevel {
		level++
		f = l.rand.Float32()
		// fmt.Println("f is", f)
	}
	return level
}

func (l *List) String() string {
	s := fmt.Sprintf("The skip list size is %d, level is %d, max level is %d\n", l.size, l.level, maxLevel)
	for i := l.level - 1; i >= 0; i-- {
		s += fmt.Sprintf("level %d: header -> ", i)

		node := l.header.nexts[i]
		for node != nil {
			s += fmt.Sprintf("%v -> ", node.data)
			node = node.nexts[i]
		}

		s += fmt.Sprintf("nil\n")
	}
	return s
}

// Size ...
func (l *List) Size() int {
	return l.size
}

// Level ...
func (l *List) Level() int {
	return l.level
}
