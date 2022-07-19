package main

func main() {
	findPeakElement([]int{1, 2, 3, 1})
}

func findPeakElement(nums []int) int {
	left := false
	right := false
	for i := 0; i < len(nums); i++ {
		// check every number
		if i == 0 {
			left = true
		} else {
			left = nums[i] > nums[i-1]
		}
		if i == len(nums)-1 {
			right = true
		} else {
			right = nums[i] > nums[i+1]
		}
		if left && right {
			return i
		}

	}
	return 0
}
