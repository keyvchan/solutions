package main

func main() {

}

var (
	up        = []int{0, 1}
	down      = []int{0, -1}
	left      = []int{-1, 0}
	right     = []int{1, 0}
	upLeft    = []int{-1, 1}
	upRight   = []int{1, 1}
	downLeft  = []int{-1, -1}
	downRight = []int{1, -1}
)

type Position struct {
	x int
	y int
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	// explore all directions
	// if queen is found, add to result
	result := [][]int{}

	// use builtin complex type
	queensMap := map[Position]bool{}
	for _, queen := range queens {
		queensMap[Position{queen[0], queen[1]}] = true
	}

	// up
	up_x, up_y := king[0], king[1]
	for up_y < 8 {
		// check if the x, y is out of bound

		if queensMap[Position{up_x, up_y}] {
			result = append(result, []int{up_x, up_y})
			break
		}
		up_y += 1
	}

	// down
	down_x, down_y := king[0], king[1]
	for down_y > 0 {
		if queensMap[Position{down_x, down_y}] {
			result = append(result, []int{down_x, down_y})
			break
		}
		down_y -= 1
	}
	// left
	left_x, left_y := king[0], king[1]
	for left_x > 0 {
		if queensMap[Position{left_x, left_y}] {
			result = append(result, []int{left_x, left_y})
			break
		}
		left_x -= 1
	}
	// right
	right_x, right_y := king[0], king[1]
	for right_x < 8 {
		if queensMap[Position{right_x, right_y}] {
			result = append(result, []int{right_x, right_y})
			break
		}
		right_x += 1
	}
	// upLeft
	upLeft_x, upLeft_y := king[0], king[1]
	for upLeft_x > 0 && upLeft_y < 8 {
		if queensMap[Position{upLeft_x, upLeft_y}] {
			result = append(result, []int{upLeft_x, upLeft_y})
			break
		}
		upLeft_x -= 1
		upLeft_y += 1
	}
	// upRight
	upRight_x, upRight_y := king[0], king[1]
	for upRight_x < 8 && upRight_y < 8 {
		if queensMap[Position{upRight_x, upRight_y}] {
			result = append(result, []int{upRight_x, upRight_y})
			break
		}
		upRight_x += 1
		upRight_y += 1
	}
	// downLeft
	downLeft_x, downLeft_y := king[0], king[1]
	for downLeft_x > 0 && downLeft_y > 0 {
		if queensMap[Position{downLeft_x, downLeft_y}] {
			result = append(result, []int{downLeft_x, downLeft_y})
			break
		}
		downLeft_x -= 1
		downLeft_y -= 1
	}
	// downRight
	downRight_x, downRight_y := king[0], king[1]
	for downRight_x < 8 && downRight_y > 0 {
		if queensMap[Position{downRight_x, downRight_y}] {
			result = append(result, []int{downRight_x, downRight_y})
			break
		}
		downRight_x += 1
		downRight_y -= 1
	}

	return result
}
