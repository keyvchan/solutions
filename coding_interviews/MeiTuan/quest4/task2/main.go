package main

import (
	"container/list"
	"fmt"
	"io"
)

func main() {
	var n, m int
	var x []int
	for {
		_, err := fmt.Scan(&n, &m)
		if err == io.EOF {
			break
		}

		x = make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&x[i])
		}

		result := Solution(n, x)
		for _, v := range result {
			fmt.Print(v, " ")
		}

	}

}

func Solution(n int, x []int) []int {
	// create a linked list for O(1) replace
	// create a map for O(1) search
	chain := list.New()
	node_map := map[int]*list.Element{}

	for i := 1; i <= n; i++ {
		chain.PushBack(i)
		element := chain.Back()
		node_map[i] = element
	}

	for _, v := range x {
		// take out the node from the linked list
		element := node_map[v]
		chain.Remove(element)
		// place to the front of the linked list
		chain.PushFront(v)
		// update the map
		node_map[v] = chain.Front()
	}

	// convert the linked list to an array
	result := []int{}
	for e := chain.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(int))
	}

	return result

}
