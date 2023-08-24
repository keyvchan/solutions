package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {

	queue := list.New()
	queue.PushBack(root)
	last_sum := 0
	temp := list.New()

	for queue.Len() > 0 {
		last_sum = 0

		for queue.Len() > 0 {
			element := queue.Remove(queue.Front()).(*TreeNode)
			last_sum += element.Val
			if element.Left != nil {
				temp.PushBack(element.Left)
			}
			if element.Right != nil {
				temp.PushBack(element.Right)
			}
		}
		queue.PushBackList(temp)
		temp.Init()
	}

	return last_sum

}
