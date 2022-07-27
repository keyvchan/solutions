package main

import "fmt"

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}

func createTargetArray(nums []int, index []int) []int {
	// create a new array to store the result
	result := make([]int, 0, len(nums))
	// create a new index array to store the index of the result

	for i := 0; i < len(index); i++ {
		// check if the index is the same as the length of the result array
		if index[i] == len(result) {
			result = append(result, nums[i])
		} else {
			// insert the number at the index
			result = append(result[:index[i]], append([]int{nums[i]}, result[index[i]:]...)...)
		}
	}

	return result
}
