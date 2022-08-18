package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	// we limit the max and min
	if root.Val <= min || root.Val >= max {
		return false
	}

	return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return dfs(root.Left, -math.MaxInt, math.MaxInt)
}
