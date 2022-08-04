package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(root *TreeNode, in_max *int) int {
	if root == nil {
		return 0
	}
	left_max := dfs(root.Left, in_max)
	right_max := dfs(root.Right, in_max)
	// check in_max and left_max+right_max+1
	if *in_max < left_max+right_max+1 {
		*in_max = left_max + right_max + 1
	}

	return max(left_max, right_max) + 1

}
func diameterOfBinaryTree(root *TreeNode) int {
	in_max := 0
	max_r := dfs(root, &in_max)

	return max(in_max, max_r)

}
