func isPalindrome(s string) bool {
	runes := []rune(s)

	leftIndex := 0
	rightIndex := len(runes) - 1

	for leftIndex < rightIndex {
		leftValue := runes[leftIndex]
		if !(unicode.IsLetter(leftValue) || unicode.IsDigit(leftValue)) {
			leftIndex++
			continue
		}

		rightValue := runes[rightIndex]
		if !(unicode.IsLetter(rightValue) || unicode.IsDigit(leftValue)) {
			rightIndex--
			continue
		}

		if unicode.ToLower(leftValue) == unicode.ToLower(rightValue) {
			leftIndex++
			rightIndex--
			continue
		}

		return false
	}

	return true
}
