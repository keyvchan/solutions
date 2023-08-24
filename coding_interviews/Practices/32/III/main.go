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

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := list.New()
	counts := 0

	queue.PushBack(root)

	for {
		current_level := []int{}
		temp_list := list.New()
		for queue.Len() > 0 {
			element := queue.Remove(queue.Front()).(*TreeNode)
			current_level = append(current_level, element.Val)
			if element.Left != nil {
				temp_list.PushBack(element.Left)
			}
			if element.Right != nil {
				temp_list.PushBack(element.Right)
			}
		}
		if counts%2 == 1 {
			// reverse slice
			for i, j := 0, len(current_level)-1; i < j; i, j = i+1, j-1 {
				current_level[i], current_level[j] = current_level[j], current_level[i]
			}
		}
		result = append(result, current_level)
		queue.PushBackList(temp_list)
		counts++

		if queue.Len() == 0 {
			break
		}

	}

	return result
}
