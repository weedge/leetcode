package main

// 0.偷窃第i间房屋，那么就不能偷窃第i-1间房屋，偷窃总金额为前i-2间房屋的最高总金额与第i间房屋的金额之和。
// 1.不偷窃第i间房屋，偷窃总金额为前i−1间房屋的最高总金额。
// dp[i]=max(dp[i−2]+nums[i],dp[i−1])
// dp[i] 前i间房屋能偷窃到的最高总金额
// 初始：
// dp[0] = nums[0]
// dp[1] = max(nums[0],nums[1])
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = Max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[n-1]
}
