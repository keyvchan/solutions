package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sign := 0
	l1_tail := l1
	l2_tail := l2

	new_head := &ListNode{}
	new_tail := new_head

	for l1_tail != nil || l2_tail != nil {
		if l1_tail == nil {
			l1_tail = &ListNode{}
		}
		if l2_tail == nil {
			l2_tail = &ListNode{}
		}
		sum := l1_tail.Val + l2_tail.Val + sign
		if sum >= 10 {
			sum -= 10
			sign = 1
		} else {
			sign = 0
		}
		new_tail.Next = &ListNode{Val: sum, Next: nil}
		new_tail = new_tail.Next

		l1_tail = l1_tail.Next
		l2_tail = l2_tail.Next
	}
	if sign == 1 {
		new_tail.Next = &ListNode{Val: 1, Next: nil}
	}

	return new_head.Next

}
