package main

func main() {
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	current := m + n - 1
	for m > 0 && n > 0 {

		// scan from right to left
		if nums1[m-1] > nums2[n-1] {
			// put nums1[i] to the end of nums1
			nums1[current] = nums1[m-1]
			m--
		} else {
			nums1[current] = nums2[n-1]
			n--
		}
		current--

	}
	if n > 0 {
		// nums2 is left
		for n > 0 {
			nums1[current] = nums2[n-1]
			n--
			current--
		}
	}

}
