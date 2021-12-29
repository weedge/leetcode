package main

import "fmt"

func findLengthV0(nums1 []int, nums2 []int) int {
	max := 0
	n := len(nums1)
	m := len(nums2)
	s := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			k := 0
			for i+k < n && j+k < m && nums1[i+k] == nums2[j+k] {
				k++
			}
			if k > max {
				max = k
				s = i
			}
		}
	}
	fmt.Println(nums1[s:s+max])
	return max
}

// dp[i][j] ：以下标i - 1为结尾的A，和以下标j - 1为结尾的B，最长重复子数组长度为dp[i][j]。
// dp[i][j] = dp[i-1][j-1] + 1 (A[i-1] == B[j-1])
func findLength(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	max := 0
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	e:=0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > max {
				max = dp[i][j]
				e=i
			}
		}
	}
	fmt.Println(nums1[e-max:e])

	return max
}

func main() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	max := findLength(nums1, nums2)
	println(max)
}
