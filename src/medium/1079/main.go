package main

import "fmt"

func main() {
	fmt.Println(numTilePossibilities("AAB"))
}

func backtracking(tiles string, current string, result map[string]bool) {
	if len(current) > 0 {
		result[current] = true
	}
	if len(tiles) == 0 {
		return
	}

	for i := 0; i < len(tiles); i++ {
		backtracking(tiles[:i]+tiles[i+1:], current+string(tiles[i]), result)
	}

}

func numTilePossibilities(tiles string) int {

	// use a hashmap to avoid duplicates
	result := map[string]bool{}

	backtracking(tiles, "", result)

	return len(result)
}
