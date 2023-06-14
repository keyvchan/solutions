package main

import "fmt"

func main() {
	tree1 := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Right: &TreeNode{},
			},
		},
		Right: &TreeNode{
			Right: &TreeNode{
				Left: &TreeNode{
					Left: &TreeNode{
						Left: &TreeNode{},
						Right: &TreeNode{
							Left: &TreeNode{},
						},
					},
				},
			},
		},
	}
	fmt.Println(minDepth(tree1))
	fmt.Println(minDepth(nil))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func count(root *TreeNode, depth int) int {
	depth += 1
	if root.Left != nil && root.Right != nil {
		left_depth := count(root.Left, depth)
		right_depth := count(root.Right, depth)
		if left_depth < right_depth {
			return left_depth
		} else {
			return right_depth
		}
	}
	if root.Left != nil {
		return count(root.Left, depth)
	}
	if root.Right != nil {
		return count(root.Right, depth)
	}
	return depth
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return count(root, 0)
}
