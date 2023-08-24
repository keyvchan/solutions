package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, nil}}}}}}}}}}
	reorderList(head1)
	for head1 != nil {
		fmt.Println(head1.Val)
		head1 = head1.Next
	}
}

func reorderList(head *ListNode) {
	stack := []*ListNode{}
	tail := head

	for tail != nil {
		stack = append(stack, tail)
		tail = tail.Next
	}

	// so, pop from the stack

	tail = head
	for i := len(stack) - 1; i >= 0; i-- {
		last_node := stack[i]
		fmt.Println("last_node", last_node)
		// insert into the list
		last_node.Next = tail.Next
		tail.Next = last_node
		tail = last_node.Next
		fmt.Println("tail:", tail)
		if last_node == tail {
			fmt.Println(last_node, tail)
			tail.Next = nil
			break
		}

	}

}
