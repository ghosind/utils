package utils

// Matrix creates a rows * columns matrix, and initialize it with the values.
//
//	Matrix[int](2, 2) // [][]int{{0, 0}, {0, 0}}
//	Matrix(2, 2, 1, 2, 3, 4) // [][]int{{1, 2}, {3, 4}}
//	Matrix(2, 2, 1, 2, 3) // [][]int{{1, 2}, {3, 0}}
func Matrix[T any](rows, columns int, values ...T) [][]T {
	rows = Max(rows, 0)
	columns = Max(columns, 0)

	matrix := make([][]T, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]T, columns)
	}

	if len(values) > 0 {
		for i, v := range values {
			r := i / columns
			c := i % columns
			if r >= rows {
				break
			}

			matrix[r][c] = v
		}
	}

	return matrix
}
