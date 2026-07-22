func mySqrt(x int) int {
	left := 1
	right := x

	for left <= right {
		mid := (left + right) / 2

		if mid <= x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
