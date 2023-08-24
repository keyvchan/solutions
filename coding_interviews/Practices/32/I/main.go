package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		element := queue.Remove(queue.Front()).(*TreeNode)
		if element.Left != nil {
			queue.PushBack(element.Left)
		}
		if element.Right != nil {
			queue.PushBack(element.Right)
		}
		result = append(result, element.Val)
	}
	return result

}
