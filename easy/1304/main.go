package main

import "fmt"

func main() {
	fmt.Println(sumZero(4))
	fmt.Println(sumZero(5))
}

func sumZero(n int) []int {
	// odd or even
	return_array := make([]int, n)
	if n%2 == 0 {
		for i, j := -n/2, 0; j < n; i++ {
			return_array[j] = i
			return_array[j+1] = -i
			j = j + 2
		}
	} else {
		n = n - 1
		for i, j := -n/2, 0; j < n; i++ {
			return_array[j] = i
			return_array[j+1] = -i
			j = j + 2
		}
	}

	return return_array
}
