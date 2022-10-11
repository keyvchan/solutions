package main

import "fmt"

func main() {
	fmt.Println(countQuadruplets([]int{1, 2, 3, 6}))
	fmt.Println(countQuadruplets([]int{3, 3, 6, 4, 5}))
	fmt.Println(countQuadruplets([]int{1, 1, 1, 3, 5}))
}

func backtracking(nums []int, index int, count int, sum int, result *int) {
	if index >= len(nums) {
		return
	}

	if count == 3 {
		// we can now check if sum is equal to current
		if nums[index] == sum {
			*result += 1
		}
		backtracking(nums, index+1, count, sum, result)
	} else {
		// not enouth numbers to check, we can choose current number or not
		// we can choose current number
		backtracking(nums, index+1, count+1, sum+nums[index], result)
		// or not
		backtracking(nums, index+1, count, sum, result)
	}

}

func countQuadruplets(nums []int) int {

	var result int
	backtracking(nums, 0, 0, 0, &result)
	return result

}
