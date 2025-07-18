package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// create a map
	numsMap := make(map[int]bool, len(nums))

	for _, num := range nums {
		if numsMap[num] {
			return true
		}

		numsMap[num] = true
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
}
