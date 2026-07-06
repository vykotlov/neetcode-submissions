func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
			continue
		}
		
		if nums[mid] < target {
			left = mid + 1
			continue
		}
		
		return mid
	}

	if nums[left] == target {
		return left
	}

	return -1
}
