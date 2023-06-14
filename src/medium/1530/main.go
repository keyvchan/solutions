package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(countPairs(root, 3))

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
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
		},
	}

	fmt.Println(countPairs(root, 3))
	fmt.Println(countPairs(root, 2))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(base *TreeNode, root *TreeNode, parent *TreeNode, adjacency_list map[*TreeNode][]*TreeNode, depth int, distance int) int {
	if root == nil {
		return 0
	}

	// if depth > distance, we can early return
	if depth > distance {
		return 0
	}
	// now we check if node is a leaf
	if root.Left == nil && root.Right == nil && root != base {
		// if it is a leaf and depth is now < distance, we can increment the count
		return 1
	} else {
		// we are in a non-leaf node, so we keep traversing
		count := 0
		for _, node := range adjacency_list[root] {
			if node != parent {
				count += dfs(base, node, root, adjacency_list, depth+1, distance)
			}
		}
		return count
	}
}

func adjacency(root *TreeNode, adjacency_list map[*TreeNode][]*TreeNode, leaf_nodes map[*TreeNode]bool) {
	if root == nil {
		return
	}

	if root.Left != nil {
		adjacency_list[root] = append(adjacency_list[root], root.Left)
		adjacency_list[root.Left] = append(adjacency_list[root.Left], root)
	}
	if root.Right != nil {
		adjacency_list[root] = append(adjacency_list[root], root.Right)
		adjacency_list[root.Right] = append(adjacency_list[root.Right], root)
	}
	if root.Left == nil && root.Right == nil {
		leaf_nodes[root] = true
	}
	adjacency(root.Left, adjacency_list, leaf_nodes)
	adjacency(root.Right, adjacency_list, leaf_nodes)
}

func countPairs(root *TreeNode, distance int) int {
	// form a adjacency list
	adjacency_list := map[*TreeNode][]*TreeNode{}
	leaf_nodes := map[*TreeNode]bool{}
	adjacency(root, adjacency_list, leaf_nodes)

	count := 0
	for node := range leaf_nodes {
		count += dfs(node, node, node, adjacency_list, 0, distance)
	}

	return count / 2
}
