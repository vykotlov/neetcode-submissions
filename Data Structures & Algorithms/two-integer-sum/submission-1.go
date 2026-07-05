func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int, len(nums))
	for i, num := range nums {
		indexes[num] = i
	}

	for i, num := range nums {
		if j, ok := indexes[target-num]; ok && j != i {
			return []int{i, j}
		}
	}

	return []int{}
}
