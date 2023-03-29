package main

func main() {

}

func twoEdit(word1 string, word2 string) bool {
	editCount := 2

	for i, char := range word1 {
		if word2[i] != byte(char) {
			editCount--
		}
		if editCount < 0 {
			return false
		}
	}
	return true

}

func twoEditWords(queries []string, dictionary []string) []string {
	// for every word in queries, check if it is in dictionary or by two edits

	result := []string{}
	for _, word := range queries {
		for _, dictWord := range dictionary {
			if (word == dictWord) || twoEdit(word, dictWord) {
				result = append(result, word)
				break
			}
		}

	}
	return result

}
