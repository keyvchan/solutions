package main

import (
	"container/list"
	"sort"
)

func main() {

}

func carFleet(target int, position []int, speed []int) int {
	tuple_list := make([][]int, 0, len(position))
	// tuple list
	for i := 0; i < len(position); i++ {
		tuple_list = append(tuple_list, []int{position[i], speed[i]})
	}

	sort.Slice(tuple_list, func(i, j int) bool {
		return tuple_list[i][0] < tuple_list[j][0]
	})

	stack := list.New()

	for i := len(tuple_list) - 1; i >= 0; i-- {
		time := float64(target-tuple_list[i][0]) / float64(tuple_list[i][1])
		if stack.Len() == 0 {
			// add to stack
			stack.PushBack(time)
		} else {
			back := stack.Back()
			if back.Value.(float64) < time {
				// collision
				stack.PushBack(time)
			}
		}

	}
	return stack.Len()

}
