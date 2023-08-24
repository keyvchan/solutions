package main

func main() {

}

type NumMatrix struct {
	col_sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	col_sum := make([][]int, 0, len(matrix[0]))
	for _, row := range matrix {
		sum := 0
		col := make([]int, 0, len(row))
		for _, v := range row {
			sum += v
			col = append(col, sum)
		}
		col_sum = append(col_sum, col)
	}
	return NumMatrix{col_sum}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	sum := 0
	for i := row1; i <= row2; i++ {
		if col1 == col2 {
			sum += this.col_sum[i][col1]
		} else {
			sum += (this.col_sum[i][col2] - this.col_sum[i][col1-1])
		}
	}
	return sum

}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
