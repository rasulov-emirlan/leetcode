package pascalstriangle

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	res := make([][]int, numRows)

	res[0] = make([]int, 1)
	res[0][0] = 1

	for row := 1; row < len(res); row++ {
		res[row] = make([]int, row+1)
		res[row][0] = 1
		for col := 1; col < len(res[row])-1; col++ {
			res[row][col] = res[row-1][col-1] + res[row-1][col]
		}
		res[row][len(res[row])-1] = 1
	}

	return res
}
