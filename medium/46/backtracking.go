package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func backtracking(nums, current []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, current)
	}

	for i, num := range nums {
		new_current := make([]int, len(current), len(current)+1)
		copy(new_current, current)
		new_current = append(new_current, num)

		new_nums := make([]int, len(nums))
		copy(new_nums, nums)
		new_nums = append(new_nums[:i], new_nums[i+1:]...)
		backtracking(new_nums, new_current, result)
	}

}

func permute(nums []int) [][]int {
	result := [][]int{}
	backtracking(nums, []int{}, &result)

	return result
}
