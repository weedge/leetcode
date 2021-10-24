package main

// dp[i] = Max(d[j]+1,dp[i]) ( 0<=j<i; nums[i]>nums[j])
// dp[i] 在i位置的最长严格递增子序列的长度
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n==0{
		return 0
	}
	Max := func(a,b int) int{
		if a>b{
			return a
		}
		return b
	}
	dp := make([]int,n)
	max := 1
	for i:=0;i<n;i++{
		dp[i] = 1
		for j:=0;j<i;j++{
			if nums[i]>nums[j]{
				dp[i] = Max(dp[j]+1,dp[i])
			}
			if max<dp[i]{
				max = dp[i]
			}
		}
	}
	return max
}

