package redblacktree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	test := assert.New(t)

	tree := New()
	test.True(tree != nil)

	node := tree.Get(1)
	test.True(node == nil)

	tree.Put(3)
	root := tree.Root
	test.Equal(root.number, 3)
	test.Equal(tree.Size, 1)

	node = tree.Get(3)
	test.Equal(node, root)

	// put 5 to right
	tree.Put(5)
	test.Equal(root.number, 3)
	test.Equal(tree.Size, 2)
	test.True(root.getColor() == black)

	test.Equal(root.right.number, 5)
	test.True(root.left == nil)
	test.True(root.right.getColor() == red)

	node = tree.Get(5)
	test.Equal(node.number, 5)
	test.Equal(node, root.right)

	// put 1 to left
	tree.Put(1)
	test.Equal(root.left.number, 1)
	test.True(root.left.getColor() == red)

	node = tree.Get(1)
	test.Equal(node.number, 1)
	test.Equal(node, root.left)

	test.True(tree.Get(10) == nil)

	test.True(tree.isLegalTree())
	test.Equal(tree.blackHight(), 1)
}

func TestSpecialSituation(t *testing.T) {
	test := assert.New(t)

	tree := New()
	test.Equal(tree.Size, 0)
	test.True(tree.Root == nil)
	test.True(tree.isLegalTree())
	tree.Remove(1)

	tree.Put(3)
	test.Equal(tree.Size, 1)
	test.True(tree.Root.number == 3)

	tree.Remove(3)
	test.Equal(tree.Size, 0)
	test.True(tree.Root == nil)
}

func TestNode(t *testing.T) {
	test := assert.New(t)

	tree := New()
	tree.Put(3)
	tree.Put(5)
	tree.Put(1)

	root := tree.Root
	right := root.right
	left := root.left

	test.Equal(root.number, 3)
	test.Equal(right.number, 5)
	test.Equal(left.number, 1)

	test.True(right.whichSide() == rightSide)
	test.False(right.whichSide() == leftSide)

	test.True(left.whichSide() == leftSide)
	test.False(left.whichSide() == rightSide)

	test.True(right.sibling() == left)
	test.True(left.sibling() == right)

	test.True(root.leftmost() == left)
	test.True(root.rightmost() == right)
}

func TestSimplePut(t *testing.T) {
	test := assert.New(t)

	tree := New()
	tree.Put(5)
	tree.Put(1)
	tree.Put(7)

	test.Equal(tree.Get(5).getColor(), black)
	test.Equal(tree.Get(1).getColor(), red)
	test.Equal(tree.Get(7).getColor(), red)

	tree.Put(8)
	tree.Put(6)

	test.Equal(tree.Root.number, 5)
	test.Equal(tree.Get(5).getColor(), black)
	test.Equal(tree.Get(1).getColor(), black)
	test.Equal(tree.Get(7).getColor(), black)
	test.Equal(tree.Get(8).getColor(), red)
	test.Equal(tree.Get(6).getColor(), red)

	test.True(tree.isLegalTree())
	test.Equal(tree.blackHight(), 2)
}

func TestRootRotation(t *testing.T) {
	test := assert.New(t)

	tree := New()
	tree.Put(5)
	tree.Put(3)
	tree.Put(7)
	tree.Put(1)
	tree.Put(4)
	tree.Put(6)
	tree.Put(8)
	// fmt.Println(tree)
	test.Equal(tree.Root.number, 5)

	tree.leftRotate(tree.Root)
	// fmt.Println(tree)
	test.Equal(tree.Root.number, 7)
	test.Equal(tree.Root.left.number, 5)
	test.Equal(tree.Root.right.number, 8)
	test.Equal(tree.Root.left.right.number, 6)

	tree.rightRotate(tree.Root)
	// fmt.Println(tree)
	test.Equal(tree.Root.number, 5)
	test.Equal(tree.Root.left.number, 3)
	test.Equal(tree.Root.right.number, 7)
	test.Equal(tree.Root.right.left.number, 6)
}

func TestRotation(t *testing.T) {
	test := assert.New(t)

	tree := New()
	tree.Put(5)
	tree.Put(3)
	tree.Put(7)
	tree.Put(1)
	tree.Put(4)
	tree.Put(6)
	tree.Put(8)
	// fmt.Println(tree)
	test.Equal(tree.Root.number, 5)

	node7 := tree.Get(7)
	tree.leftRotate(node7)
	// fmt.Println(tree)
	test.Equal(tree.Root.number, 5)
	test.Equal(tree.Root.left.number, 3)
	test.Equal(tree.Root.right.number, 8)
	test.Equal(tree.Root.right.left.number, 7)

	node8 := tree.Get(8)
	tree.rightRotate(node8)
	test.Equal(tree.Root.number, 5)
	test.Equal(tree.Root.left.number, 3)
	test.Equal(tree.Root.right.number, 7)
}

func TestPrevious(t *testing.T) {
	test := assert.New(t)

	tree := New()
	test.True(tree.Root.precursor() == nil)
	tree.Put(5)
	test.True(tree.Root.precursor() == nil)
	tree.Put(3)
	tree.Put(7)
	tree.Put(1)
	tree.Put(4)
	tree.Put(6)
	tree.Put(8)
	// fmt.Println(tree)

	numbers := make([]int, 0)
	node := tree.Root.rightmost()
	for node != nil {
		numbers = append(numbers, node.number)
		node = node.precursor()
	}

	test.Equal(numbers, []int{8, 7, 6, 5, 4, 3, 1})
	test.True(node == nil)

	node1 := tree.Get(1)
	test.True(node1.precursor() == nil)
}

func TestSuccessor(t *testing.T) {
	test := assert.New(t)

	tree := New()
	test.True(tree.Root.successor() == nil)
	tree.Put(5)
	test.True(tree.Root.successor() == nil)
	tree.Put(3)
	tree.Put(7)
	tree.Put(1)
	tree.Put(4)
	tree.Put(6)
	tree.Put(8)
	// fmt.Println(tree)

	numbers := make([]int, 0)
	node := tree.Root.leftmost()
	for node != nil {
		numbers = append(numbers, node.number)
		node = node.successor()
	}

	test.Equal(numbers, []int{1, 3, 4, 5, 6, 7, 8})
	test.True(node == nil)

	node8 := tree.Get(8)
	test.True(node8.successor() == nil)
}

func TestRemove1(t *testing.T) {
	test := assert.New(t)

	tree := New()
	numbers := []int{5, 6, 7, 3, 4, 1, 2}
	for _, i := range numbers {
		tree.Put(i)
	}
	// fmt.Println(tree)
	test.True(tree.isLegalTree())

	// 3
	// 4-2-1
	tree.Remove(6)
	test.True(tree.isLegalTree())
	// fmt.Println(tree)

	// 4-1
	tree.Remove(5)
	tree.Remove(7)
	test.True(tree.isLegalTree())
	// fmt.Println(tree)

	// 4-2-2
	tree.Remove(3)
	tree.Remove(4)
	test.True(tree.isLegalTree())
	// fmt.Println(tree)
}

func TestRemove2(t *testing.T) {
	test := assert.New(t)

	tree := New()
	numbers := []int{1, 28, 58, 11, 88, 91, 56, 62, 90, 94, 6, 3, 61}
	for _, i := range numbers {
		tree.Put(i)
	}
	numbers = []int{3, 58, 61, 56, 90, 94}
	for _, i := range numbers {
		tree.Remove(i)
		test.True(tree.isLegalTree())
	}
	// 全黑
	test.True(tree.isLegalTree())

	tree.Remove(62)
	test.True(tree.isLegalTree())
	// fmt.Println(tree)
}

func TestOrder(t *testing.T) {
	test := assert.New(t)
	tree := New()
	numbers := []int{5, 6, 7, 3, 4, 1, 2}
	for _, i := range numbers {
		tree.Put(i)
	}
	test.True(tree.isLegalTree())

	// preorder
	index := 0
	preArray := make([]int, tree.Size)
	tree.traverse = func(n *Node) {
		preArray[index] = n.number
		index++
	}
	tree.Preorder()
	test.Equal(preArray, []int{6, 4, 2, 1, 3, 5, 7})

	// inorder
	index = 0
	inArray := make([]int, tree.Size)
	tree.traverse = func(n *Node) {
		inArray[index] = n.number
		index++
	}
	tree.Inorder()
	test.Equal(inArray, []int{1, 2, 3, 4, 5, 6, 7})

	// postorder
	index = 0
	postArray := make([]int, tree.Size)
	tree.traverse = func(n *Node) {
		postArray[index] = n.number
		index++
	}
	tree.Postorder()
	test.Equal(postArray, []int{1, 3, 2, 5, 4, 7, 6})

	// levelorder
	index = 0
	levelArray := make([]int, tree.Size)
	tree.traverse = func(n *Node) {
		levelArray[index] = n.number
		index++
	}
	tree.Levelorder()
	test.Equal(levelArray, []int{6, 4, 7, 2, 5, 1, 3})
}

// TODO: convert this test to benchmark test
// lifeTime unit is seconds
func randomTest(lifeTime int64) bool {
	t := New()
	maxNumber := 2000
	startTime := time.Now().Unix()
	for {
		if time.Now().Unix() > lifeTime+startTime {
			break
		}

		t.Put(rand.Intn(maxNumber))
		if !t.isLegalTree() {
			return false
		}

		t.Remove(rand.Intn(maxNumber))
		if !t.isLegalTree() {
			return false
		}

		// time.Sleep(1 * time.Microsecond)
	}

	return true
}

func TestRandomPutAndRemove(t *testing.T) {
	test := assert.New(t)
	test.True(randomTest(2))
}
