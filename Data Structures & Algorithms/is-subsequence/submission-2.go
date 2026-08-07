func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	runes := []rune(s)
	index := 0

	for _, symbol := range t {
		if runes[index] == symbol {
			index++

			if index >= len(runes) {
				return true
			}
		}
	}

	return false
}
