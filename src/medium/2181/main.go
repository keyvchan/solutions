package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	sum := 0

	tail := head
	newHead := &ListNode{}
	newTail := newHead

	for tail != nil {
		if tail.Val == 0 {
			// we check if its the first zero
			if sum == 0 {
				// we are in the first zero, so do nothing
			} else {
				// we are in consecutive zeros, so we add the sum to the new head
				newTail.Val = sum
				if tail.Next != nil {
					newTail.Next = &ListNode{}
					newTail = newTail.Next
				} else {
					newTail.Next = nil
				}
			}
			sum = 0
		}
		sum += tail.Val
		tail = tail.Next
	}
	// set the new tail to nil
	newTail.Next = nil

	return newHead

}
