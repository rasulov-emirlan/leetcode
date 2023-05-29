package rotateimage

// solved with help from neetcode

func rotate(matrix [][]int) {
	if len(matrix) < 2 {
		return
	}

	l, r := 0, len(matrix)-1

	for l < r {
		bounds := r - l
		for i := 0; i < bounds; i++ {
			top, bottom := l, r

			topLeft := matrix[top][l+i]
			matrix[top][l+i] = matrix[bottom-i][l]
			matrix[bottom-i][l] = matrix[bottom][r-i]
			matrix[bottom][r-i] = matrix[top+i][r]
			matrix[top+i][r] = topLeft
		}
		r--
		l++
	}
}
