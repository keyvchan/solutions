package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	// bfs
	queue := list.New()
	queue.PushBack(root)
	level := 1
	largest_sum := root.Val
	largest_level := 1
	for queue.Len() != 0 {

		level_sum := 0
		for i := 0; i < queue.Len(); i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			level_sum += node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		if level_sum > largest_sum {
			largest_sum = level_sum
			largest_level = level
		}
		level++
	}
	return level

}
