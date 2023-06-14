package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, array *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, array)
	*array = append(*array, root.Val)
	dfs(root.Right, array)
}

func kthSmallest(root *TreeNode, k int) int {

	array := []int{}

	// inorder traversal
	dfs(root, &array)
	fmt.Println(array)
	return array[k-1]
}
