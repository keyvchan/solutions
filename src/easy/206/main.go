package main

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	new_head := reverseList(head)

	for new_head != nil {
		fmt.Println(new_head.Val)
		new_head = new_head.Next
	}

}

func reverseList(head *ListNode) *ListNode {
	stack := list.New()
	for head != nil {
		stack.PushBack(head.Val)
		head = head.Next
	}
	if stack.Len() == 0 {
		return nil
	}
	new_head := &ListNode{}
	new_tail := new_head
	for stack.Len() != 0 {
		val := stack.Remove(stack.Back())
		new_tail.Val = val.(int)
		if stack.Len() == 0 {
			new_tail.Next = nil
		} else {
			new_tail.Next = &ListNode{}
		}
		new_tail = new_tail.Next
	}

	return new_head
}
