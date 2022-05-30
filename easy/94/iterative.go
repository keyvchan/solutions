package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func inorderTraversal(root *TreeNode) []int {

	result := []int{}
	current_node := root
	stack := list.New()
	for current_node != nil || stack.Len() != 0 {
		for current_node != nil {
			stack.PushBack(current_node)
			current_node = current_node.Left
		}
		current_node = stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, current_node.Val)
		current_node = current_node.Right

	}
	return result
}
