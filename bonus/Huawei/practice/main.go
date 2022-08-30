package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	for {
		r := bufio.NewReader(os.Stdin)
		input, err := r.ReadString('\n')
		fmt.Printf("%.2f\n", Solution(input))
		if err == io.EOF {
			break
		}
	}

}

func Solution(s string) float64 {
	ss := strings.Split(s, " ")

	result := 0

	for _, sss := range ss {
		result += len(sss)
	}

	return float64(result) / float64(len(ss))

}
