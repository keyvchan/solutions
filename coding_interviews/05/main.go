package main

import "strings"

func main() {

}

func replaceSpace(s string) string {
	// convert input to rune array
	news := strings.ReplaceAll(s, " ", "%20", )

	return news
}
