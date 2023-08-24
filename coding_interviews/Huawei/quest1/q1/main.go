package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Err() == io.EOF {
			break
		}
		input := scanner.Text()
		scanner.Scan()
		if scanner.Err() == io.EOF {
			break
		}
		words := scanner.Text()
		new_words := strings.Split(words, " ")

		if len(new_words) == 0 || len(input) == 0 {
			break
		}

		fmt.Println(Solution(input, new_words))
	}

}

func Solution(s string, words []string) string {
	if s == "" {
		return ""
	}
	if len(words) == 0 {
		return s
	}
	ss := strings.Split(s, " ")

	quate := false

	for i, word := range ss {
		// search for index
		// check if words is in " "
		// check words is "
		if strings.Contains(word, "\"") {
			if quate == false {
				quate = true
			} else {
				quate = false
			}
			continue
		}

		if quate == true {
			continue
		}

		last_index := -1
		for j, w := range words {
			// match ignore case
			if strings.EqualFold(word, w) {
				// replace with index
				last_index = j
			}
		}
		if last_index != -1 {
			ss[i] = strconv.Itoa(last_index)
		}
	}

	// join
	result := strings.Join(ss, " ")

	return result
}
