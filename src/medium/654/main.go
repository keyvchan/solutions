package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	new_nums := make([]int, len(nums))
	copy(new_nums, nums)
	sort.Ints(new_nums)

	i := 0
	for i < len(nums) {
		if nums[i] == new_nums[len(new_nums)-1] {
			break
		}
		i++
	}
	root := &TreeNode{
		Val: nums[i],
	}
	root.Left = dfs(nums[:i])
	root.Right = dfs(nums[i+1:])
	return root
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	// create a new

	root := dfs(nums)

	return root
}
