package main

import "container/list"

func main() {

}

type Position struct {
	x    int
	y    int
	step int
}

func WallsAndGates(rooms [][]int) {
	// write your code here
	queue := list.New()
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				queue.PushBack(Position{i, j, 0})
			}
		}
	}
	// BFS
	for queue.Len() > 0 {
		element := queue.Remove(queue.Front())
		posion := element.(Position)
		if posion.x > 0 && rooms[posion.x-1][posion.y] == 2147483647 {
			// check step
			if rooms[posion.x-1][posion.y] > posion.step+1 {
				rooms[posion.x-1][posion.y] = posion.step + 1
			}
			queue.PushBack(Position{posion.x - 1, posion.y, posion.step + 1})
		}
		if posion.x < len(rooms)-1 && rooms[posion.x+1][posion.y] == 2147483647 {
			// check step
			if rooms[posion.x+1][posion.y] > posion.step+1 {
				rooms[posion.x+1][posion.y] = posion.step + 1
			}
			queue.PushBack(Position{posion.x + 1, posion.y, posion.step + 1})
		}
		if posion.y > 0 && rooms[posion.x][posion.y-1] == 2147483647 {
			if rooms[posion.x][posion.y-1] > posion.step+1 {
				rooms[posion.x][posion.y-1] = posion.step + 1
			}
			queue.PushBack(Position{posion.x, posion.y - 1, posion.step + 1})
		}
		if posion.y < len(rooms[0])-1 && rooms[posion.x][posion.y+1] == 2147483647 {
			if rooms[posion.x][posion.y+1] > posion.step+1 {
				rooms[posion.x][posion.y+1] = posion.step + 1
			}
			queue.PushBack(Position{posion.x, posion.y + 1, posion.step + 1})
		}

	}
}
