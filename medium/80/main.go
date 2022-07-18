package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	largest := nums[0]
	current_index := 0
	duplicates := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == largest {
			// duplicate
			if duplicates < 2 {
				// current_index = nums[i]
				nums[current_index] = nums[i]
				current_index++
			}
			duplicates++
		} else {
			// next number must be larger than largest
			duplicates = 1
			largest = nums[i]
			nums[current_index] = nums[i]
			current_index++
		}

	}
	fmt.Println(nums)
	return current_index

}
