package main

import (
	"fmt"
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(binaryTreePaths(root))
}

func traversal(root *TreeNode, path string, paths *[]string) {

	// check left and right empty
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path+strconv.Itoa(root.Val))
	}
	if root.Left != nil {
		traversal(root.Left, path+strconv.Itoa(root.Val)+"->", paths)
	}
	if root.Right != nil {
		traversal(root.Right, path+strconv.Itoa(root.Val)+"->", paths)
	}

}

func binaryTreePaths(root *TreeNode) []string {

	var paths = []string{}
	traversal(root, "", &paths)
	return paths
}
