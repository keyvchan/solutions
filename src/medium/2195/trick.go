package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(minimalKSum([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(minimalKSum([]int{1, 4, 25, 10, 25}, 2))
	fmt.Println(minimalKSum([]int{5, 6}, 6))
}

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)

	last := 0
	sum := int64(0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == last {
			continue
		}
		// check if we can have k numbers
		if nums[i]-last-1 > k {
			// we have k number, range is last+1 to last+1+k
			sum += int64((last + 1 + last + k) * k / 2)
			k = 0
			break
		} else {
			// not enough numbers, range is last+1 to nums[i]
			sum += int64((last + 1 + nums[i] - 1) * (nums[i] - last - 1) / 2)
			k -= nums[i] - last - 1
		}
		last = nums[i]
	}

	for k > 0 {
		sum += int64(last + k)
		k--
	}
	return sum

}
