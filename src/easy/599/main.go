package main

import "fmt"

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
}

func findRestaurant(list1 []string, list2 []string) []string {

	// use a hashmap to store the index if each restaurant
	// key is the restaurant name, value is the index
	indexMap := make(map[string]int)
	for i, restaurant := range list1 {
		indexMap[restaurant] = i
	}

	// create a map to store all the index sum of each restaurant
	index_sums := map[int][]string{}
	for i, restaurant := range list2 {
		if index, ok := indexMap[restaurant]; ok {
			index_sums[i+index] = append(index_sums[i+index], restaurant)
		}
	}

	// find the smallest index sum
	smallest_index_sum := -1
	for index_sum := range index_sums {
		if smallest_index_sum == -1 || index_sum < smallest_index_sum {
			smallest_index_sum = index_sum
		}
	}
	return index_sums[smallest_index_sum]
}
