package main

import (
	"container/list"
	"fmt"
)

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	newhead := deleteDuplicates(head)
	for newhead != nil {
		fmt.Println(newhead.Val)
		newhead = newhead.Next
	}

	fmt.Println()

	head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	newhead = deleteDuplicates(head)
	for newhead != nil {
		fmt.Println(newhead.Val)
		newhead = newhead.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	stack := list.New()

	tail := head
	var poped *ListNode
	for tail != nil {

		if poped != nil {
			if poped.Val == tail.Val {
				tail = tail.Next
				continue
			}
		}

		if stack.Len() == 0 {
			stack.PushBack(tail)
			tail = tail.Next
			continue
		}

		// stack.Len > 0
		last := stack.Back().Value.(*ListNode)
		if tail.Val == last.Val {
			poped = last
			stack.Remove(stack.Back())
			tail = tail.Next
		} else {
			stack.PushBack(tail)
			tail = tail.Next
		}
	}

	var newHead *ListNode
	var newtail *ListNode

	for stack.Len() > 0 {
		node := stack.Remove(stack.Front()).(*ListNode)
		if newHead == nil {
			newHead = node
			newtail = node
			continue
		}
		newtail.Next = node
		newtail = newtail.Next
	}
	if newtail != nil {
		newtail.Next = nil
	}

	return newHead
}
