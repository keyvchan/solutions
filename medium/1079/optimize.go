package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numTilePossibilities("AAB"))
	fmt.Println(numTilePossibilities("CDC"))
	fmt.Println(numTilePossibilities("ABA"))
}

func backtracking(tiles []rune, result *int, used map[int]bool) {
	*result++

	for i, t := range tiles {
		if used[i] {
			continue
		}
		if (i > 0 && t == tiles[i-1]) && !used[i-1] {
			continue
		}
		used[i] = true
		backtracking(tiles, result, used)
		used[i] = false
	}
}

func numTilePossibilities(tiles string) int {

	new_tiles := []rune(tiles)

	sort.Slice(new_tiles, func(i, j int) bool {
		return new_tiles[i] < new_tiles[j]
	})

	used := map[int]bool{}

	result := 0
	backtracking(new_tiles, &result, used)

	return result - 1
}
