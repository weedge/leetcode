package main

import "github.com/weedge/lib/mathx"

// like maximum-length-of-repeated-subarray.go
// 不要求是连续的了，但要有相对顺序
// dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列为dp[i][j]
// dp[i][j] = dp[i-1][j-1] + 1 (text1[i-1] == text2[j-1])
// dp[i][j] = max(dp[i-1][j], dp[i][j-1]) (text1[i-1] != text2[j-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = mathx.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n][m]
}
