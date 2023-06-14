package main

import "fmt"

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

func main() {

	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("DCXXI"))
	fmt.Println(romanToInt("MCMIV"))

}

func romanToInt(s string) int {

	I := 1
	V := 5
	X := 10
	L := 50
	C := 100
	D := 500
	M := 1000

	var arr []int

	for _, c := range s {
		switch c {
		case 'I':
			arr = append(arr, I)
		case 'V':
			arr = append(arr, V)
		case 'X':
			arr = append(arr, X)
		case 'L':
			arr = append(arr, L)
		case 'C':
			arr = append(arr, C)
		case 'D':
			arr = append(arr, D)
		case 'M':
			arr = append(arr, M)
		}
	}

	sum := 0

	for i := 0; i < len(arr); i++ {
		s1 := arr[i]
		if i+1 < len(arr) {
			if s1 >= arr[i+1] {
				sum += s1
			} else {
				sum += arr[i+1] - s1
				i++
			}
		} else {
			sum += s1
		}

	}

	return sum
}
