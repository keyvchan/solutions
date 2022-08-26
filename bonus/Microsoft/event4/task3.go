package main

func main() {

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int, B []int) int {
	min_num := 1000001
	max_num := -1
	for i := 0; i < len(A); i++ {
		if A[i] == B[i] {
			// A[i] == B[i], its not possible to exclude this element
			max_num = max(max_num, A[i])
			continue
		}
		if A[i] < B[i] {
			// we could exclude A[i]
			min_num = min(A[i], min_num)
			max_num = max(max_num, B[i])
		} else {
			// we could exclude B[i]
			min_num = min(B[i], min_num)
			max_num = max(max_num, A[i])
		}
	}
	// write your code in Go (Go 1.4)
	if min_num == 1000001 {
		return max_num + 1
	}
	return min_num
}
