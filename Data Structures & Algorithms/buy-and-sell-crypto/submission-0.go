func maxProfit(prices []int) int {
	low := prices[0]
	high := prices[0]
	res := 0
	for _,price := range prices {
		if price < low {
			low = price
			high = price
		} else if price > high {
			high = price
			res = max(res, high-low)
		}


	}
	return res
}
