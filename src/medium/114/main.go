package main

import "fmt"

func main() {
	// test case
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	flatten(root)

	// output
	for root != nil {
		fmt.Println(root.Val, " ")
		root = root.Right
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}

	*nodes = append(*nodes, root)
	dfs(root.Left, nodes)
	dfs(root.Right, nodes)
}

func flatten(root *TreeNode) {
	// use a array to store the nodes
	nodes := []*TreeNode{}
	// pre order traversal
	dfs(root, &nodes)
	// attch the nodes to the root
	for i := 0; i < len(nodes); i++ {
		if root.Left != nil {
			root.Left = nil
		}
		if i == 0 {
			continue
		}
		root.Right = nodes[i]
		root = root.Right
		if i == len(nodes)-1 {
			root = nil
		}
	}

}
