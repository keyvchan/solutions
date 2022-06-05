package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, parent int, max int, min int, side string) bool {
	if side == "left" {
		// check if root.Val is less than min
		if root.Val >= max {
			return false
		}
		if root.Val >= parent {
			return false
		}
	}
	if side == "right" {
		// check if root.Val is greater than max
		if root.Val <= min {
			return false
		}
		if root.Val <= parent {
			return false
		}
	}
	if root == nil {
		return true
	}
	// set max and min
	if root.Val > max {
		max = root.Val
	}
	if root.Val < min {
		min = root.Val
	}
	return dfs(root.Left, root.Val, max, min, "left") && dfs(root.Right, root.Val, max, min, "right")

}

func isValidBST(root *TreeNode) bool {
	return dfs(root, root.Val, root.Val, root.Val, "root")
}
