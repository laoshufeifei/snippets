package binaryheap

import (
	"fmt"
	"gossips/list/arraylist"
	"gossips/utils"
	"strings"
)

// Heap 默认是实现是按照最小堆来实现的，如果需要最大堆，则修改 comparator 的返回值
type Heap struct {
	list       *arraylist.ArrayList
	comparator utils.Comparator
}

// NewWithComparator instantiates a new emtpy heap tree with custom comparator
func NewWithComparator(comparator utils.Comparator) *Heap {
	return &Heap{list: arraylist.New(), comparator: comparator}
}

// NewWithIntComparator instantiates a new emtpy int heap tree
func NewWithIntComparator() *Heap {
	return &Heap{list: arraylist.New(), comparator: utils.IntComparator}
}

// Push 插入一个或多个元素
func (h *Heap) Push(values ...interface{}) {
	if len(values) == 1 {
		// 先插入到最后一个位置，然后通过节点上浮来调整正确的顺序
		h.list.Push(values[0])
		h.bubbleUp()
	} else {
		// 相当于构造二叉堆，所有非叶子节点倒序下沉
		for _, value := range values {
			h.list.Push(value)
		}

		index := h.list.Size()>>1 - 1
		for i := index; i >= 0; i-- {
			h.bubbleDownFromIndex(i)
		}
	}
}

// Pop 删除堆顶的元素，然后将最后一个元素放到堆顶，并从 0 开始下沉
func (h *Heap) Pop() (value interface{}, ok bool) {
	size := h.list.Size()
	if size == 0 {
		return
	}

	value, _ = h.list.Get(0)
	lastIndex := size - 1
	h.list.Swap(0, lastIndex)
	h.list.Remove(lastIndex)
	h.bubbleDownFromIndex(0)
	return value, true
}

// Peek without remove top
func (h *Heap) Peek() (value interface{}, ok bool) {
	return h.list.Get(0)
}

// String ...
func (h *Heap) String() string {
	values := []string{}
	for _, v := range h.list.Values() {
		values = append(values, fmt.Sprintf("%v", v))
	}
	return strings.Join(values, ", ")
}

// Size ...
func (h *Heap) Size() int {
	return h.list.Size()
}

// Clear ...
func (h *Heap) Clear() {
	h.list.Clear()
}

//////////////////////////////////////
// bubbleUp 节点上浮, 倒序
func (h *Heap) bubbleUp() {
	index := h.list.Size() - 1
	for index > 0 {
		value, _ := h.list.Get(index)
		parentIndex := (index - 1) >> 1
		parentValue, _ := h.list.Get(parentIndex)

		if h.comparator(parentValue, value) <= 0 {
			break
		}

		h.list.Swap(index, parentIndex)
		index = parentIndex
	}
}

// bubleDown 节点下沉, 从 index 到 min(leftIndex, rightIndex), 直到末尾
func (h *Heap) bubbleDownFromIndex(index int) {
	size := h.list.Size()
	leftIndex := 2*index + 1

	for leftIndex < size {
		value, _ := h.list.Get(index)

		// leftIndex := 2*index + 1
		leftValue, _ := h.list.Get(leftIndex)

		rightIndex := 2*index + 2
		rightValue, _ := h.list.Get(rightIndex)

		smallerIndex, smallerValue := leftIndex, leftValue
		if rightIndex < size && h.comparator(leftValue, rightValue) > 0 {
			smallerIndex = rightIndex
			smallerValue = rightValue
		}

		if h.comparator(value, smallerValue) > 0 {
			h.list.Swap(index, smallerIndex)
			// fmt.Println("swap", index, smallerIndex)
		}

		// index = (index - 1) >> 1
		index = smallerIndex
		leftIndex = 2*index + 1
	}
}
