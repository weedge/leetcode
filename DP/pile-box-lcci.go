package main

import (
	"fmt"
	"sort"

	"github.com/weedge/leetcode/utils"
)

// sort -> dp
// dp[i] = Max(dp[j]+1,dp[i])
// d[i] = Max(d[j]+box[i][2],d[i])
// ( 0<=j<i; box[i][0]>box[j][0]&&box[i][1]>box[j][1]&&box[i][2]>box[j][2])
// dp[i] 在i位置满足叠加盒子最大盒子数目
// d[i] 在i位置满足叠加盒子最大盒子累计深度
// 举一反三：可以衍生出相应的属性集合，求最大连续子问题
func pileBox(box [][]int) (int, int) {
	n := len(box)
	if n == 0 {
		return 0, 0
	}

	bx := utils.TwoDimensionalArray(box)
	sort.Sort(bx)
	fmt.Println(bx)

	dp := make([]int, n)
	d := make([]int, n)
	max := 1
	maxD := bx[0][2]
	for i := 0; i < n; i++ {
		dp[i] = 1
		d[i] = bx[i][2]
		for j := 0; j < i; j++ {
			if bx[i][0] > bx[j][0] && bx[i][1] > bx[j][1] && bx[i][2] > bx[j][2] {
				d[i] = utils.Max(d[i], d[j]+bx[i][2])
				dp[i] = utils.Max(dp[i], dp[i]+1)
			}
			if dp[i] > max {
				max = dp[i]
			}
			if d[i] > maxD {
				maxD = d[i]
			}
		}
	}

	return max, maxD
}

func main() {
	testCases := [][][]int{
		{{1, 1, 1}, {2, 3, 4}, {2, 6, 7}, {3, 4, 5}},
		{{1, 1, 1}},
		{{1, 1, 1}, {1, 1, 1}},
	}
	for _, box := range testCases {
		max, maxD := pileBox(box)
		println(max, maxD)
	}
}
