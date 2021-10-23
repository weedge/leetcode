package main

// twice -> k times
func maxProfitIV(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	dp := make([][]int, n)
	dp[0] = make([]int, 2*k+1)
	dp[0][0] = 0
	for j := 1; j < 2*k+1; j++ {
		if j%2 == 0 {
			dp[0][j] = 0
		} else {
			dp[0][j] = -prices[0]
		}
	}
	for i := 1; i < n; i++ {
		dp[i] = make([]int, 2*k+1)
		dp[i][0] = dp[i-1][0]
		for j := 1; j < 2*k+1; j++ {
			if j%2 == 0 {
				dp[i][j] = Max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			}
		}
	}

	return dp[n-1][2*k]
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	println(maxProfitIV(3, prices))

	return
}
