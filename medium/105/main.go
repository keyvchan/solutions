package main

import (
	"fmt"
)

func main() {
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Right)

	// print the tree BFS
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		fmt.Println()
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtree(indexMap map[int]int, preorder []int, inorder []int, preorder_index *int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	root := &TreeNode{Val: preorder[*preorder_index]}
	*preorder_index++
	root.Left = subtree(indexMap, preorder, inorder, preorder_index, left, indexMap[root.Val]-1)
	root.Right = subtree(indexMap, preorder, inorder, preorder_index, indexMap[root.Val]+1, right)

	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := map[int]int{}
	for i, v := range inorder {
		indexMap[v] = i
	}
	preorder_index := 0
	left := 0
	right := len(inorder) - 1
	return subtree(indexMap, preorder, inorder, &preorder_index, left, right)
}
