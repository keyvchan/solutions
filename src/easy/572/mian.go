package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sameTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil {
		left_result := sameTree(root1.Left, root2.Left)
		right_result := sameTree(root1.Right, root2.Right)
		return left_result && right_result
	}
	return false
}

func dfs(root *TreeNode, subRoot *TreeNode, sub_root_list *[]*TreeNode) {
	if root == nil {
		// can't find the subRoot
		return
	}
	if root.Val == subRoot.Val {
		// check result and continue traversing
		*sub_root_list = append(*sub_root_list, root)
	}
	dfs(root.Left, subRoot, sub_root_list)
	dfs(root.Right, subRoot, sub_root_list)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// find all nodes in the tree val is subRoot.Val
	sub_root_list := []*TreeNode{}
	// find the root of the subtree in the tree
	dfs(root, subRoot, &sub_root_list)

	for _, sub_root := range sub_root_list {
		if sameTree(sub_root, subRoot) {
			return true
		}
	}
	return false
}
