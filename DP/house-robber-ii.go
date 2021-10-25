package main

import "github.com/weedge/lib/mathx"

// 首尾相连，分两种情况：
// 1. 首部开始抢，尾部不能抢，0<=i<n-1 获取最大金额 max1
// 2. 首部不能抢，尾部可以抢，0<i<=n-1 获取最大金额 max2
// 最终获取最大金额： max = max(max1,max2)
// 获取最大金额方式和没有首尾相连的情况一样
// dp[i] = max(dp[i-2]+nums[i],dp[i-1])
// 初始：
// dp[0] = nums[0]
// dp[1] = max(nums[0],nums[1])
func robII(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	max1 := Rob(nums[:n-1])
	max2 := Rob(nums[1:n])

	return mathx.Max(max1, max2)
}

func Rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = mathx.Max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = mathx.Max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[n-1]
}
