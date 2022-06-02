package main

import "container/heap"

func main() {

}

type hp []int

func (hp hp) Len() int {
	return len(hp)
}

func (hp hp) Less(i, j int) bool {
	return hp[i] > hp[j]
}
func (hp *hp) Pop() interface{} {
	old := *hp
	n := len(old)
	x := old[n-1]
	*hp = old[0 : n-1]
	return x
}

func (hp *hp) Push(x interface{}) {
	*hp = append(*hp, x.(int))
}
func (hp hp) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func lastStoneWeight(stones []int) int {
	if len(stones) < 3 {
		if stones[0] > stones[1] {
			return stones[0] - stones[1]
		} else {
			return stones[1] - stones[0]
		}
	}

	h := hp(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		a, b := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if a != b {
			heap.Push(&h, a-b)
		}
	}

	return h[0]

}
