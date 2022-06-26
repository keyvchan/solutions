package main

import "fmt"

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	fmt.Println(maxScore([]int{2, 2, 2}, 2))
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7))
}

func maxScore(cardPoints []int, k int) int {
	// sub array sum min

	// take len(cardPoints) - k as the sub array length

	min := 2147483647
	last := 0
	for i := len(cardPoints) - k - 1; i < len(cardPoints); i++ {
		if i == -1 {
			min = 0
			break
		}
		if i == len(cardPoints)-k-1 {
			// calculate the sum of the first k elements
			for j := 0; j < len(cardPoints)-k-1; j++ {
				last += cardPoints[j]
			}
		}
		if last += cardPoints[i]; last < min {
			min = last
		}
		last -= cardPoints[i-len(cardPoints)+k+1]
	}

	total_sum := 0
	for _, v := range cardPoints {
		total_sum += v
	}
	return total_sum - min

}
