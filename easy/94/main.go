package main

import "fmt"

// 1
//   2
// 3
// inorder: 1 3 2

func main() {
	var root TreeNode
	root.Val = 1
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	inorderTraversal(&root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// found the left most node

	var leftNodes = []TreeNode{}

	for root != nil {
		leftNodes = append(leftNodes, *root)
		root = root.Left
	}

	orders := []int{}
	for i := len(leftNodes) - 1; i >= 0; i-- {

		fmt.Println("root", leftNodes[i].Val)
		orders = append(orders, leftNodes[i].Val)

		if leftNodes[i].Right != nil {
			fmt.Println("right", leftNodes[i].Right.Val)
			orders_nodes := inorderTraversal(leftNodes[i].Right)
			for _, v := range orders_nodes {
				orders = append(orders, v)
			}
		}
	}
	fmt.Println(orders)

	return orders

}
