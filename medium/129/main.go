package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(sumNumbers(root))
	root = &TreeNode{Val: 4,
		Left: &TreeNode{Val: 9,
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{Val: 0},
	}
	fmt.Println(sumNumbers(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, current_num string, sums *[]int) {
	if root == nil {
		return
	}
	current_num += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		num, err := strconv.Atoi(current_num)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(num)
		*sums = append(*sums, num)
		return
	}
	dfs(root.Left, current_num, sums)
	dfs(root.Right, current_num, sums)

}

func sumNumbers(root *TreeNode) int {
	// sum array
	sums := []int{}
	// dfs
	dfs(root, "", &sums)

	sum := 0
	for _, v := range sums {
		sum += v
	}
	return sum
}
