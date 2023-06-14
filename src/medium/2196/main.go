package main

import "fmt"

func main() {
	fmt.Println(createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	// create a hashmap to index the certain node
	nodeMap := map[int]*TreeNode{}
	hasParent := map[int]bool{}
	for _, description := range descriptions {
		parent := description[0]
		child := description[1]
		is_left := false
		if description[2] == 1 {
			is_left = true
		} else {
			is_left = false
		}
		var parentNode *TreeNode
		if _, ok := nodeMap[parent]; ok {
			parentNode = nodeMap[parent]
		} else {
			parentNode = &TreeNode{Val: parent}
			nodeMap[parent] = parentNode
		}
		var childNode *TreeNode
		if _, ok := nodeMap[child]; ok {
			childNode = nodeMap[child]
		} else {
			childNode = &TreeNode{Val: child}
			nodeMap[child] = childNode
		}
		if is_left {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}

		hasParent[child] = true
	}

	// diff
	for _, node := range nodeMap {
		if !hasParent[node.Val] {
			return node
		}
	}
	return nil
}
