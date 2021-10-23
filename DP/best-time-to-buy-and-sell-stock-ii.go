package main

// many times
// DP[i] = Max(DP[i-1],DP[i-1]+prices[i]-prices[i-1])
func maxProfitII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	dp := make([]int, len(prices))
	dp[0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i] = Max(dp[i-1], dp[i-1]+prices[i]-prices[i-1])
		if max < dp[i] {
			max = dp[i]
		}
	}

	return max

}

// many times
// 需求：
// 多次交易获取最大收益
// 两个状态：
// 0. 持有股票(多次)
// 1. 没有持有股票(多次)
// 递推推演 定义模型公式：
// 第i天持有股票最大收益 dp[i][0] = max(dp[i-1][0],dp[i-1][1]-price[i])
// 第i天没有持有股票最大收益 dp[i][1] = max(dp[i-1][1],dp[i-1][0]+price[i])
// 初始：
// dp[0][0] = -price[0]
// dp[0][1] = 0
// 最大收益
// dp[n-1][1]
func maxProfitIIV2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[n-1][1]
}
