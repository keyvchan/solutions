package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// create a array to store all val
	arr := []int{}

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	left, right := 0, len(arr)-1
	pairSum := 0

	for left < right {

		temp := arr[left] + arr[right]
		if temp > pairSum {
			pairSum = temp
		}
		left++
		right--
	}

	return pairSum

}
