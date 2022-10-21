package main

func main() {

}

func containsNearbyDuplicate(nums []int, k int) bool {
	// check for each number

	for i := 0; i < len(nums); i++ {
		// for nums[i], check nums[i] to nums[i+k]
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}

	}
	return false

}
