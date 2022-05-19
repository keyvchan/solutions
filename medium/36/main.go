package main

import "fmt"

func main() {
	fmt.Println(isValidSudoku([][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}))
	fmt.Println(isValidSudoku([][]byte{
		[]byte("83..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}))
}

func isValidSudoku(board [][]byte) bool {

	// check each row
	row_map := make(map[byte]bool, 9)
	column_map := make(map[byte]bool, 9)
	cell_map := make(map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// check row
			if _, ok := row_map[board[i][j]]; ok {
				if board[i][j] != '.' {
					return false
				}
			} else {
				row_map[board[i][j]] = true
			}

			// check column
			if _, ok := column_map[board[j][i]]; ok {
				if board[j][i] != '.' {
					return false
				}
			} else {
				column_map[board[j][i]] = true
			}

			// check cell
			// (j/3, j%3) (j/3, j%3) (j/3, j%3) || (j/3, i*3 + j%3) (j/3, i*3 + j%3) (j/3, i*3 + j%3)
			// (j/3, j%3) (j/3, j%3) (j/3, j%3) ||
			// (j/3, j%3) (j/3, j%3) (j/3, j%3) ||
			// (i%3*3 + j/3, j%3) (i%3*3 + j/3, j%3) (i%3*3 + j/3, i%3*3 + j%3) || (i%3*3, i*3 + j%3) (i%3*3, i*3 + j%3) (i%3*3, i*3 + j%3)

			// fmt.Println(j/3, j%3)
			if _, ok := cell_map[board[i/3*3+j/3][i%3*3+j%3]]; ok {
				if board[i/3*3+j/3][i%3*3+j%3] != '.' {
					return false
				}
			} else {
				cell_map[board[i/3*3+j/3][i%3*3+j%3]] = true
			}
		}
		row_map = map[byte]bool{}
		column_map = map[byte]bool{}
		cell_map = map[byte]bool{}
	}

	return true
}
