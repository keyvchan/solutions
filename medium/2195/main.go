package main

func main() {

}

func minimalKSum(nums []int, k int) int64 {
	// use a set to store all number
	nums_set := map[int64]int{}
	for _, num := range nums {
		nums_set[int64(num)]++
	}

	counter := int64(1)
	sum_of_k := int64(0)
	for k > 0 {
		// test counter
		if _, ok := nums_set[counter]; ok {
			// skip this number
			counter++
		} else {
			nums_set[counter]++
			k--
			counter++
			sum_of_k += counter
		}

	}

	return sum_of_k
}
