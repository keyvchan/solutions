package main

import (
	"container/heap"
	"math"
	"sort"
)

func main() {

}

type num struct {
	val int
	abs float64
}

type H []num

// Len
func (h H) Len() int { return len(h) }

// Less
func (h H) Less(i, j int) bool {
	if h[i].abs == h[j].abs {
		return h[i].val < h[j].val
	}
	return h[i].abs < h[j].abs
}

// swap
func (h H) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *H) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(num))
}

func (h *H) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	// use a heap to store every distance of arr[i] to x
	h := H{}
	heap.Init(&h)

	// calcuate distance of every arr[i] to x
	for i := 0; i < len(arr); i++ {
		heap.Push(&h, num{arr[i], math.Abs(float64(arr[i] - x))})
	}

	// pop k elements from heap
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(&h).(num).val)
	}

	sort.Ints(result)
	return result
}
