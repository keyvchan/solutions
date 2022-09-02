package main

func main() {

}

func arithmeticTriplets(nums []int, diff int) int {
	// use hashmap for O(1) lookup
	nums_map := map[int]bool{}
	// for every num, search doubles
	for _, num := range nums {
		nums_map[num] = true
	}

	counts := 0
	for _, num := range nums {
		// for every num, search triples
		temp := num + diff
		temp2 := temp + diff
		// check temp and temp2
		if _, ok := nums_map[temp]; ok {
			if _, ok := nums_map[temp2]; ok {
				counts++
			}
		}

	}

	return counts

}
