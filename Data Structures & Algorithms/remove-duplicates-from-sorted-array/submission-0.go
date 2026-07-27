func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	uniqueIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[uniqueIndex] {
			continue
		}
		
		uniqueIndex++
		nums[uniqueIndex] = nums[i]
	}

	return uniqueIndex + 1
}
