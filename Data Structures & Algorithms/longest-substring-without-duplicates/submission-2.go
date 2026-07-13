func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[rune]int)
	maxLength := 0
	left := 0

	for right, symbol := range s {
		if prevIdx, ok := lastSeen[symbol]; ok && prevIdx >= left {
			left = prevIdx + 1
		}

		lastSeen[symbol] = right
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
