package main

func main() {

}

func runningSum(nums []int) []int {

	res := make([]int, len(nums))
	last := 0

	for i, v := range nums {
		res[i] = v + last
		last = res[i]
	}

	return res

}
