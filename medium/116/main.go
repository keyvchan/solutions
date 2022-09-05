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

	parent := list.New()
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var temp *Node
		for queue.Len() > 0 {
			element := queue.Remove(queue.Back()).(*Node)
			element.Next = temp
			temp = element

			if element.Right != nil {
				parent.PushFront(element.Right)
			}
			if element.Left != nil {
				parent.PushFront(element.Left)
			}
		}
		queue.PushBackList(parent)
		parent.Init()
	}

	return root
}
