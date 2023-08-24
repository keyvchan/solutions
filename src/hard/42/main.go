package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println(trap([]int{4, 2, 3}))
}

func check_trap(i, j int, height []int) (int, int) {
	forward := 0
	pred := 0
	for {
		if height[i] > height[j] {
			pred += height[j]
			j++
		} else {
			// we found a peak, so we calculate the area
			forward += height[i]*(j-i-1) - pred
			pred = 0
			i = j
			j++
		}
		if j == len(height) {
			j--
			break
		}
	}

	return i, forward
}
func check_back(limit int, height []int) int {
	i, j := len(height)-2, len(height)-1
	backward := 0
	pred := 0

	for {
		if height[i] < height[j] {
			pred += height[i]
			i--
		} else {
			// we found a peak, so we calculate the area
			fmt.Println(i, j)
			backward += height[j]*(j-i-1) - pred
			pred = 0
			j = i
			i--
		}
		if i < limit {
			i--
			break
		}
	}

	return backward
}

func trap(height []int) int {
	// forwards, start from index 0
	if len(height) == 1 {
		return 0
	}

	counts := 0
	// we got a area not covered by the peaks
	i, trap := check_trap(0, 1, height)
	counts += trap

	trap = check_back(i, height)
	counts += trap

	return counts
}
