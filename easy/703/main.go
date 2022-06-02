package main

import "container/heap"

func main() {

}

type hp []int

func (hp hp) Len() int {
	return len(hp)
}
func (hp hp) Less(i, j int) bool {
	return hp[i] < hp[j]
}
func (hp hp) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
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

type KthLargest struct {
	k  int
	hp hp
}

func Constructor(k int, nums []int) KthLargest {
	h := hp{}
	heap.Init(&h)
	for _, v := range nums {
		heap.Push(&h, v)
	}
	for len(h) > k {
		heap.Pop(&h)
	}
	return KthLargest{k, h}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.hp, val)
	if len(this.hp) > this.k {
		heap.Pop(&this.hp)
	}
	return this.hp[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
