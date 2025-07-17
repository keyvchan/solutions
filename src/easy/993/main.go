package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	parentVal int
	Depth     int
}

func dfs(parentVal int, root *TreeNode, x int, y int, nodeX *Node, nodeY *Node, depth int) {

	if root == nil {
		return
	}

	if nodeX.Depth != 0 && nodeY.Depth != 0 {
		return
	}

	if root.Val == x {
		nodeX.parentVal = parentVal
		nodeX.Depth = depth
	}

	if root.Val == y {
		nodeY.parentVal = parentVal
		nodeY.Depth = depth
	}

	dfs(root.Val, root.Left, x, y, nodeX, nodeY, depth+1)
	dfs(root.Val, root.Right, x, y, nodeX, nodeY, depth+1)
}

func isCousins(root *TreeNode, x int, y int) bool {
	// find depth of each node
	nodeX := Node{0, 0}
	nodeY := Node{0, 0}

	dfs(0, root, x, y, &nodeX, &nodeY, 0)

	return nodeX.parentVal != nodeY.parentVal && nodeX.Depth == nodeY.Depth
}

func main() {
	fmt.Println(isCousins(nil, 0, 0))
	fmt.Println(isCousins(nil, 0, 1))
	// [1,2,3,null,4,null,5]

	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(isCousins(tree, 5, 4))
}
