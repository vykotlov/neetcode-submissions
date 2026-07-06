func isPalindrome(s string) bool {
	runes := []rune(strings.ToLower(s))

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

		if leftValue == rightValue {
			leftIndex++
			rightIndex--
			continue
		}

		return false
	}

	return true
}