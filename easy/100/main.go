package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func preorder(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}

	return preorder(root1.Left, root2.Left) && preorder(root1.Right, root2.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	return preorder(p, q)
}
