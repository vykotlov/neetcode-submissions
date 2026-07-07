func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}

		currentProfit := prices[i] - minPrice
		if currentProfit > profit {
			profit = currentProfit
		}
	}

	return profit
}
