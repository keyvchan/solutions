package main

import (
	"container/list"
)

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5}, 2)
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5}, 2)
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2}, 1)
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{}, 0)
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i, j := 0, 0

	queue := list.New()
	if n == 0 {
		nums1 = nums1[:m]
		return
	}

	if m == 0 {
		for i, val := range nums2 {
			nums1[i] = val
		}
		return
	}

	for {
		if nums1[i] <= nums2[j] {
			queue.PushBack(nums1[i])
			if i != m-1 {
				i++
			} else {
				for j < n {
					queue.PushBack(nums2[j])
					j++
				}
				break
			}
		} else {
			queue.PushBack(nums2[j])
			if j != n-1 {
				j++
			} else {
				for i < m {
					queue.PushBack(nums1[i])
					i++
				}
				break
			}
		}
		if i >= m && j >= n {
			break
		}
	}

	index := 0
	for {
		elem := queue.Front()
		if elem == nil {
			break
		}
		nums1[index] = queue.Remove(elem).(int)
		index++

	}

}
