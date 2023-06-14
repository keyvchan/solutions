package main

import "fmt"

func main() {
	solve([][]byte{
		{'O', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'O', 'X'},
		{'X', 'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'X'},
	})
}

type Position struct {
	x int
	y int
}

func dfs(board [][]byte, i, j int, visited map[Position]bool) {
	if board[i][j] == 'X' {
		return
	}
	visited[Position{i, j}] = true
	if i > 0 && !visited[Position{i - 1, j}] {
		dfs(board, i-1, j, visited)
	}
	if i < len(board)-1 && !visited[Position{i + 1, j}] {
		dfs(board, i+1, j, visited)
	}
	if j > 0 && !visited[Position{i, j - 1}] {
		dfs(board, i, j-1, visited)
	}
	if j < len(board[0])-1 && !visited[Position{i, j + 1}] {
		dfs(board, i, j+1, visited)
	}
}

func solve(board [][]byte) {

	// all positions can't be flaped
	all_positions := map[Position]bool{}
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == 'O' && !all_positions[Position{0, i}] {
			visited := map[Position]bool{}
			dfs(board, 0, i, visited)
			// add all positions to the map
			for k := range visited {
				all_positions[k] = true
			}
		}

		if board[len(board)-1][i] == 'O' && !all_positions[Position{len(board) - 1, i}] {
			visited := map[Position]bool{}
			dfs(board, len(board)-1, i, visited)
			// add all positions to the map
			for k := range visited {
				all_positions[k] = true
			}
		}
	}

	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' && !all_positions[Position{i, 0}] {

			visited := map[Position]bool{}
			dfs(board, i, 0, visited)
			// add all positions to the map
			for k := range visited {
				all_positions[k] = true
			}
		}
		if board[i][len(board[0])-1] == 'O' && !all_positions[Position{i, len(board[0]) - 1}] {

			visited := map[Position]bool{}
			dfs(board, i, len(board[0])-1, visited)
			// add all positions to the map
			for k := range visited {
				all_positions[k] = true
			}
		}
	}
	fmt.Println(all_positions)

	// edit board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if !all_positions[Position{i, j}] && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}
