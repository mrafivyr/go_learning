package main

func createMatrix(rows, cols int) [][]int {
	rowsSlice := make([][]int, rows)
	for rowIndex := 0; rowIndex < rows; rowIndex++ {
		colsSlice := make([]int, cols)
		for colIndex := 0; colIndex < cols; colIndex++ {
			colsSlice[colIndex] = rowIndex * colIndex
		}
		rowsSlice[rowIndex] = colsSlice
	}
	return rowsSlice
}

// non negative check at the start scenario
/*
func createMatrix(rows, cols int) ([][]int, error) {
	if rows < 0 || cols < 0 {
		return nil, errors.New("rows and columns must be non-negative")
	}
	// ... rest of the code
}
*/
