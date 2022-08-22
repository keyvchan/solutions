package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))
	fmt.Println(find132pattern([]int{1, 0, 1, -4, -3}))
	fmt.Println(find132pattern([]int{3, 5, 0, 3, 4}))
	fmt.Println(find132pattern([]int{-2, 1, -2}))
	fmt.Println(find132pattern([]int{1, 2, 3, 4, -4, -3, -5, -1}))
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Value struct {
	val      int
	prev_min int
}

func find132pattern(nums []int) bool {
	// use a stack
	prev_min := nums[0]
	stack := list.New()

	for _, num := range nums[1:] {

		for stack.Len() > 0 && num >= stack.Back().Value.(Value).val {
			stack.Remove(stack.Back())
		}
		if stack.Len() > 0 && num > stack.Back().Value.(Value).prev_min {
			return true
		}
		stack.PushBack(Value{num, prev_min})
		prev_min = min(prev_min, num)

	}
	return false

}
