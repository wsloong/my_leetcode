package main

import "math"

func maxProfit(prices []int) int {
	minPrice, maxPrice := math.MaxInt32, 0

	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxPrice = max(maxPrice, price-minPrice)
	}
	return maxPrice
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
