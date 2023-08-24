package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	// list to array
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	// array to tree
	return sortedArrayToBST(nums)
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
