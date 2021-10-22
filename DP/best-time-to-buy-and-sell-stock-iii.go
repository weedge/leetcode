package main

// twice
// 一天一共就有五个状态:
//	0.没有操作
//	1.第一次买入
//	2.第一次卖出
//	3.第二次买入
//	4.第二次卖出
// DP[i][j]中 i表示第i天，j为 [0 - 4] 五个状态，DP[i][j]表示第i天状态j所剩最大现金。
// 第i天没有操作 DP[i][0] = DP[i-1][0]
// 第i天第一次买入所剩最大现金 DP[i][1] = max(DP[i-1][0] - prices[i], DP[i-1][1]);
// 第i天第一次卖出所剩最大现金 DP[i][2] = max(DP[i-1][1] + prices[i], DP[i-1][2]);
// 第i天第二次买入所剩最大现金 DP[i][3] = max(DP[i-1][2] - prices[i], DP[i-1][3]);
// 第i天第二次卖出所剩最大现金 DP[i][4] = max(DP[i-1][3] + prices[i], DP[i-1][4]);
// 初始：
// 第0天没有操作 DP[0][0] = 0 // 穷光🥚，初始没钱作为启动资金
// 第0天第一次买入所剩最大现金 DP[0][1] = -prices[0]
// 第0天第一次卖出所剩最大现金 DP[0][2] = 0
// 第0天第二次买入所剩最大现金 DP[0][3] = -prices[0]
// 第0天第二次卖出所剩最大现金 DP[0][4] = 0
func maxProfitIII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max := 0
	n := len(prices)
	dp := make([][]int, n)
	dp[0] = make([]int, 5)
	dp[0][0], dp[0][2], dp[0][4] = 0, 0, 0
	dp[0][1], dp[0][3] = -prices[0], -prices[0]
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 5)
		dp[i][0] = dp[i-1][0]
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = Max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = Max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = Max(dp[i-1][4], dp[i-1][3]+prices[i])

		if dp[i][4] > max {
			max = dp[i][4]
		}
	}

	return max
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	println(maxProfitIII(prices))

	return
}
