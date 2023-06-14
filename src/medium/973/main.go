package main

import "container/heap"

func main() {

}

type Point struct {
	xy   []int
	dist int
}

type hp []Point

// implement sort.Interface
func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// implement heap.Interface
func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {

	hp := hp{}
	// add all points to heap
	heap.Init(&hp)
	for _, point := range points {
		heap.Push(&hp, Point{
			xy:   point,
			dist: point[0]*point[0] + point[1]*point[1],
		})

	}

	for k > 0 {
		heap.Pop(&hp)
		k--
	}
	res := [][]int{}
	last := heap.Pop(&hp)
	res = append(res, last.(Point).xy)
	for hp.Len() != 0 {

		if hp[0].dist == last.(Point).dist {
			res = append(res, heap.Pop(&hp).(Point).xy)
		} else {
			break
		}

	}
	return res

}
