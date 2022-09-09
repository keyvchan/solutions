package main

import (
	"fmt"
)

func main() {

	root := &TreeNode{Val: 1}
	fmt.Println(lcaDeepestLeaves(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, depth int) (int, *TreeNode) {
	if root == nil {
		return depth, nil
	}

	leftDepth, leftParent := dfs(root.Left, depth+1)
	rightDepth, rightParent := dfs(root.Right, depth+1)
	if leftDepth == rightDepth {
		return leftDepth, root
	}
	if leftDepth > rightDepth {
		return leftDepth, leftParent
	}
	return rightDepth, rightParent
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	// dfs find the deepest node and matain the parent
	_, result := dfs(root, 0)

	return result
}
