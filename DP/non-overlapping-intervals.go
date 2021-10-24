package main

import (
	"sort"

	"github.com/weedge/lib/mathx"
	"github.com/weedge/lib/sortx"
)

// dp[i] = max(dp[i],dp[j]+1)
// (0<=j<i; intervals[i][0]>=intervals[j][1])
// dp[i] 在i位置满足区间不重叠的最大区间数量
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	it := sortx.TwoDimensionalArray(intervals)
	sort.Sort(it)

	max := 1
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if it[i][0] >= it[j][1] {
				dp[i] = mathx.Max(dp[i], dp[j]+1)
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}

	return n - max
}

// 贪心
// sort by right asc -> compare
func eraseOverlapIntervalsV2(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	//it := utils.TwoDimensionalArray(intervals)
	//sort.Sort(it)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	it := intervals

	max := 1
	r := 0
	for i := 1; i < n; i++ {
		if it[i][0] >= it[r][1] {
			max++
			r = i
		}
	}

	return n - max
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	minErase := eraseOverlapIntervals(intervals)
	minEraseV2 := eraseOverlapIntervalsV2(intervals)

	println(minErase,minEraseV2)
}
