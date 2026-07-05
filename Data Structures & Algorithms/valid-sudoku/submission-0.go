func isValidSudoku(board [][]byte) bool {
	subsets := make([][9]byte, 0, 27)

	for i := 0; i < 9; i++ {
		row := [9]byte{}
		col := [9]byte{}
		box := [9]byte{}

		// Collect rows and cols
		for j := 0; j < 9; j++ {
			row[j] = board[i][j]
			col[j] = board[j][i]
		}

		// Collect box
		for j := 0; j < 9; j++ {
			x := i%3*3 + j%3
			y := i/3*3 + j/3

			box[j] = board[x][y]
		}

		subsets = append(subsets, row)
		subsets = append(subsets, col)
		subsets = append(subsets, box)
	}

	for _, subset := range subsets {
		numbersExistence := [9]bool{}

		for i := 0; i < 9; i++ {
			if subset[i] == '.' {
				continue
			}

			number := subset[i]

			if number < '0' || number > '9' || numbersExistence[number-'0'-1] {
				return false
			}

			numbersExistence[number-'0'-1] = true
		}
	}

	return true
}