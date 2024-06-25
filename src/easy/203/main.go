package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// loop through the whole list
	result := head

	prev := head
	for head != nil {
		if head.Val == val {
			// remove it
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}
	if result != nil && result.Val == val {
		result = result.Next
	}

	return result
}
