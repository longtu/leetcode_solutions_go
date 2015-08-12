package main

type N006ZigZag struct {
}

func (this *N006ZigZag) convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := numRows
	length := len(s)
	cols := length

	chars := make([][]byte, rows)
	for i := range chars {
		chars[i] = make([]byte, cols)
	}

	index := 0
	i := 0
	j := 0
	result := []byte(s) // Convert string to bytes slice
	for index < length {
		//cout << i << "," << j << endl;
		chars[i][j] = s[index]
		index++
		if j%(rows-1) == 0 {
			if i < rows-1 {
				i++
			} else {
				j++
				i--
			}
		} else {
			j++
			i--
		}
	}

	index = 0
	cols = j + 1
	for ii := 0; ii < rows; ii++ {
		for jj := 0; jj < cols; jj++ {
			if chars[ii][jj] != 0 {
				result[index] = chars[ii][jj]
				index++
			}
		}
	}

	return string(result) // Convert bytes slice to string
}
