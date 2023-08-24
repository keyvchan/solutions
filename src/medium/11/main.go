package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

func maxArea(height []int) int {

	max := 0
	for i := 0; i < len(height); i++ {
		for j := len(height) - 1; j > i; j-- {
			area := (j - i) * min(height[i], height[j])
			if area > max {
				max = area
			}
			// trace back to largest height
			max_j := height[j]
			for j > i {
				if max_j > height[j-1] {
					j -= 1
				} else {
					break
				}
			}
		}
		// trace forward to largest height
		max_i := height[i]
		for i < len(height)-1 {
			if max_i > height[i+1] {
				i += 1
			} else {
				break
			}
		}
	}
	return max
}
