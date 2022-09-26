package main

import (
	"container/list"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {

	result := [][]int{}
	if root == nil {
		return result
	}

	queue := list.New()

	parent_level := list.New()

	queue.PushBack(root)
	for queue.Len() > 0 {
		current_level := []int{}

		// add its children to the queue
		for queue.Len() > 0 {
			element := queue.Remove(queue.Front()).(*TreeNode)
			if element.Left != nil {
				parent_level.PushBack(element.Left)
			}
			if element.Right != nil {
				parent_level.PushBack(element.Right)
			}
			current_level = append(current_level, element.Val)
		}
		result = append(result, current_level)
		queue.PushBackList(parent_level)
		parent_level.Init()
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
