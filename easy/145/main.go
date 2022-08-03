package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFS(root *TreeNode, value *[]int) {
	if root == nil {
		return
	}

	DFS(root.Left, value)
	DFS(root.Right, value)
	*value = append(*value, root.Val)

}
func postorderTraversal(root *TreeNode) []int {
	value := []int{}
	DFS(root, &value)

	return value
}
