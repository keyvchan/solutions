package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	fmt.Println(isSymmetric(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(left *TreeNode, Right *TreeNode) bool {
	if left == nil && Right == nil {
		return true
	}
	if left == nil || Right == nil {
		return false
	}
	if left.Val != Right.Val {
		return false
	}
	return dfs(left.Left, Right.Right) && dfs(left.Right, Right.Left)

}

func isSymmetric(root *TreeNode) bool {
	// dfs
	return dfs(root.Left, root.Right)
}
