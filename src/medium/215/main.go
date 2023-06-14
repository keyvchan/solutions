package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	old := *pq
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func findKthLargest(nums []int, k int) int {
	pq := PriorityQueue(nums)
	heap.Init(&pq)

	largest := 0
	for i := 0; i < k; i++ {
		largest = heap.Pop(&pq).(int)
	}

	return largest
}
