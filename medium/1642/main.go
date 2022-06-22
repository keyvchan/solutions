package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
	fmt.Println(furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
	fmt.Println(furthestBuilding([]int{14, 3, 19, 3}, 17, 0))
	fmt.Println(furthestBuilding([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 100, 0))
}

type Myheap []int

// hp as a heap
func (h Myheap) Len() int           { return len(h) }
func (h Myheap) Less(i, j int) bool { return h[i] > h[j] }
func (h Myheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Myheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *Myheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

type Gap struct {
	start int
	end   int
	gap   int
}

func furthestBuilding(heights []int, bricks int, ladders int) int {

	// get all the gaps we need to fill
	gaps := []Gap{}

	for i := 0; i < len(heights); i++ {
		if i == 0 {
			continue
		}
		if heights[i] > heights[i-1] {
			gaps = append(gaps, Gap{i, i - 1, heights[i] - heights[i-1]})
		}
	}
	if len(gaps) == 0 {
		return len(heights) - 1
	}

	// sum from start to the end, the ladders we have can be used as a replacement for max heights.
	// use a heap to store ladders largest value

	hp := Myheap{}
	heap.Init(&hp)
	for _, gap := range gaps {
		if gap.gap <= bricks {
			// it means we can fill this gap with bricks, add it to the heap
			heap.Push(&hp, gap.gap)
			bricks -= gap.gap
		} else {
			if ladders == 0 {
				// we reach the end
				return gap.end
			}
			if hp.Len() == 0 {
				// we can't fill this gap with bricks, but we have ladders, so we can fill it with ladders
				ladders--
			} else {
				// we don't have enough bricks to fill this gap, we need to use ladders, but we get the privious largest gap
				if hp[0] > gap.gap {
					// fill the gap with the ladder, and we now get a butchs of bricks available
					bricks = bricks + hp[0] - gap.gap
					// pop the largest ladder
					heap.Pop(&hp)
					if bricks < 0 {
						return gap.start
					}
					ladders -= 1
					// push current gap back to the heap
					heap.Push(&hp, gap.gap)
				} else {
					// we can use the ladder to fill the gap, no need to push it back to the heap
					ladders -= 1
				}
			}

		}
	}

	return len(heights) - 1

}
