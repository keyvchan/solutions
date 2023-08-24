package main

import (
	"container/heap"
	"fmt"
)

// so we will have a heap

func main() {

}

type Heap struct {
	underlying []Pair
}

type Pair struct {
	index1 int
	index2 int
	sum    int
}

func (h Heap) Len() int { return len(h.underlying) }
func (h Heap) Less(i, j int) bool {
	return h.underlying[i].sum > h.underlying[j].sum
}
func (h Heap) Swap(i, j int) {
	h.underlying[i], h.underlying[j] = h.underlying[j], h.underlying[i]
}
func (h *Heap) Push(x interface{}) {
	xPair := x.(Pair)
	(*h).underlying = append((*h).underlying, xPair)
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old.underlying)
	x := old.underlying[n-1]
	(*h).underlying = old.underlying[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	hp := Heap{
		underlying: []Pair{},
	}

	heap.Init(&hp)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if hp.Len() < k {
				heap.Push(&hp, Pair{i, j, nums1[i] + nums2[j]})
			} else {
				fmt.Println(hp.underlying[0].sum, nums1[i]+nums2[j], i, j)
				if nums1[i]+nums2[j] < hp.underlying[0].sum {
					heap.Pop(&hp)
					heap.Push(&hp, Pair{i, j, nums1[i] + nums2[j]})
				} else {
					break
				}
			}
		}

	}

	res := [][]int{}
	for i := 0; i < k; i++ {
		if hp.Len() == 0 {
			break
		}
		pair := heap.Pop(&hp).(Pair)
		res = append(res, []int{nums1[pair.index1], nums2[pair.index2]})
	}

	// reverse res
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}
