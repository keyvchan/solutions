package main

func main() {

}

func pivotArray(nums []int, pivot int) []int {
	// create two array
	left := []int{}
	right := []int{}
	equal := []int{}

	for _, num := range nums {
		if num < pivot {
			left = append(left, num)
			continue
		}
		if num > pivot {
			right = append(right, num)
			continue
		}
		equal = append(equal, num)
	}

	return append(append(left, equal...), right...)

}
