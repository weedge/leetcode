package main

// dp[i] = Max(d[j]+1,dp[i]) ( 0<=j<i; nums[i]>nums[j])
// dp[i] 在i位置的最长严格递增子序列的长度
// cn[i] 在i位置最长递增子序列的个数
// 在nums[i] > nums[j]前提下，如果在[0, i-1]的范围内，找到了j，使得dp[j] + 1 > dp[i]，说明找到了一个更长的递增子序列。
// 那么在j位置最长递增子序列的个数，就是最新的在i位置最长递增子序列的个数，即：count[i] = count[j]。
// 在nums[i] > nums[j]前提下，如果在[0, i-1]的范围内，找到了j，使得dp[j] + 1 == dp[i]，说明找到了两个相同长度的递增子序列。
// 那么在i位置最长递增子序列的个数 就应该加上在j位置最长递增子序列的个数，即：count[i] += count[j];
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	cn := make([]int, n)
	max := 1
	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < n; i++ {
		dp[i] = 1
		cn[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					cn[i] = cn[j]
				} else if dp[j]+1 == dp[i] {
					cn[i] += cn[j]
				}
				dp[i] = Max(dp[j]+1, dp[i])
			}

			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	result := 0
	//fmt.Printf("dp:%v cn:%v max:%d",dp,cn,max)
	for i := 0; i < n; i++ {
		if max == dp[i] {
			result += cn[i]
		}
	}
	return result
}

func main() {
	nums := []int{2, 2, 2, 2, 2}
	println(findNumberOfLIS(nums))
}
