package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(preorderTraversal(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}

	// non-recursion
	stack := list.New()

	// keep push left child into stack
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			result = append(result, root.Val)
			root = root.Left
		}
		root = stack.Remove(stack.Back()).(*TreeNode)
		root = root.Right
	}

	return result
}
