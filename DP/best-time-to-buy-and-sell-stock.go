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

// At most once
// 需求：
// 一次交易获取最大收益
// 两个状态：
// 0. 一次持有股票
// 1. 一次没有持有股票
// 递推推演 定义模型公式：
// 第i天一次持有股票最大收益 dp[i][0] = max(dp[i-1][0],-price[i])
// 第i天一次没有持有股票最大收益 dp[i][1] = max(dp[i-1][1],dp[i-1][0]+price[i])
// 初始：
// dp[0][0] = -price[0]
// dp[0][1] = 0
// 最大收益
// dp[n-1][1]
func maxProfitV3(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], -prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[n-1][1]
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
