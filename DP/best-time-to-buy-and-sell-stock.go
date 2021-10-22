package main

import "fmt"
func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := 1<<32 - 1
	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

// At most once
// DP[i] = max(DP[i-1]+prices[i]-prices[i-1],prices[i]-prices[i-1])
func maxProfitV2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := 0
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		dp[i] = Max(diff, dp[i-1]+diff)
		if dp[i] > max {
			max = dp[i]
		}
	}
	fmt.Printf("DP:%v", dp)
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(maxProfitV2(prices))

	return
}
