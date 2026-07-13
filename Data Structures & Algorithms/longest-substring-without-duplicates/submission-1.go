func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	letters := make(map[rune]int)
	maxLength := 0

	for i := 0; i < len(runes); i++ {
		if index, ok := letters[runes[i]]; ok {
			maxLength = max(maxLength, len(letters))
			letters = make(map[rune]int)

			for j := index; j < i; j++ {
				letters[runes[j]] = j
			}
		}

		letters[runes[i]] = i
	}

	return max(maxLength, len(letters))
}
