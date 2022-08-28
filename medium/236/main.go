package main

import "fmt"

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}

	fmt.Println(lowestCommonAncestor(root, root.Left, root.Left.Right.Right))
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root, p, q *TreeNode) (bool, bool, *TreeNode) {
	if root == nil {
		return false, false, nil
	}
	// fmt.Println(root)

	var child_root *TreeNode

	p_l, q_l, child_root := dfs(root.Left, p, q)
	// fmt.Println(root.Val, "<--", "left", p_l, "right", q_l)
	if p_l && q_l {
		return true, true, child_root
	}
	p_r, q_r, child_root := dfs(root.Right, p, q)
	// fmt.Println(root.Val, "-->", "left", p_r, "right", q_r)
	if p_r && q_r {
		return true, true, child_root
	}

	if (p_l || p_r) && (q_l || q_r) {
		return true, true, root
	}

	if p_l || p_r {
		// either left_l or left_r is true, we check q
		if root == q {
			return true, true, root
		} else {
			return true, false, nil
		}
	}

	if q_l || q_r {
		// either left_l or left_r is true, we check q
		if root == p {
			return true, true, root
		} else {
			return false, true, nil
		}
	}

	if root == p {
		return true, false, nil
	}
	if root == q {
		return false, true, nil
	}

	return false, false, nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, _, result := dfs(root, p, q)

	return result
}
