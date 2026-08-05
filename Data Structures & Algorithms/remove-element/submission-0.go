func removeElement(nums []int, val int) int {
	current, counter := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			counter++
			continue
		}

		nums[current] = nums[i]
		current++
	}

	for i := current; i < len(nums); i++ {
		nums[i] = 0
	}

	return len(nums) - counter
}
