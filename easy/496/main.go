package main

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// implement a brute force solution
	result := []int{}

	// for every element in nums1, find the next greater element in nums2
	for _, num1 := range nums1 {
		for j, num2 := range nums2 {
			if num1 == num2 {
				// find the next greater element
				var k int
				for k = j + 1; k < len(nums2); k++ {
					if nums2[k] > num1 {
						result = append(result, nums2[k])
						break
					}
				}
				if k == len(nums2) {
					result = append(result, -1)
				}
				break
			}
		}

	}
	return result

}
