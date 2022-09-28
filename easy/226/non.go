package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	root = invertTree(root)
	fmt.Println(root)
	fmt.Println(root.Left.Left.Val)
	fmt.Println(root.Left.Right.Val)
	fmt.Println(root.Right.Left.Val)
	fmt.Println(root.Right.Right.Val)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	stack := list.New()
	current := root

	for root != nil || stack.Len() > 0 {
		if root != nil {
			stack.PushBack(root)
			root.Left, root.Right = root.Right, root.Left
			root = root.Left
			continue
		}
		root = stack.Remove(stack.Back()).(*TreeNode)
		root = root.Right
	}

	return current
}
