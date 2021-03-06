package redblacktree

import (
	"fmt"
)

// 自己实现的一个红黑树，为了简化操作，只存储 int 类型

// NodeColor ...
type NodeColor bool

const (
	red   NodeColor = true
	black NodeColor = false
)

type nodeSide int

const (
	none nodeSide = iota
	leftSide
	rightSide
)

type traverseFunc func(n *Node)

// Node ...
type Node struct {
	number int
	color  NodeColor

	parent *Node
	left   *Node
	right  *Node
}

// Tree ...
type Tree struct {
	Root     *Node
	Size     int
	traverse traverseFunc
}

// New ...
func New() *Tree {
	return &Tree{
		traverse: traversePrint,
	}
}

// Get ...
func (t *Tree) Get(number int) *Node {
	node := t.Root
	for node != nil {
		switch {
		case number > node.number:
			node = node.right
		case number < node.number:
			node = node.left
		default:
			return node
		}
	}

	return nil
}

// Put 增加一个元素
// https://juejin.im/post/5e0da754f265da5d2202025a
func (t *Tree) Put(number int) {
	if t.Root == nil {
		t.Root = &Node{number: number, color: black}
		t.Size++
		return
	}

	newNode := &Node{number: number, color: red}

	node := t.Root
	loop := true
	for loop {
		switch {
		case number > node.number:
			if node.right == nil {
				newNode.parent = node
				node.right = newNode
				loop = false
			} else {
				node = node.right
			}
		case number < node.number:
			if node.left == nil {
				newNode.parent = node
				node.left = newNode
				loop = false
			} else {
				node = node.left
			}
		default:
			return
		}
	}
	t.Size++

	t.fixPut(newNode)
}

func (t *Tree) fixPut(node *Node) {
	for {
		parent := node.parent
		if parent == nil {
			// 依然有必要这里判断的原因是有可能上溢走到这里
			node.color = black
			break
		}

		if parent.getColor() == black {
			break
		}

		// else parent.color == red 所以 parent 肯定有父节点，也就是说 uncle 肯定存在
		uncle := node.uncle()
		grandparent := node.parent.parent
		// 这里应该用 uncle 的颜色来判断，因为有可能上溢
		if uncle.getColor() == black {
			// 需要旋转操作
			if parent.whichSide() == leftSide {
				if node.whichSide() == leftSide {
					// LL
					t.rightRotate(grandparent)
					parent.color = black
					grandparent.color = red
				} else {
					// LR
					t.leftRotate(parent)
					t.rightRotate(grandparent)
					node.color = black
					grandparent.color = red
				}
			} else {
				if node.whichSide() == leftSide {
					// RL
					t.rightRotate(parent)
					t.leftRotate(grandparent)
					node.color = black
					grandparent.color = red
				} else {
					// RR
					t.leftRotate(grandparent)
					parent.color = black
					grandparent.color = red
				}
			}

			break
		}

		// uncle is not nil
		if uncle.getColor() == red {
			parent.color = black
			uncle.color = black
			grandparent.color = red
			node = grandparent
		}
	}
}

// Remove a node
func (t *Tree) Remove(number int) {
	node := t.Get(number)
	if node == nil {
		return
	}

	defer func() {
		t.Size--
	}()

	// 转为删除其后继
	if node.right != nil {
		successor := node.successor()
		// swap successor's number and node's number
		node.number = successor.number
		node = successor
	}

	// 接下来 node 的度不会为 2
	// node 的左右子节点要么全没有，要么只有一个

	if node.getColor() == red {
		// 红结点肯定有父节点
		if node.whichSide() == leftSide {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
		return
	}

	// 以下 node 全为 黑色

	child := node.getRedChild()
	if child != nil {
		// 如果含有红孩子，删除红孩子即可
		// 删除的做法也可以用 child 替换 node 的位置
		node.number = child.number

		// red child 的兄弟节点一定为 nil
		// 所以不用判断 red child 是那一边，全都置为 nil
		node.left = nil
		node.right = nil
		return
	}

	// 通过替换来删除 node
	// child 就是用来替换掉 node 的，
	child = node.getChild()

	// 先调整、再删除
	t.fixRemove(node)

	// 注意 child 有可能为 nil
	// 代码还有精简的余地
	if node.parent == nil {
		t.Root = child
		if child != nil {
			child.parent = nil
			child.color = black
		}
	} else {
		if node.whichSide() == leftSide {
			node.parent.left = child
		} else {
			node.parent.right = child
		}

		if child != nil {
			child.parent = node.parent
		}
	}
}

func (t *Tree) fixRemove(n *Node) {
	if n == nil {
		return
	}

	parent := n.parent
	if parent == nil {
		// root
		n.setColor(black)
		return
	}

	if n.getColor() == red {
		return
	}

	sibling := n.sibling()
	if sibling.getColor() == red {
		// 兄弟是红色，父节点降级
		sibling.color = black
		parent.color = red
		if sibling.whichSide() == leftSide {
			t.rightRotate(parent)
		} else {
			t.leftRotate(parent)
		}
		// 注意需要重新计算兄弟节点
		sibling = n.sibling()
	}

	nephew := sibling.getRedChild()
	if nephew != nil {
		// 从黑兄弟那里借红孩子
		pc := parent.color
		if sibling.whichSide() == leftSide {
			if nephew.whichSide() == leftSide {
				// LL
				nephew.color = black
				sibling.color = pc
				parent.color = black
				t.rightRotate(parent)
			} else {
				// LR
				nephew.color = pc
				sibling.color = black
				parent.color = black
				t.leftRotate(sibling)
				t.rightRotate(parent)
			}
		} else {
			if nephew.whichSide() == leftSide {
				// RL
				nephew.color = pc
				sibling.color = black
				parent.color = black
				t.rightRotate(sibling)
				t.leftRotate(parent)
			} else {
				// RR
				nephew.color = black
				sibling.color = pc
				parent.color = black
				t.leftRotate(parent)
			}
		}

		return
	}

	if parent.getColor() == red {
		// parent 是红色 下溢 染色
		sibling.color = red
		parent.color = black
		return
	}

	sibling.color = red
	// parent.color = black
	t.fixRemove(parent)
}

// leftRotate ...
func (t *Tree) leftRotate(n *Node) {
	p := n.parent
	r := n.right

	// 先处理 parent
	if p == nil {
		t.Root = r
	} else {
		if n.whichSide() == leftSide {
			p.left = r
		} else {
			p.right = r
		}
	}

	if r != nil {
		r.parent = p
	}

	n.right = r.left
	if r.left != nil {
		r.left.parent = n
	}

	r.left = n
	n.parent = r
}

func (t *Tree) rightRotate(n *Node) {
	p := n.parent
	l := n.left

	// 先处理 parent
	if p == nil {
		t.Root = l
	} else {
		if n.whichSide() == leftSide {
			p.left = l
		} else {
			p.right = l
		}
	}

	if l != nil {
		l.parent = p
	}

	n.left = l.right
	if l.right != nil {
		l.right.parent = n
	}

	l.right = n
	n.parent = l
}

func (t *Tree) String() string {
	if t.Root == nil {
		return ""
	}

	str := ""
	output(t.Root, "", &str, true)
	return str
}

func output(node *Node, prefix string, str *string, isRight bool) {
	if node.right != nil {
		// 如果觉得这里不好理解，先把这里给改成下面这句看看效果
		// newPrefix := prefix + "    "
		newPrefix := prefix
		if isRight {
			newPrefix += "|   "
		} else {
			newPrefix += "    "
		}
		output(node.right, newPrefix, str, false)
	}

	*str += prefix
	if isRight {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"

	if node.left != nil {
		newPrefix := prefix
		if isRight {
			newPrefix += "    "
		} else {
			newPrefix += "|   "
		}
		output(node.left, newPrefix, str, true)
	}
}

func (t *Tree) isLegalTree() bool {
	root := t.Root
	if t.Size == 0 || root == nil {
		return true
	}

	if root.color == red {
		return false
	}

	bh := t.blackHight()
	if bh < 1 {
		return false
	}

	// 层次遍历
	queue := make(chan *Node, t.Size)
	queue <- root
	for len(queue) > 0 {
		n := <-queue
		if !n.checkColor() {
			return false
		}

		if n.left == nil || n.right == nil {
			if bh != n.blackCountWithinRoot() {
				return false
			}
		}

		if n.left != nil {
			queue <- n.left
		}

		if n.right != nil {
			queue <- n.right
		}
	}

	return true
}

// 一路向左一直到 nil 所经过的黑节点(不含叶子)个数
func (t *Tree) blackHight() (hight int) {
	n := t.Root
	for n != nil {
		if n.getColor() == black {
			hight++
		}
		n = n.left
	}
	return
}

func traversePrint(n *Node) {
	fmt.Println(n)
}

// Preorder ...
func (t *Tree) Preorder() {
	t.Root.preorderImple(t.traverse)
}

func (n *Node) preorderImple(f traverseFunc) {
	if n == nil {
		return
	}

	f(n)
	n.left.preorderImple(f)
	n.right.preorderImple(f)
}

// Inorder ...
func (t *Tree) Inorder() {
	t.Root.inorderImple(t.traverse)
}

func (n *Node) inorderImple(f traverseFunc) {
	if n == nil {
		return
	}

	n.left.inorderImple(f)
	f(n)
	n.right.inorderImple(f)
}

// Postorder ...
func (t *Tree) Postorder() {
	t.Root.postorderImple(t.traverse)
}

func (n *Node) postorderImple(f traverseFunc) {
	if n == nil {
		return
	}

	n.left.postorderImple(f)
	n.right.postorderImple(f)
	f(n)
}

// Levelorder ...
func (t *Tree) Levelorder() {
	if t.Root == nil {
		return
	}

	queue := make(chan *Node, t.Size)
	queue <- t.Root
	// nextLevelSize 可以用来判断是什么是一行遍历完了
	// nextLevelSize := 1

	for len(queue) > 0 {
		n := <-queue
		t.traverse(n)

		if n.left != nil {
			queue <- n.left
		}

		if n.right != nil {
			queue <- n.right
		}

		// nextLevelSize--
		// if nextLevelSize == 0 {
		// 	nextLevelSize = len(queue)
		// }
	}
}

// getColor get node color
func (n *Node) getColor() NodeColor {
	if n == nil {
		return black
	}

	return n.color
}

func (n *Node) setColor(c NodeColor) {
	if n != nil {
		n.color = c
	}
}

func (n *Node) checkColor() bool {
	if n.getColor() == black {
		return true
	}

	// node color is red
	if n.left.getColor() == red || n.right.getColor() == red {
		return false
	}
	return true
}

// 一定要注意 parent 为 nil 的情况
func (n *Node) whichSide() nodeSide {
	if n == nil || n.parent == nil {
		return none
	}

	if n.parent.left == n {
		return leftSide
	}

	return rightSide
}

func (n *Node) uncle() *Node {
	parent := n.parent
	return parent.sibling()
}

func (n *Node) sibling() *Node {
	parent := n.parent
	if parent.left == n {
		return parent.right
	}

	return parent.left
}

func (n *Node) getRedChild() *Node {
	if n == nil {
		return nil
	}

	if n.left == nil && n.right == nil {
		return nil
	}

	// 优先使用判断同侧的
	child1, child2 := n.left, n.right
	if n.whichSide() == rightSide {
		child1, child2 = child2, child1
	}

	if child1.getColor() == red {
		return child1
	}

	if child2.getColor() == red {
		return child2
	}

	return nil
}

func (n *Node) getChild() *Node {
	if n == nil {
		return nil
	}

	if n.left == nil && n.right == nil {
		return nil
	}

	// 优先使用判断同侧的
	child1, child2 := n.left, n.right
	if n.whichSide() == rightSide {
		child1, child2 = child2, child1
	}

	if child1 != nil {
		return child1
	}

	return child2
}

// 前驱结点 中序遍历的前一个
func (n *Node) precursor() *Node {
	if n == nil {
		return nil
	}

	if n.left != nil {
		return n.left.rightmost()
	}

	for n != nil {
		if n.whichSide() == rightSide {
			return n.parent
		}
		n = n.parent
	}

	return nil
}

// 后继结点 中序遍历的后一个
func (n *Node) successor() *Node {
	if n == nil {
		return nil
	}

	if n.right != nil {
		return n.right.leftmost()
	}

	for n != nil {
		if n.whichSide() == leftSide {
			return n.parent
		}
		n = n.parent
	}

	return nil
}

// 最左边
func (n *Node) leftmost() *Node {
	left := n
	for left != nil && left.left != nil {
		left = left.left
	}
	return left
}

// 最右边
func (n *Node) rightmost() *Node {
	right := n
	for right != nil && right.right != nil {
		right = right.right
	}
	return right
}

// 从 node 开始一直到 root 过程中经过的黑色节点个数
func (n *Node) blackCountWithinRoot() (count int) {
	for n != nil {
		if n.color == black {
			count++
		}
		n = n.parent
	}
	return
}

// 彩色打印，有些情况下可能不好使
func (n *Node) String() string {
	if n.getColor() == black {
		return fmt.Sprintf("%d", n.number)
	}

	// return fmt.Sprintf("%v", n.number)
	// return fmt.Sprintf("\033[31m%v\033[0m", n.number)
	return fmt.Sprintf("%v(r)", n.number)
}
