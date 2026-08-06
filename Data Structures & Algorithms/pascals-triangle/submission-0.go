func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	result := [][]int{}
	result = append(result, []int{1})

	for i := 1; i < numRows; i++ {
		curr := []int{1}
		prev := result[i - 1]

		for j := 1; j < len(prev); j++ {
			curr = append(curr, prev[j - 1] + prev[j])
		}

		curr = append(curr, 1)
		result = append(result, curr)
	}

	return result
}
