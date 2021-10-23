package main

import (
	"fmt"
)

func maxSubarraySumCircularV1(nums []int) int {
	max := nums[0]
	n := len(nums)
	tmpNums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		tmpNums = []int{}
		tmpNums = append(tmpNums, nums[i:]...)
		tmpNums = append(tmpNums, nums[0:i]...)
		tmpMax := MaxSubArray(tmpNums)
		if tmpMax > max {
			max = tmpMax
		}
		fmt.Printf("%v\n", tmpNums)
	}

	return max
}

// 1. 最大子序列和在原数组中 MaxSubArray
// 2. 最大子序列和在环形数组中 Total - MinSubArray(最小子序列和在原数组中)
// maxSubarraySumCircular = min(MaxSubArray,Total - MinSubArray)
// 原数组都是负数的情况，Total = MinSubArray, 最大值应该是MaxSubArray
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	total := 0
	for i := 0; i < n; i++ {
		total += nums[i]
	}

	max := MaxSubArray(nums)
	min := MinSubArray(nums)
	//fmt.Printf("max:%d min:%d total:%d\n", max, min, total)

	circularMax := total - min
	if total-min == 0 {
		circularMax = max
	}

	return Max(max, circularMax)
}

func MaxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	max := nums[0]
	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = Max(dp[i-1]+nums[i], nums[i])
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func MinSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	Min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	min := nums[0]
	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = Min(dp[i-1]+nums[i], nums[i])
		if min > dp[i] {
			min = dp[i]
		}
	}

	return min
}

func main() {
	//nums := []int{5, -3, 5}
	nums := []int{-2, -3, -1}
	println(maxSubarraySumCircular(nums))
}
