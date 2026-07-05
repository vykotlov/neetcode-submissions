func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		subset := [9]byte{}

		// Row [i][]
		for j := 0; j < 9; j++ {
			subset[j] = board[i][j]
		}

		if !isValidSubset(subset) {
			return false
		}

		// Column [][j]
		for j := 0; j < 9; j++ {
			subset[j] = board[j][i]
		}

		if !isValidSubset(subset) {
			return false
		}

		// Box [i % 3 * 3][i / 3 * 3]
		for j := 0; j < 9; j++ {
			subset[j] = board[i%3*3+j%3][i/3*3+j/3]
		}

		if !isValidSubset(subset) {
			return false
		}
	}

	return true
}

func isValidSubset(subset [9]byte) bool {
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

	return true
}