package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {

	// use a array to store the nodes
	nodes := []*ListNode{}

	tail := head

	for tail != nil {
		nodes = append(nodes, tail)
		tail = tail.Next
	}

	middle := len(nodes) / 2

	if middle == 0 {
		// don't have prev node, so just remove the next node
		return nodes[middle].Next
	} else if middle == len(nodes) {
		// don't have next node, so set the prev node's next to nil, then return the head
		nodes[middle-1].Next = nil

		return head
	} else {
		// remove the middle node
		nodes[middle-1].Next = nodes[middle+1]
		return head
	}

}
