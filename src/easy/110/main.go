package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	left_height := height(root.Left)
	right_height := height(root.Right)

	if left_height-right_height > 1 || right_height-left_height > 1 {
		panic("not balanced")
	}

	if left_height > right_height {
		return left_height + 1
	} else {
		return right_height + 1
	}
}

func isBalanced(root *TreeNode) bool {
	defer func() bool {
		return recover() == nil
	}()
	height(root)

	return true
}
