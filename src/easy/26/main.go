package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

}

func removeDuplicates(nums []int) int {
	largest := nums[0]
	last_index := 0
	for _, value := range nums {
		if value > largest {
			// replace the value at the last index + 1
			nums[last_index+1] = value
			largest = value
			last_index += 1
		}
	}
	fmt.Println(nums)

	return last_index + 1
}
