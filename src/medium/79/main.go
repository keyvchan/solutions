package main

import "fmt"

func main() {

	board := [][]byte{
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		// "AAAAAAAAAAAAAAB"
	}
	fmt.Println(exist(board, "AAAAAAAAAAAAAAB"))
}

type Position struct {
	x int
	y int
}

func backtracking(board [][]byte, current_position Position, current_word string, position_map map[Position]bool, target_word string) bool {

	if current_word == target_word {
		return true
	}
	result := false
	// current_word is part of the target word
	// check current_position left right up down
	// check left right up down if they are in the board
	if current_position.x-1 >= 0 && board[current_position.x-1][current_position.y] == target_word[len(current_word)] {
		left := Position{current_position.x - 1, current_position.y}
		if !position_map[left] {
			// process left
			// add to position_map
			// copy position_map to new map
			position_map[left] = true
			result = result || backtracking(board, left, current_word+string(board[left.x][left.y]), position_map, target_word)
			delete(position_map, left)
		}
	}
	if current_position.x+1 < len(board) && board[current_position.x+1][current_position.y] == target_word[len(current_word)] {
		right := Position{current_position.x + 1, current_position.y}
		if !position_map[right] {
			// process right
			// add to position_map
			// copy position_map to new map
			position_map[right] = true
			result = result || backtracking(board, right, current_word+string(board[right.x][right.y]), position_map, target_word)
			delete(position_map, right)
		}
	}
	if current_position.y-1 >= 0 && board[current_position.x][current_position.y-1] == target_word[len(current_word)] {
		up := Position{current_position.x, current_position.y - 1}
		if !position_map[up] {
			// process up
			// add to position_map
			// copy position_map to new map
			position_map[up] = true
			result = result || backtracking(board, up, current_word+string(board[up.x][up.y]), position_map, target_word)
			delete(position_map, up)
		}
	}
	if current_position.y+1 < len(board[0]) && board[current_position.x][current_position.y+1] == target_word[len(current_word)] {
		down := Position{current_position.x, current_position.y + 1}
		if !position_map[down] {
			// process down
			// add to position_map
			// copy position_map to new map
			position_map[down] = true
			result = result || backtracking(board, down, current_word+string(board[down.x][down.y]), position_map, target_word)
			delete(position_map, down)
		}
	}
	return result

}

func exist(board [][]byte, word string) bool {

	searched_positions := map[Position]bool{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				searched_positions[Position{i, j}] = true
				if backtracking(board, Position{i, j}, string(board[i][j]), searched_positions, word) {
					return true
				}
				delete(searched_positions, Position{i, j})
			}
		}
	}
	return false

}
