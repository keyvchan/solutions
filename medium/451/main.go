package main

import "fmt"

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}

func frequencySort(s string) string {
	bucket := make([][]rune, len(s)+1)
	counts_map := make(map[rune]int)
	for _, v := range s {
		if _, ok := counts_map[v]; ok {
			counts_map[v] = counts_map[v] + 1
		} else {
			counts_map[v] = 1
		}
	}
	for k, v := range counts_map {
		bucket[v] = append(bucket[v], k)
	}
	var s1 string
	for bucket_idx := len(bucket) - 1; bucket_idx >= 0; bucket_idx-- {
		if bucket[bucket_idx] == nil {
			continue
		}
		for _, v := range bucket[bucket_idx] {
			for i := 0; i < counts_map[v]; i++ {
				s1 = s1 + string(v)
			}
		}
	}

	return s1
}
