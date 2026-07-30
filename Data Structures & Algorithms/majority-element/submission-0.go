func majorityElement(nums []int) int {
	target := len(nums) / 2
	if len(nums) % 2 != 0 {
		target++
	}

	hashmap := make(map[int]int)
	
	for _, num := range nums {
		hashmap[num] = hashmap[num] + 1

		if hashmap[num] == target {
			return num
		}
	}

	return -1
}