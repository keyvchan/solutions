package main

import "fmt"

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}

func exchange(nums []int) []int {
	// start from start and end, swap if needed
	i, j := 0, len(nums)-1

	for {
		if i >= j {
			break
		}
		// i is odd
		// j is even
		if nums[i]%2 == 0 {
			if nums[j]%2 == 1 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			} else {
				j--
			}
		} else {
			i++
		}

	}
	return nums
}
