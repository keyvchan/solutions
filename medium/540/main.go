package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 1, 2, 2, 3}))
}

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			// check i-1 and i+1
			if i+2 >= len(nums) {
				return nums[i]
			}
			if i == 0 {
				return nums[i]
			}
			if nums[i] != nums[i-1] && nums[i+1] == nums[i+2] {
				return nums[i]
			} else {
				return nums[i+1]
			}
		}
	}
	return nums[len(nums)-1]
}
