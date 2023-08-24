package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	root = invertTree(root)
	fmt.Println(root)
	fmt.Println(root.Left.Left.Val)
	fmt.Println(root.Left.Right.Val)
	fmt.Println(root.Right.Left.Val)
	fmt.Println(root.Right.Right.Val)

}

func invert(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left

	invert(root.Left)
	invert(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {

	invert(root)

	return root
}
