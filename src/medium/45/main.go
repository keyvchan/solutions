package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{2, 0}))
	fmt.Println(jump([]int{0}))
	fmt.Println(jump([]int{1, 1, 1, 1}))
	fmt.Println(jump([]int{1}))
}

func jump(nums []int) int {
	// find minimum steps to reach end in every step
	steps := map[int]int{}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			steps[i] = 0
			continue
		}
		// distance from end
		dist := len(nums) - 1 - i
		// check distance
		if nums[i] >= dist {
			// we can use 1 step to the end
			steps[i] = 1
		} else {
			min := steps[i+1] + 1
			// we gonna find the minimum steps to reach end in every step
			for j := i + 1; j <= i+nums[i]; j++ {
				if steps[j] < min {
					min = steps[j]
				}
			}
			steps[i] = min + 1
		}
	}
	return steps[0]
}
