package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 3}
	fmt.Println(zigzagLevelOrder(root))

	root = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// bfs
	queue := list.New()

	next_level := list.New()

	queue.PushBack(root)

	count := 1

	result := [][]int{}
	result = append(result, []int{root.Val})

	for queue.Len() > 0 {
		next_array := []int{}

		for queue.Len() > 0 {
			element := queue.Remove(queue.Front()).(*TreeNode)

			if element.Left != nil {
				next_level.PushBack(element.Left)
			}
			if element.Right != nil {
				next_level.PushBack(element.Right)
			}
		}
		if count%2 == 0 {
			// left to right
			for start := next_level.Front(); start != nil; start = start.Next() {
				next_array = append(next_array, start.Value.(*TreeNode).Val)
			}
		} else {
			for start := next_level.Back(); start != nil; start = start.Prev() {
				next_array = append(next_array, start.Value.(*TreeNode).Val)
			}
		}
		count++
		result = append(result, next_array)
		// queue is not empty, swap it with next_level
		queue, next_level = next_level, queue

	}
	return result[:len(result)-1]

}
