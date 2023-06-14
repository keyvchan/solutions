package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	ancestor := root
	// search the root, until we find p and q
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			// search left
			root = root.Left
			continue
		}
		if p.Val > root.Val && q.Val > root.Val {
			// search right
			root = root.Right
			continue
		}

		ancestor = root
		break
	}

	return ancestor
}
