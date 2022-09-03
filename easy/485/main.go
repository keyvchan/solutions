package main

func main() {

}

type window struct {
	left  int
	right int
}

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	// use a sliding window
	w := window{}

	for w.right < len(nums) {
		if nums[w.right] == 1 {
			w.right++
		} else {
			if w.right-w.left > max {
				max = w.right - w.left
			}
			// reset window start point
			w.left = w.right + 1
			w.right++

		}

	}

	// check if the last window is the max
	if w.right-w.left > max {
		max = w.right - w.left
	}

	return max

}
