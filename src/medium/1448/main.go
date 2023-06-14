package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, max int, total *int) {
	if root == nil {
		return
	}

	if root.Val <= max {
		*total += 1
	} else {
		// set max to value
		max = root.Val
	}
	dfs(root.Left, max, total)
	dfs(root.Right, max, total)

}

func goodNodes(root *TreeNode) int {

	total := 0
	dfs(root, root.Val, &total)

	return total

}
