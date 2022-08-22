package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// type Ordered interface {
//   comparable
//   // FIXME: type of passed
//   Less(other any) bool
// }

type Heap[E constraints.Ordered] struct {
	// we use a array to store all nodes
	nodes []E
}

func New[E constraints.Ordered]() Heap[E] {

	h := Heap[E]{}
	h.nodes = make([]E, 1)

	return h
}

func (h *Heap[E]) Push(x E) {
	// check capacity
	h.nodes = append(h.nodes, x)
	// push a new node to the heap




	// heapify the new node
	// - first check if the new node is satisfy the heap property
	// - if not, then heapify the new node
	h.up(len(h.nodes) - 1)
}

func (h *Heap[E]) Pop() any {
	// swap the top node and the last node
	h.nodes[0], h.nodes[len(h.nodes)-1] = h.nodes[len(h.nodes)-1], h.nodes[0]
	top := h.nodes[len(h.nodes)-1]
	// shink
	h.nodes = h.nodes[:len(h.nodes)-1]

	h.down(0)

	return top
}

func (h *Heap[E]) down(index int) {
	for {
		// compare the left child and the right child, swap the smaller one with the current node
		left := index*2 + 1
		right := index*2 + 2

		if left >= len(h.nodes) && right >= len(h.nodes) {
			break
		}

		swaped := 0

		// check if the left child is out of range
		if left < len(h.nodes) && right < len(h.nodes) {
			// compare the left child and the right child
			if h.nodes[left] < (h.nodes[right]) {
				// swap the left child with the current node
				swaped = left
			} else {
				// swap the right child with the current node
				swaped = right
			}
		} else if left >= len(h.nodes) {
			// swap right child with the current node
			swaped = right
		} else if right >= len(h.nodes) {
			// swap left child with the current node
			swaped = left
		}
		h.nodes[index], h.nodes[swaped] = h.nodes[swaped], h.nodes[index]
		index = swaped
		if left >= len(h.nodes) || right >= len(h.nodes) {
			break
		}
	}
}

func (h *Heap[E]) up(index int) {

	for {
		// make sure the parent node is less than the current node
		parent := (index - 1) / 2
		// check if parent node index out of range
		if parent < 0 && parent >= len(h.nodes) {
			// this node don't have a parent node
			return
		}
		if index == parent {
			// we are at the top of the heap
			return
		}

		if h.nodes[parent] < (h.nodes[index]) {
			// parent node is less than current node
			return
		}
		// update the indexs map
		// swap the parent node and current node
		h.nodes[parent], h.nodes[index] = h.nodes[index], h.nodes[parent]
		// update the index
		index = parent
	}

}

type MyInt int

// Less
func (i MyInt) Less(j any) bool {
	return i < j.(MyInt)
}

func main() {
	h := New[MyInt]()
	h.Push(1)
	h.Push(2)
	fmt.Println(h)
	h.Push(0)
	fmt.Println(h)
	h.Push(3)
	fmt.Println(h)
	h.Push(1)
	fmt.Println(h)
	h.Pop()
	fmt.Println(h)
	h.Pop()
	fmt.Println(h)
	h.Pop()
	fmt.Println(h)
}
