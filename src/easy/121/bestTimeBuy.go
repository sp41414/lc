package bestTimeBuy

func MaxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v-minPrice > maxProfit {
			maxProfit = v - minPrice
		}
	}

	return maxProfit
}
