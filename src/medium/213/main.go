package main

func main() {

}

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	if len(nums) == 1 {
		return nums[0]
	}

	rob_map1 := map[int]int{}

	for _, i := range nums[:len(nums)-1] {
		rob_map1[i] = max(rob_map1[i-1], rob_map1[i-2]+i)
	}

	rob_map2 := map[int]int{}
	for _, i := range nums[1:] {
		rob_map2[i] = max(rob_map2[i-1], rob_map2[i-2]+i)
	}

	return max(rob_map1[len(nums)-2], rob_map2[len(nums)-2])
}
