package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	counts := 0
	for i, value := range flowerbed {
		if value == 0 {
			// check left and right
			left := 0
			right := 0
			if i == 0 {
				left = 0
			} else {
				left = flowerbed[i-1]

			}
			if i == len(flowerbed)-1 {
				right = 0
			} else {
				right = flowerbed[i+1]

			}
			// check index out of range
			if left == 1 || right == 1 {
				continue
			} else {
				flowerbed[i] = 1
				counts++
			}
		}
		if counts == n {
			return true
		}

	}

	return false
}
