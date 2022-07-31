package main

func main() {
	// test case

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func post_order(root *TreeNode, parent *TreeNode, left bool) {
	if root == nil {
		return
	}
	post_order(root.Left, root, true)
	post_order(root.Right, root, false)
	if root.Val == 0 {
		// check its children
		if root.Left == nil || root.Right == nil {
			if left {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
	}
}

func pruneTree(root *TreeNode) *TreeNode {
	// DFS on tree, post-order
	post_order(root, root, true)

	return root
}
