package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// use a array to store the nodes
	nodes := []*ListNode{}

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	return nodes[len(nodes)-k]

}
