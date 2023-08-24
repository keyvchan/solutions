package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
	fmt.Println(reconstructQueue([][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}))
	fmt.Println(reconstructQueue([][]int{{9, 0}, {7, 0}, {1, 9}, {3, 0}, {2, 7}, {5, 3}, {6, 0}, {3, 4}, {6, 2}, {5, 2}}))
}

func reconstructQueue(people [][]int) [][]int {

	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	index_array := make([]int, len(people))
	for i := range people {
		index_array[i] = i
	}

	for i, v := range people {
		if i > v[1] {
			// fmt.Println(i, v)
			index_array[i] = v[1]
			temp_index_array := make([]int, len(people))
			copy(temp_index_array, index_array)
			for j := 0; j < i; j++ {
				if temp_index_array[j] >= v[1] {
					index_array[j] += 1
				}
			}
		}
	}
	temp_people := make([][]int, len(people))
	copy(temp_people, people)
	for i, v := range index_array {
		people[v] = temp_people[i]
	}

	return people
}
