package main

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	head = partition(head, 3)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {

	tail := head
	var less_head *ListNode
	less_tail := less_head
	var greater_head *ListNode
	greater_tail := greater_head

	// loop through the list
	for tail != nil {
		// check if the current node is less than x
		if tail.Val < x {
			// if it is, add it to the new list
			if less_head == nil {
				less_head = tail
				less_tail = less_head
			} else {
				// attach the current node to the end of the new list
				less_tail.Next = tail
				less_tail = less_tail.Next
			}
		} else {
			if greater_head == nil {
				greater_head = tail
				greater_tail = greater_head
			} else {
				greater_tail.Next = tail
				greater_tail = greater_tail.Next
			}
		}
		// shift the tail
		tail = tail.Next

	}
	// combine the two lists
	if less_tail != nil {
		if greater_tail != nil {
			greater_tail.Next = nil
		}
		less_tail.Next = greater_head
	} else {
		// check if greater_tail is nil
		if greater_tail != nil {
			less_head = greater_head
		}
	}
	return less_head
}
