package main

import "fmt"

func main() {
	fmt.Println(percentageLetter("Hello World", 'l'))
	fmt.Println(percentageLetter("foobar", 'o'))
	fmt.Println(percentageLetter("jjjj", 'k'))
}

func percentageLetter(s string, letter byte) int {

	// count the number of all letters in the string

	var count int

	for _, b := range s {

		if byte(b) == letter {
			count++
		}
	}

	return int(float64(count) * 100 / float64(len(s)))
}
