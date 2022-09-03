package main

func main() {

}

func numOfSubarrays(arr []int, k int, threshold int) int {
	// create a sliding window of size k
	left, right := 0, k-1
	// calcuate the sum from left to right
	sum := 0
	for i := left; i < right; i++ {
		sum += arr[i]
	}

	counts := 0
	for right < len(arr) {
		// check every subarray of size k
		sum += arr[right]

		if float64(sum)/float64(k) >= float64(threshold) {
			counts++
		}

		sum -= arr[left]
		left++
		right++

	}

	return counts
}
