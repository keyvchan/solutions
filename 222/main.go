package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func left_depth(root *TreeNode) int {
	left_depth := 0

	for root != nil {
		left_depth += 1
		root = root.Left
	}

	return left_depth
}

func right_depth(root *TreeNode) int {
	right_depth := 0

	for root != nil {
		right_depth += 1
		root = root.Right
	}

	return right_depth
}
func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left_depth := left_depth(root)
	right_depth := right_depth(root)
	if left_depth == right_depth {
		return int(math.Pow(float64(2), float64(left_depth))) - 1
	}

	fmt.Println(left_depth, right_depth)
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func main() {
	tree1 := TreeNode{}
	tree1.Left = &TreeNode{}
	tree1.Right = &TreeNode{}
	tree1.Left.Left = &TreeNode{}
	tree1.Left.Right = &TreeNode{}
	tree1.Right.Left = &TreeNode{}
	tree1.Right.Right = &TreeNode{}
	tree1.Left.Left.Left = &TreeNode{}
	tree1.Left.Left.Right = &TreeNode{}
	tree1.Left.Right.Left = &TreeNode{}
	fmt.Println(countNodes(&tree1))
}
