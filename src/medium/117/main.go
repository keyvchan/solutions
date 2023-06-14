package main

import "container/list"

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	// BFS
	queue := list.New()
	queue.PushBack(root)

	var last_node *Node
	for {
		newQueue := list.New()
		for queue.Len() > 0 {
			// BFS from back of queue
			element := queue.Remove(queue.Back()).(*Node)
			element.Next = last_node
			last_node = element
			if element.Left != nil {
				newQueue.PushFront(element.Left)
			}
			if element.Right != nil {
				newQueue.PushFront(element.Right)
			}
		}
		if newQueue.Len() == 0 {
			break
		}
		queue.PushBackList(newQueue)
		last_node = nil
	}

	return root
}
