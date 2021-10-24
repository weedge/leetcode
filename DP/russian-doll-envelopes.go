package main

import (
	"fmt"
	"sort"
)

type TwoDimensionalArray [][]int

func (e TwoDimensionalArray) Less(i, j int) bool {
	for k := 0; k < len(e[i]); k++ {
		if e[i][k] < e[j][k] {
			return true
		} else if e[i][k] == e[j][k] {
			continue
		} else {
			return false
		}
	}
	return true
}

func (e TwoDimensionalArray) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
	return
}

func (e TwoDimensionalArray) Len() int {
	return len(e)
}

// dp[i] = Max(d[j]+1,dp[i]) ( 0<=j<i; envelopes[i][0]>envelopes[j][0]&&envelopes[i][1]>envelopes[j][1])
// dp[i] 在i位置的最长严格递增子序列的长度
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}
	es := TwoDimensionalArray(envelopes)
	sort.Sort(es)
	fmt.Printf("%v\n", es)

	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	max := 1
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if es[i][0] > es[j][0] && es[i][1] > es[j][1] {
				dp[i] = Max(dp[i], dp[j]+1)
				if max < dp[i] {
					max = dp[i]
				}
			}
		}
	}

	return max
}

func main() {
	testCases := [][][]int{
		{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
		{{10, 17}, {10, 19}, {16, 2}, {19, 18}, {5, 6}},
		{{1, 1}, {1, 1}},
	}
	for _, envelopes := range testCases {
		max := maxEnvelopes(envelopes)
		println(max)
	}
}
