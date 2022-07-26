package main

import "fmt"

func main() {
	// test case 1
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	newHead := insertionSortList(head)
	for newHead != nil {
		fmt.Println("hhh", newHead.Val)
		newHead = newHead.Next
	}

	// test case 2
	head = &ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	}
	newHead = insertionSortList(head)
	for newHead != nil {
		fmt.Println("hhh", newHead.Val)
		newHead = newHead.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	tail := head
	new_head := &ListNode{head.Val, nil}

	// start from second node
	tail = tail.Next
	for tail != nil {
		// loop through the list
		new_tail := new_head
		var last_tail *ListNode
		for new_tail != nil {
			if new_tail.Val <= tail.Val {
				if new_tail.Next == nil {
					// insert at the end
					new_tail.Next = &ListNode{tail.Val, nil}
					break
				} else {
					// current node is smaller than the inserted node, keep looping
					last_tail = new_tail
					new_tail = new_tail.Next
				}

			} else {
				// insert before new_tail
				if last_tail == nil {
					// insert at the head
					new_head = &ListNode{
						Val:  tail.Val,
						Next: new_head,
					}
				} else {
					last_tail.Next = &ListNode{
						Val:  tail.Val,
						Next: new_tail,
					}
				}
				break
			}

		}
		tail = tail.Next
	}
	return new_head

}
