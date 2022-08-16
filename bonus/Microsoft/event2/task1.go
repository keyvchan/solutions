package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(Solution([]int{5, 19, 8, 1}))
	fmt.Println(Solution([]int{10, 10}))
	fmt.Println(Solution([]int{3, 0, 5}))
}

type Heap []float32

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(float32))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Solution1(A []int) int {
	// write your code in Go (Go 1.4)
	// create a max heap
	pq := Heap{}
	heap.Init(&pq)

	var total float32
	// push to heap and sumup the total pollution
	for _, v := range A {
		heap.Push(&pq, float32(v))
		total += float32(v)
	}

	var counts int
	var reduction float32
	// pop the max element from the heap
	// and push the next element to the heap
	for reduction < total/2 {
		max_element := heap.Pop(&pq)
		// apply the filer
		heap.Push(&pq, max_element.(float32)/2)
		reduction += max_element.(float32) / 2
		counts++
	}

	return counts
}
