package qns8_12

func NQueens(n int) [][]int {
	gridSize := n
	var res [][]int

	var iter func(row int, cols []int)
	iter = func(row int, cols []int) {
		if row == gridSize {
			newCols := make([]int, len(cols))
			copy(newCols, cols)
			res = append(res, newCols)
			return
		}
		for col := 0; col < gridSize; col++ {
			if checkValid(cols, row, col) {
				cols[row] = col
				iter(row+1, cols)
			}
		}
	}

	iter(0, make([]int, 8))
	return res
}

func checkValid(cols []int, row1 int, col1 int) bool {
	for row2 := 0; row2 < row1; row2++ {
		col2 := cols[row2]

		if col1 == col2 {
			return false
		}

		columnDistance := abs(col2, col1)
		rowDistance := row1 - row2
		if columnDistance == rowDistance {
			return false
		}
	}
	return true
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
