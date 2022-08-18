package main

import "fmt"

func main() {
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
	fmt.Println(minSetSize([]int{7, 7, 7, 7, 7, 7}))
}

func minSetSize(arr []int) int {
	// count words frequency
	freq := map[int]int{}
	for _, v := range arr {
		freq[v]++
	}
	// use a array to store word frequency
	array := make([][]int, len(arr)+1)

	original_len := len(arr)

	for k, v := range freq {
		array[v] = append(array[v], k)
	}

	counts := 0
	new_len := original_len
	// start from most frequent word, delete it and check len
	for i := len(array) - 1; i >= 0; i-- {
		if len(array[i]) == 0 {
			continue
		}
		for j := 0; j < len(array[i]); j++ {
			counts++
			new_len -= i
			if new_len*2 <= original_len {
				return counts
			}
		}
	}

	return counts
}
