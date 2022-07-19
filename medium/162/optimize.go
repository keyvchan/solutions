package main

func main() {
	findPeakElement([]int{1, 2, 3, 1})
}

func findPeakElement(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// check next number
		if i == len(nums)-1 {
			return i
		}

		if nums[i] > nums[i+1] {
			return i
		}
	}
	return 0
}
