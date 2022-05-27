package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 15,
				Left: &TreeNode{
					Val: 7,
				},
			},
		},
	}
	fmt.Println(maxDepth(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	left_depth := dfs(root.Left, depth+1)
	right_depth := dfs(root.Right, depth+1)

	if left_depth > right_depth {
		return left_depth
	} else {
		return right_depth
	}
}

func maxDepth(root *TreeNode) int {

	depth := dfs(root, 0)
	return depth
}
