package main

func main() {

}

func majorityElement(nums []int) int {
	// use a map to count num freq
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	for k, v := range freq {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
