package main

func main() {

}

func containsNearbyDuplicate(nums []int, k int) bool {
	// create a hashset that size k
	freq := make(map[int]bool, k)

	for i := 0; i <= k; i++ {
		if i > len(nums)-1 {
			return false
		}
		if freq[nums[i]] {
			return true
		}
		freq[nums[i]] = true
	}

	// the first k number is not duplicate

	for i := k + 1; i < len(nums); i++ {
		// remove the first number
		delete(freq, nums[i-k-1])
		// check the last number
		if freq[nums[i]] {
			return true
		}
		// add the last number
		freq[nums[i]] = true

	}
	return false
}
