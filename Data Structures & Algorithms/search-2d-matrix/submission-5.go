func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rowsCount := len(matrix)
	colsCount := len(matrix[0])

	left := 0
	right := rowsCount - 1

	for left <= right {
		mid := (left + right) / 2

		if matrix[mid][colsCount-1] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left >= rowsCount {
		return false
	}

	row := matrix[left]

	left = 0
	right = colsCount - 1

	for left <= right {
		mid := (left + right) / 2

		switch {
		case row[mid] < target:
			left = mid + 1
		case row[mid] > target:
			right = mid - 1
		default:
			return true
		}
	}

	return false
}
