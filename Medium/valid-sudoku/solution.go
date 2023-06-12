package validsudoku

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}       // [row][digit]
	columns := [9][9]bool{}    // [column][digit]
	squares := [3][3][9]bool{} // [row/3][column/3][digit]

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]

			if cell == '.' {
				continue
			}

			digit := int(cell-'0') - 1

			if rows[i][digit] || columns[j][digit] || squares[i/3][j/3][digit] {
				return false
			}

			rows[i][digit] = true
			columns[j][digit] = true
			squares[i/3][j/3][digit] = true
		}
	}

	return true
}
