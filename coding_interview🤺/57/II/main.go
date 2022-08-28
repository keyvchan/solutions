package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(findContinuousSequence(9))
}

func findContinuousSequence(target int) [][]int {

	// find the first sequence that sums >= target

	sum := 0
	sequence := list.New()
	for i := 0; i <= target; i++ {
		sum += i
		if sum > target {
			break
		}
		sequence.PushBack(i)
	}

}
