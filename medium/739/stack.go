package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
}

func dailyTemperatures(temperatures []int) []int {
	stack := list.New()
	res := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		fmt.Println(stack.Back())
		if stack.Len() == 0 {
			res[i] = 0
			stack.PushBack([]int{i, temperatures[i]})
			continue
		}
		index := 0
		for stack.Back() != nil {
			if temperatures[i] >= stack.Back().Value.([]int)[1] {
				stack.Remove(stack.Back())
				index = i
			} else {
				index = stack.Back().Value.([]int)[0]
				break
			}
		}
		res[i] = index - i
		stack.PushBack([]int{i, temperatures[i]})
	}
	return res
}
