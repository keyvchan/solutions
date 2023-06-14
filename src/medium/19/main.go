package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	removeNthFromEnd(head, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tail := head
	stack := []*ListNode{}

	for tail != nil {
		stack = append(stack, tail)
		tail = tail.Next
	}

	// check n
	if n == len(stack) && n != 1 {
		// delete the first one in the list
		return head.Next
	}

	if n == 1 {
		// remove the last
		if len(stack) == 1 {
			head = nil
		} else {
			stack[len(stack)-2].Next = nil
		}
		return head
	}

	prev := stack[len(stack)-n-1]
	next := stack[len(stack)-n+1]

	prev.Next = next

	return head
}
