package main

type N048RotateImage struct {
}

func (this *N048RotateImage) rotate(pmatrix *[][]int) {
	matrix := *pmatrix
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])
	newmatrix := make([][]int, cols)
	for i := 0; i < cols; i++ {
		newrow := make([]int, rows)
		for j := 0; j < rows; j++ {
			newrow[j] = matrix[rows-j-1][i]
		}
		newmatrix[i] = newrow
	}
	*pmatrix = newmatrix
}
