package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sortSentence("is2 This1 Test4 a3"))
	fmt.Println(sortSentence("Myself2 Me1 I4 and3"))
}

func sortSentence(s string) string {
	sbustrings := strings.Split(s, " ")
	sortstrings := make([]string, len(sbustrings))
	for _, substring := range sbustrings {
		index, err := strconv.Atoi(string(substring[len(substring)-1]))
		if err != nil {
			log.Fatal(err)
		}

		sortstrings[index-1] = substring[:len(substring)-1]
	}

	var final_string string
	for i, str := range sortstrings {

		if i == len(sortstrings)-1 {
			final_string += str
		} else {
			final_string += str + " "
		}
	}

	return final_string
}
