package main

import "fmt"

func main() {
	// test case with random pointer
	// create some nodes
	node1 := &Node{Val: 7, Random: nil}
	node2 := &Node{Val: 13, Random: nil}
	node3 := &Node{Val: 11, Random: nil}
	node4 := &Node{Val: 10, Random: nil}
	node5 := &Node{Val: 1, Random: nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node1.Random = nil
	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1

	newHead := copyRandomList(node1)
	for newHead != nil {
		fmt.Println(newHead)
		// fmt.Println(newHead.Val, newHead.Random.Val)
		newHead = newHead.Next
	}

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// creat a new list
	var newHead *Node
	var newTail *Node

	// use a hashmap to store the map index to new node
	created_nodes := map[*Node]*Node{}

	tail := head
	// loop through the list
	for tail != nil {
		new_node := &Node{Val: tail.Val, Random: nil}
		created_nodes[tail] = new_node
		if newHead == nil {
			newHead = new_node
			newTail = newHead
		} else {
			// append the new node to the new list
			newTail.Next = new_node
			newTail = newTail.Next
		}

		tail = tail.Next
	}
	// second iteration, place the random pointer
	tail = head
	newTail = newHead
	for tail != nil {
		fmt.Println(tail, newTail)
		// place the random pointer
		if tail.Random != nil {
			newTail.Random = created_nodes[tail.Random]
		}
		tail = tail.Next
		newTail = newTail.Next
	}

	return newHead
}
