package main

import (
	"container/list"
	"fmt"
)

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = reverseBetween(head, 2, 4)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	fmt.Println()

	head = &ListNode{5, nil}
	head = reverseBetween(head, 1, 1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println()

	head = &ListNode{3, &ListNode{5, nil}}
	head = reverseBetween(head, 1, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println()
	head = &ListNode{3, &ListNode{5, nil}}
	head = reverseBetween(head, 1, 1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

	fmt.Println()
	head = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	head = reverseBetween(head, 1, 1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func appent_list(stack *list.List, direction string, new_head *ListNode, new_tail *ListNode) (*ListNode, *ListNode) {
	for {
		var element *list.Element
		if direction == "back" {
			element = stack.Back()
		} else {
			element = stack.Front()
		}
		if element == nil {
			break
		}
		node := &ListNode{stack.Remove(element).(int), nil}

		if new_head == nil {
			new_head = node
			new_tail = new_head
		} else {
			new_tail.Next = node
			new_tail = node
		}
	}

	return new_head, new_tail

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// queue based solution

	stack := list.New()
	var new_head *ListNode
	var new_tail *ListNode
	index := 1

	new_head = nil
	new_tail = new_head

	for {
		if head == nil {
			if index > right+1 {
				new_head, new_tail = appent_list(stack, "front", new_head, new_tail)

			} else {
				new_head, new_tail = appent_list(stack, "back", new_head, new_tail)
			}
			break
		} else if index == left {
			new_head, new_tail = appent_list(stack, "front", new_head, new_tail)
		} else if index == right+1 {
			new_head, new_tail = appent_list(stack, "back", new_head, new_tail)

		}
		stack.PushBack(head.Val)
		head = head.Next

		index++
	}
	return new_head

}
