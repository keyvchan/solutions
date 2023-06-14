package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	new_head := mergeTwoLists(head1, head2)
	for new_head != nil {
		print(new_head.Val)
		new_head = new_head.Next
	}

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var new_head *ListNode
	var new_tail *ListNode

	for {
		if list1 != nil && list2 != nil {
			// check list element and list2 element
			if list1.Val < list2.Val {
				if new_tail == nil {
					// we initialize new_head and new_tail
					new_head = &ListNode{Val: list1.Val, Next: nil}
					new_tail = new_head
					list1 = list1.Next
				} else {
					// we append new_tail
					new_tail.Next = &ListNode{Val: list1.Val, Next: nil}
					new_tail = new_tail.Next
					list1 = list1.Next
				}
			} else {
				if new_tail == nil {
					// we initialize new_head and new_tail
					new_head = &ListNode{Val: list2.Val, Next: nil}
					new_tail = new_head
					list2 = list2.Next
				} else {
					// we append new_tail
					new_tail.Next = &ListNode{Val: list2.Val, Next: nil}
					new_tail = new_tail.Next
					list2 = list2.Next
				}
			}

		} else {
			if list1 == nil {
				// append list2
				if list2 != nil {
					if new_tail == nil {
						return list2
					}
					new_tail.Next = list2
				}
			} else {
				// append list1
				if list1 != nil {
					if new_tail == nil {
						return list1
					}
					new_tail.Next = list1
				}
			}
			break

		}

	}

	return new_head
}
