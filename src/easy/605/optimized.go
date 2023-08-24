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

	// we from left to right, check every 0, if it is surrounded by 0, we can plant a flower there

	// we can add 0 at the beginning and end of the array, so we don't need to check the edge case

	for i := 0; i < len(flowerbed); i++ {
		if n <= 0 {
			return true
		}

		// if we can plant a flower, we plant it, and we need to skip the next 0
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			i++
		}
	}

	return n <= 0

}
