package main

func main() {

}

func matrixReshape(mat [][]int, r int, c int) [][]int {

	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	if len(mat)%r != 0 || len(mat[0])%c != 0 {
		return mat
	}

	return mat
}
