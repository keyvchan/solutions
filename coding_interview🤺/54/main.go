package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, result *[]int) {
	// inorder traversal
	if root == nil {
		return
	}
	dfs(root.Left, result)
	*result = append(*result, root.Val)
	dfs(root.Right, result)

}

func kthLargest(root *TreeNode, k int) int {
	result := []int{}
	dfs(root, &result)

	return result[len(result)-k]

}
