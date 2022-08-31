package main

func main() {

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) (int, int) {
	// dfs
	if root == nil {
		// check current depth
		return 0, 0
	}

	if root.Left == nil && root.Right == nil {
		return 1, root.Val
	}

	depth_l, val_l := dfs(root.Left)
	depth_r, val_r := dfs(root.Right)

	// check depth_l and depth_r
	if depth_l > depth_r {
		return depth_l + 1, val_l
	}
	if depth_l < depth_r {
		return depth_r + 1, val_r
	}
	return depth_l + 1, val_l + val_r
}

func deepestLeavesSum(root *TreeNode) int {
	// dfs

	_, sum := dfs(root)
	return sum
}
