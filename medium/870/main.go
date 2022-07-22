package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}

type node struct {
	value int
	index int
}
type PQ []node

// Len
func (pq PQ) Len() int {
	return len(pq)
}

// Less
func (pq PQ) Less(i, j int) bool {
	return pq[i].value > pq[j].value
}

// Pop
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// Push
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(node))
}

// Swap
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func advantageCount(nums1 []int, nums2 []int) []int {
	nums2_max_heap := PQ{}
	heap.Init(&nums2_max_heap)
	for i := 0; i < len(nums2); i++ {
		heap.Push(&nums2_max_heap, node{nums2[i], i})
	}
	newNums := make([]int, len(nums1))
	sort.Ints(nums1)

	j := 0
	for i := len(nums1) - 1; i >= j; i-- {
		node := heap.Pop(&nums2_max_heap).(node)
		if node.value < nums1[i] {
			// it's ok to use nums1[i] here
			newNums[node.index] = nums1[i]
		} else {
			// it can't be a advantage, so we use the smallest one
			newNums[node.index] = nums1[j]
			j++
			i++
		}
	}
	return newNums
}
