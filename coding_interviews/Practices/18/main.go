package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	var last *ListNode
	tail := head

	for tail != nil {
		if tail.Val == val {
			if last == nil {
				head = tail.Next
			} else {
				last.Next = tail.Next
			}
			break
		} else {
			last = tail
		}

		tail = tail.Next
	}

	return head
}
