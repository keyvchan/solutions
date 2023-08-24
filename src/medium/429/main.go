package main

import "container/list"

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	queue.PushBack(root)

	parent := list.New()

	result := [][]int{}
	for queue.Len() > 0 {
		temp := make([]int, 0, queue.Len())
		for queue.Len() > 0 {
			element := queue.Remove(queue.Front()).(*Node)
			temp = append(temp, element.Val)

			for _, child := range element.Children {
				if child != nil {
					parent.PushBack(child)
				}
			}
		}
		result = append(result, temp)
		queue.PushBackList(parent)
		parent.Init()
	}

	return result
}
