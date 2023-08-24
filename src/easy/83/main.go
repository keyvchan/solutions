package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tail := head

	largest := tail.Val
	current_index := tail

	for tail != nil {
		if tail.Val > largest {
			// insert into current_index.next
			current_index.Next = tail
			current_index = current_index.Next
			largest = tail.Val
		}
		tail = tail.Next
	}
	current_index.Next = nil

	return head
}
