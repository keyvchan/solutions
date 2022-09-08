package main

import "container/list"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// stack
	stack := list.New()

	var newHead *ListNode
	var newTail *ListNode

	for head != nil {
		if stack.Len() == 2 {
			for stack.Len() > 0 {
				element := stack.Remove(stack.Back()).(*ListNode)
				if newHead == nil {
					newHead = element
					newTail = newHead
				} else {
					newTail.Next = element
					newTail = newTail.Next
				}
			}
		} else {
			// push
			stack.PushBack(head)
		}
	}
	for stack.Len() > 0 {
		element := stack.Remove(stack.Back()).(*ListNode)
		if newHead == nil {
			newHead = element
			newTail = newHead
		} else {
			newTail.Next = element
			newTail = newTail.Next
		}
	}
	newTail.Next = nil

	return newHead
}
