package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := list.New()

	pushed_index, popped_index := 0, 0

	for {
		if stack.Len() == 0 {
			if pushed_index == len(pushed) {
				break
			}
			stack.PushBack(pushed[pushed_index])
			pushed_index++
			continue
		}

		// we check the top of the stack can be popped
		if stack.Back().Value == popped[popped_index] {
			// can be popped
			stack.Remove(stack.Back())
			popped_index++
		} else {
			// can't be popped
			if pushed_index == len(pushed) {
				break
			}
			stack.PushBack(pushed[pushed_index])
			pushed_index++
		}
	}

	return popped_index == len(popped)
}
