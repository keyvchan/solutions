package main

import "fmt"

func main() {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Left.Left)
	fmt.Println(root.Right)
	fmt.Println(root.Right.Left)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func factory(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// create a root node
	root := &TreeNode{Val: nums[len(nums)/2]}

	// left side of the array
	root.Left = factory(nums[:len(nums)/2])
	// right side of the array
	root.Right = factory(nums[len(nums)/2+1:])

	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	root := factory(nums)

	return root

}
