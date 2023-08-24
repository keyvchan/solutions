package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	tail := node.Next
	// store all the following nodes in a slice
	nodes := []*ListNode{}

	// remove the last node
	for tail != nil {
		nodes = append(nodes, tail)
		tail = tail.Next

	}
	// node is the last node,
	tail = node

	for i := 0; i < len(nodes); i++ {
		// reassign the values
		tail.Val = nodes[i].Val
		if tail.Next != nil {
			if tail.Next.Next == nil {
				// remove the last node
				tail.Next = nil
			}

		}
		tail = tail.Next
	}
	fmt.Println(tail)

}
