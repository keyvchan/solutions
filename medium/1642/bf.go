package main

import "fmt"

func main() {
	fmt.Println(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
	fmt.Println(furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
	fmt.Println(furthestBuilding([]int{14, 3, 19, 3}, 17, 0))
}

func backtracking(height []int, current int, bricks int, ladders int, further int) int {
	if bricks < 0 || ladders < 0 {
		further -= 1
		return further
	}
	if current == len(height)-1 {
		return further
	}

	if height[current] >= height[current+1] {
		return backtracking(height, current+1, bricks, ladders, further+1)
	} else {
		use_bricks := backtracking(height, current+1, bricks-height[current+1]+height[current], ladders, further+1)
		use_ladders := backtracking(height, current+1, bricks, ladders-1, further+1)
		if use_bricks > use_ladders {
			return use_bricks
		}
		return use_ladders
	}

}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	// brute force

	return backtracking(heights, 0, bricks, ladders, 0)

}
