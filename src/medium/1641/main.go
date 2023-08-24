package main

func main() {

}

// 5
// 5 + 4 + 3 + 2 + 1
// 5 + 4 + 3 + 2 + 1 4 + 3 + 2 + 1 3 + 2 + 1 2 + 1 1

func countVowelStrings(n int) int {
	expand := map[int]int{}
	expand[5] = 15
	expand[4] = 10
	expand[3] = 6
	expand[2] = 3
	expand[1] = 1

	for i := 0; i < n; i++ {
	}

}
