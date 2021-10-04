package helpers

func CreateMatrixInt(n, m int) [][]int {
	output := make([][]int, n)
	for i := range output {
		output[i] = make([]int, m)
	}
	return output
}

func FillMatrix(matrix [][]int, value int) {
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = value
		}
	}
}
