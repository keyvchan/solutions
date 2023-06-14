package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, current_sum int, target int, result *[][]int, current_nodes []int) {
	if root == nil {
		return
	}

	current_nodes = append(current_nodes, root.Val)
	// copy current_nodes
	new_current_nodes := make([]int, len(current_nodes))
	copy(new_current_nodes, current_nodes)
	dfs(root.Left, current_sum+root.Val, target, result, new_current_nodes)
	dfs(root.Right, current_sum+root.Val, target, result, new_current_nodes)
	if root.Left == nil && root.Right == nil && current_sum+root.Val == target {
		*result = append(*result, current_nodes)
	}

}

func pathSum(root *TreeNode, target int) [][]int {
	result := [][]int{}
	dfs(root, 0, target, &result, []int{})

	return result
}
