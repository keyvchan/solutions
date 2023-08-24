package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}

func minDistance(word1 string, word2 string) int {

	delete_count := 0
	if len(word1) > len(word2) {
		delete_count = len(word1) - len(word2)

	}

	insert_count := 0
	if len(word1) < len(word2) {
		insert_count = len(word2) - len(word1)
	}

	sub_string_count := map[string]int{}
	sub_string := ""
	for i := 0, 0; i < len(word1); i++ {

	}

	fmt.Println(sub_string_count)

	max := 0
	for _, v := range sub_string_count {
		if v > max {
			max = v
		}
	}

	fmt.Println(delete_count)
	fmt.Println(insert_count)

	return delete_count + insert_count + len(word2) - max
}
