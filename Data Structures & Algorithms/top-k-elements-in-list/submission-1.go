func topKFrequent(nums []int, k int) []int {
	var freqs = make(map[int]int)
	for _, number := range nums {
		freqs[number]++
	}

	var buckets = make([][]int, len(nums) + 1)
	for number, freq := range freqs {
		if buckets[freq] == nil {
			buckets[freq] = make([]int, 0)
		}

		buckets[freq] = append(buckets[freq], number)
	}

	var result = make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		var bucket = buckets[i]
		if bucket == nil {
			continue
		}

		for j := 0; j < len(bucket) && len(result) < k; j++ {
			result = append(result, bucket[j])
		}
	}

	return result
}
