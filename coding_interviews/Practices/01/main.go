package main

func main() {

}

func findRepeatNumber(nums []int) int {

	// use a hashmap to find the number that is repeated
	num_map := map[int]bool{}

	for _, num := range nums {
		if num_map[num] {
			return num
		} else {
      num_map[num] = true
    }
	}
	return 0
}
