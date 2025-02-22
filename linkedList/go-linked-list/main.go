package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Root *Node
}

func NewNode(val int) *Node {
	newNode := &Node{
		Val:  val,
		Next: nil,
	}

	return newNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Root: nil,
	}
}

func (ll *LinkedList) Add(val int) {
	var newNode *Node = NewNode(val)

	if ll.Root == nil {
		ll.Root = newNode
		return
	}

	var curr *Node = ll.Root
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

func (ll *LinkedList) PrintList() {
	if ll.Root == nil {
		fmt.Println("root is nil")
		return
	}

	curr := ll.Root

	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	ll := NewLinkedList()

	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	ll.PrintList()
}