// Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.
//
// Given linked list -- head = [4,5,1,9], which looks like following:
//
//
//
//
//
// Example 1:
//
// Input: head = [4,5,1,9], node = 5
// Output: [4,1,9]
// Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.
// Example 2:
//
// Input: head = [4,5,1,9], node = 1
// Output: [4,5,9]
// Explanation: You are given the third node with value 1, the linked list should become 4 -> 5 -> 9 after calling your function.
//
//
// Note:
//
// The linked list will have at least two elements.
// All of the nodes' values will be unique.
// The given node will not be the tail and it will always be a valid node of the linked list.
// Do not return anything from your function.
//

package DeleteNodeInaLinkedList

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	dummyHead *ListNode
	size      int
}

func (l *LinkedList) addListNode(node *ListNode, val int) *ListNode {

	newNode := &ListNode{
		Val: val,
	}

	if node == nil {
		node = newNode
		l.size++
		return node
	}

	if node.Next != nil {
		node.Next = l.addListNode(node.Next, val)
	} else {
		node.Next = newNode
		l.size++
	}

	return node

}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	node = nil
}

func toString(head *ListNode) {
	var allString []int
	for node := head; node != nil; node = node.Next {
		allString = append(allString, node.Val)

	}
	fmt.Printf("Items in the Linked List: %v\n", allString)
}
func TestConvertBinaryNumberInALinkedListtoInteger(t *testing.T) {
	input := []int{4, 5, 1, 9}
	var dummyHead = &ListNode{}

	l := &LinkedList{
		dummyHead: dummyHead,
		size:      0,
	}

	for _, val := range input {
		l.addListNode(dummyHead, val)

	}
	toString(l.dummyHead.Next)
	t.Logf("Input size: %d", l.size)

	deleteNode(l.dummyHead.Next.Next.Next)
	toString(l.dummyHead.Next)
}
