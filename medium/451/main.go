package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}

type counts struct {
	k rune
	v int
}

func frequencySort(s string) string {
	counts_map := make(map[rune]int)
	for _, v := range s {
		counts_map[v] = counts_map[v] + 1
	}

	sorted_counts := make([]counts, len(counts_map))

	for k, v := range counts_map {
		sorted_counts = append(sorted_counts, counts{k, v})
	}
	sort.Slice(sorted_counts, func(i, j int) bool {
		return sorted_counts[i].v > sorted_counts[j].v
	})

	var s1 strings.Builder
	s1.Grow(len(s))

	for _, s := range sorted_counts {
		s1.Write([]byte(strings.Repeat(string(s.k), s.v)))
	}

	return s1.String()
}
