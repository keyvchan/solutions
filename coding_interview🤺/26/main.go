package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: -4,
			},
			Right: &TreeNode{
				Val: -3,
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -4,
		},
	}

	fmt.Println(isSubStructure(root, root2))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func check(A *TreeNode, B *TreeNode) bool {
	// check same tree
	if A == nil || B == nil {
		return true
	}

	if A.Val != B.Val {
		return false
	}
	result1 := check(A.Left, B.Left)
	fmt.Println(A.Left, B.Left, result1)
	result2 := check(A.Right, B.Right)
	fmt.Println(A.Right, B.Right, result2)
	fmt.Println("Final",result1 && result2)
	return result1 && result2
}

func dfs(A *TreeNode, B *TreeNode) bool {
	if A == nil {
		return false
	}
	if A.Val == B.Val {
		// check
		if check(A, B) {
			return true
		}
	}

	return dfs(A.Left, B) || dfs(A.Right, B)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return dfs(A, B)

}
