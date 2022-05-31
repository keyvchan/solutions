package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{{root.Val}}

	queue := list.New()
	queue.PushBack(root)

	for {
		next_level := list.New()
		level := []int{}
		// take root from queue, and expand it to its children
		if queue.Len() != 0 {
			element := queue.Remove(queue.Front())
			if element.(*TreeNode).Left != nil {
				level = append(level, element.(*TreeNode).Left.Val)
				next_level.PushBack(element.(*TreeNode).Left)
			}
			if element.(*TreeNode).Right != nil {
				level = append(level, element.(*TreeNode).Right.Val)
				next_level.PushBack(element.(*TreeNode).Right)
			}
		}
		res = append(res, level)
		if next_level.Len() == 0 {
			break
		}
		queue.PushBackList(next_level)
	}
	return res
}
