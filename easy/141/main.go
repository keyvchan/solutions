package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	one := &ListNode{3, nil}
	two := &ListNode{2, nil}
	three := &ListNode{0, nil}
	four := &ListNode{-4, nil}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = two

	fmt.Println(hasCycle(one))
}

func hasCycle(head *ListNode) bool {

	set := make(map[*ListNode]bool)
	for {
		if head == nil {
			return false
		}
		_, ok := set[head]
		if ok {
			return true
		} else {
			set[head] = true
		}

		head = head.Next
	}

}
