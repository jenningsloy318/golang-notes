package bst

import (
	"fmt"
	"testing"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func ToString(node *Node) string {
	if node == nil {
		return ""
	}
	if node.left == nil {
		return fmt.Sprintf("left:0 - value: %v - right: %v\n", node.value, ToString(node.right))
	}
	if node.right == nil {
		return fmt.Sprintf("left:%v - value: %v - right: 0\n", ToString(node.left), node.value)
	}
	return fmt.Sprintf("left:%v - value: %v - right: %v\n", ToString(node.left), node.value, ToString(node.right))

}

type BsTree struct {
	root *Node
	size int
}

func NewBsTree() *BsTree {
	return &BsTree{
		root: nil,
		size: 0,
	}
}

func (b *BsTree) Size() int {
	return b.size
}

func (b *BsTree) IsEmpty() bool {
	return b.size == 0
}

func (b *BsTree) Add(value int) {
	b.root = b.add(b.root, value)

}

func (b *BsTree) add(node *Node, value int) *Node {
	if node == nil {
		b.size++
		return NewNode(value)
	}
	if value < node.value {
		node.left = b.add(node.left, value)
	}

	if value > node.value {

		node.right = b.add(node.right, value)
	}

	return node

}

func (b *BsTree) Contains(value int) bool {
	return contains(b.root, value)
}

func contains(node *Node, value int) bool {

	if node == nil {
		return false
	}

	if node.value == value {
		return true
	} else if value < node.value {
		return contains(node.left, value)
	} else // value > node.value
	{
		return contains(node.right, value)
	}

}

func (b *BsTree) PreOrderTraverse() {
	preOrderTraverse(b.root)

}

func preOrderTraverse(node *Node) {

	if node == nil {
		return
	}

	fmt.Println(node.value)
	preOrderTraverse(node.left)
	preOrderTraverse(node.right)

}

func (b *BsTree) InOrderTraverse() {
	inOrderTraverse(b.root)

}
func inOrderTraverse(node *Node) {
	if node == nil {
		return
	}

	inOrderTraverse(node.left)
	fmt.Println(node.value)
	inOrderTraverse(node.right)

}

func (b *BsTree) PostOrderTraverse() {
	postOrderTraverse(b.root)

}

func postOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	postOrderTraverse(node.left)
	postOrderTraverse(node.right)
	fmt.Println(node.value)

}

// implemented queue
//SliceQueue define a stack
type SliceQueue []*Node

//NewSliceQueue create new SliceQueue
func NewSliceQueue() SliceQueue {
	return make(SliceQueue, 0)
}

//Push  an element to SliceQueue
func (sq *SliceQueue) Enqueue(element *Node) {

	*sq = append(*sq, element)
}

//Dequeue will get the pop the front element from SliceQueue
func (sq *SliceQueue) Dequeue() (element *Node) {
	length := len(*sq)
	if length == 0 {
		fmt.Printf("No element in the queue")
		return nil
	}

	ret := (*sq)[0]

	*sq = (*sq)[1:]

	return ret
}

//GetFront get the front element
func (sq *SliceQueue) GetFront() (element *Node) {
	return (*sq)[0]
}

//IsEmpty
func (sq *SliceQueue) IsEmpty() bool {
	return len(*sq) == 0
}

// implement level traverse using queue

func (b *BsTree) LevelOrderTraverse() {

	if b.root == nil {
		return
	}
	queueBST := NewSliceQueue()

	queueBST.Enqueue(b.root)
	for !queueBST.IsEmpty() {
		curNode := queueBST.Dequeue()
		fmt.Println(curNode.value)

		if curNode.left != nil {
			queueBST.Enqueue(curNode.left)
		}
		if curNode.right != nil {
			queueBST.Enqueue(curNode.right)
		}

	}

}
func TestBstreeBasic(t *testing.T) {

	newBsTree := NewBsTree()

	numbers := []int{33, 21, 77, 99, 36, 15, 76}
	for _, number := range numbers {
		newBsTree.Add(number)
	}

	t.Log("Pre-order traverse: \n")
	newBsTree.PreOrderTraverse()

	t.Log("In-order traverse: \n")
	newBsTree.InOrderTraverse()

	t.Log("Post-order traverse: \n")
	newBsTree.PostOrderTraverse()

	t.Log("Level-order traverse: \n")
	newBsTree.LevelOrderTraverse()
}
