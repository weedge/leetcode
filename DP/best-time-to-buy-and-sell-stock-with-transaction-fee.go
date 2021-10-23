package main

// many times
// 需求：
// 多次交易获取最大收益,每次交易需要交付费用
// 两个状态：
// 0. 持有股票(多次)
// 1. 没有持有股票(多次)
// 递推推演 定义模型公式：
// 第i天持有股票最大收益 dp[i][0] = max(dp[i-1][0],dp[i-1][1]-price[i])
// 第i天没有持有股票最大收益 dp[i][1] = max(dp[i-1][1],dp[i-1][0]+price[i]-fee)
// 初始：
// dp[0][0] = -price[0]
// dp[0][1] = 0
// 最大收益
// max(dp[n-1][0],dp[n-1][1])
func maxProfitWithTransactionFree(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][2]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return Max(dp[n-1][0], dp[n-1][1])
}
