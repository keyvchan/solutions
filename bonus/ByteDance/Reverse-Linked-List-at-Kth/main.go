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
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val:  8,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	newHead := solution(head, 3)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}

func solution(head *ListNode, k int) *ListNode {

	stack := list.New()

	// use a stack to store all the node
	counts := 0
	tail := head
	var newHead *ListNode
	var newTail *ListNode
	for tail != nil {
		if counts%k == 0 {
			for stack.Len() > 0 {
				element := stack.Remove(stack.Back()).(*ListNode)

				// if tail is nil, we are having the first node
				if newTail == nil {
					newTail = element
					newHead = newTail
					continue
				}
				newTail.Next = element
				newTail = newTail.Next
			}

		}

		stack.PushBack(tail)
		tail = tail.Next
		counts++
	}

	for stack.Len() > 0 {
		element := stack.Remove(stack.Front()).(*ListNode)

		// if tail is nil, we are having the first node
		if newTail == nil {
			newTail = element
			newHead = newTail
			continue
		}
		newTail.Next = element
		newTail = newTail.Next
	}
	newTail.Next = nil

	return newHead

}
