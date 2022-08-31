package main

import "sort"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode 148: https://leetcode.com/problems/sort-list/
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// convert list to array
	var nodes []*ListNode

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}
