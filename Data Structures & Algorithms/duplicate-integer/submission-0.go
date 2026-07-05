func hasDuplicate(nums []int) bool {
	unique := make(map[int]bool, len(nums))

	for _, num := range nums {
		if _, contains := unique[num]; contains {
			return true
		}

		unique[num] = true
	}

	return false
}
