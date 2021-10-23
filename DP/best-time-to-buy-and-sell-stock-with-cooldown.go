package main

// many times with cool down
// 需求：
// 多次交易，没有持股后冷冻期一天，获取最大收益
// 三个状态：
// 0.不持股且当天没卖出
// 1.持股(冷冻期后买入)
// 2.不持股且当天卖出了(持股后卖出)
// 递推推演 定义模型公式：
// 第i天不持股且当天没卖出最大收益 dp[i][0] = max(dp[i-1][0],dp[i-1][2])
// 第i天持有股票最大收益 dp[i][1] = max(dp[i-1][1],dp[i-1][0]-price[i])
// 第i天不持股且当天卖出最大收益 dp[i][2] = dp[i-1][1]+price[i]
// 初始：
// dp[0][0] = 0
// dp[0][1] = -price[0]
// dp[0][2] = 0
// 最大收益
// max(dp[n-1][0],dp[n-1][2])
func maxProfitWithCoolDown(prices []int) int {
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
	dp := make([][3]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}

	return Max(dp[n-1][0], dp[n-1][2])
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	println(maxProfitWithCoolDown(prices))
	return
}
