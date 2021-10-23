package main

import (
	"fmt"
)

// 两个状态：
// 0. 当前最大
// 1. 当前最小
// 第i个元素当前最大乘积（当前元素是正数或者负数的情况） dp[i][0] = max(dp[i-1][0]*nums[i],max(dp[i-1][1]*nums[i],nums[i]))
// 第i个元素当前最小乘积（当前元素是正数或者负数的情况） dp[i][1] = min(dp[i-1][1]*nums[i],min(dp[i-1][0]*nums[i],nums[i]))
// 初始：
// dp[0][0] = nums[0]
// dp[0][1] = nums[1]
// 最大乘积
// max(dp[][0]...)
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, n)
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]

	max := nums[0]

	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	Min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	s, e := 0, 0

	for i := 1; i < n; i++ {
		// 最大(当前值为正 和 负的情况)
		dp[i][0] = Max(dp[i-1][0]*nums[i], Max(dp[i-1][1]*nums[i], nums[i]))
		// 最小(当前值为正 和 负的情况)
		dp[i][1] = Min(dp[i-1][1]*nums[i], Min(dp[i-1][0]*nums[i], nums[i]))

		if max < dp[i][0] {
			max = dp[i][0]
			if dp[i][0] == nums[i] {
				s = i
			}
			e = i
		}
	}
	fmt.Printf("%v\n", nums[s:e+1])

	return max
}

func main() {
	nums := []int{2, 3, -2, 4, -1, -3}
	println(maxProduct(nums))
}
