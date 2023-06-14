package main

func main() {

}

func firstMissingPositive(nums []int) int {

	nums_map := map[int]bool{}

	for _, num := range nums {
		if num < 0 {
			continue
		}
		nums_map[num] = true
	}

	// check from 1 to len(nums)
	for i := 1; i <= len(nums); i++ {
		if _, ok := nums_map[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}
