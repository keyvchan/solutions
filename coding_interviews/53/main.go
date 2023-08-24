package main

func main() {

}

func missingNumber(nums []int) int {

	// standard formula for sum of 1 to n
	n := len(nums)
	expectedSum := (n * (n + 1)) / 2

	for _, num := range nums {
		expectedSum -= num
	}

	return expectedSum
}
