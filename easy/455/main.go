package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println(findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8}))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	fmt.Println(g)
	fmt.Println(s)

	count := 0
	j := 0
	i := 0
	for {
		fmt.Println(i, j)
		if i > len(g)-1 || j > len(s)-1 {
			break
		}
		if g[i] <= s[j] {
			count++
			i++
			j++
		} else {
			j++
		}
	}

	return count
}
