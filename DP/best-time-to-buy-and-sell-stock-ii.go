package main

// many times
// DP[i] = Max(DP[i-1],DP[i-1]+prices[i]-prices[i-1])
func maxProfitII(prices []int) int {
	if len(prices) == 0{
		return 0
	}
	max := 0
	dp := make([]int,len(prices))
	dp[0]=0
	for i:=1;i<len(prices);i++{
		dp[i] = Max(dp[i-1],dp[i-1]+prices[i]-prices[i-1])
		if max<dp[i]{
			max=dp[i]
		}
	}

	return max

}

