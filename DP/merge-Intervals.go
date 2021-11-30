package main

import (
	"fmt"
	"sort"
)

// 直男思维(greed)：从局部解 => 下一步局部解 => 解
// 合并区间: 左边排序
// [[2,4],[5,7],[8,11]] [12,13]  => [[2,4],[5,7],[8,11],[12,13]]
// [[2,4],[5,7],[8,11]] [9,10]  => [[2,4],[5,7],[8,11]]
// [[2,4],[5,7],[8,11]] [10,12]  => [[2,4],[5,7],[8,12]]
func mergeIntervalsV2(arr [][2]int) [][2]int {
	n := len(arr)
	if n == 0 {
		return [][2]int{}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] <= arr[j][0]
	})
	//fmt.Println(arr)

	res := [][2]int{arr[0]}
	for i := 1; i < n; i++ {
		m := len(res)
		if res[m-1][1] < arr[i][0] {
			res = append(res, arr[i])
		} else {
			tmp := [2]int{res[m-1][0], res[m-1][1]}
			if res[m-1][1] < arr[i][1] {
				tmp = [2]int{res[m-1][0], arr[i][1]}
			}
			res[m-1] = tmp
		}
	}

	return res
}

//dp[i]{l,h} = dp[i-1]{min(l,arr[i][0]),max(h,arr[i][1])} (arr[i][0]<h)
//l = dp[i][0][0]
//h = dp[i][len(dp[i])-1][1]
//dp[i]{l,h} = dp[i-1]{l,h},{arr[i][0],arr[i][1]} (arr[i][0]>h)
// 时间复杂度：O(n*log(n))
// 合并重复区间:(右边排序)
// [[2,4],[5,7],[8,9]] [6,10] => [[2,4],[5,10]]
// [[2,4],[5,7],[8,9]] [10,11] => [[2,4],[5,7],[8,9],[10,11]]
func mergeIntervals(arr [][2]int) [][2]int {
	n := len(arr)
	if n == 0 {
		return [][2]int{}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] <= arr[j][1]
	})
	//fmt.Println(arr)

	res := [][2]int{arr[0]}
	for i := 1; i < n; i++ {
		m := len(res)
		l := 0
		h := m - 1
		if arr[i][0] <= res[h][1] {
			// search max less
			/*
				for j := 0; j < m; j++ {
					if res[j][0] <= arr[i][0] {
						l = j
					} else {
						break
					}
				}
			*/

			pos := searchMaxLess(res, arr[i][0])
			if pos > 0 {
				l = pos
			}

			tmp := [2]int{res[l][0], arr[i][1]}
			if arr[i][0] < res[l][0] {
				tmp = [2]int{arr[i][0], arr[i][1]}
			}
			res = res[:l]
			res = append(res, tmp)
		} else {
			res = append(res, arr[i])
		}
		//fmt.Println(res)
	}

	return res
}

//O(log(n))
func searchMaxLess(arr [][2]int, target int) int {
	pos := -1
	l := 0
	h := len(arr) - 1
	for l <= h {
		m := l + (h-l)/2
		if arr[m][0] > target {
			h = m - 1
		} else if arr[m][0] < target {
			l = m + 1
		} else {
			return m
		}
	}
	pos = l - 1
	return pos
}

func testSearchLess() {
	type testCase struct {
		arr    [][2]int
		target int
	}
	testCases := []testCase{
		{[][2]int{{1}, {3}, {7}, {9}, {11}}, 0},
	}

	for _, testCase := range testCases {
		pos := searchMaxLess(testCase.arr, testCase.target)
		fmt.Println(testCase.arr, testCase.target, pos)
	}
}

func main() {
	//testSearchLess()
	//return
	testCases := [][][2]int{
		{{4, 6}, {2, 5}, {8, 9}},
		{{4, 6}, {2, 5}, {8, 9}, {1, 7}},
		{{4, 6}, {2, 5}, {8, 9}, {0, 10}},
		{{2, 4}, {5, 7}, {8, 9}, {6, 10}},
		{{2, 4}, {5, 7}, {8, 9}, {0, 10}},
		{{2, 4}, {5, 7}, {1, 8}, {8, 9}},
	}

	for _, testCase := range testCases {
		//res := mergeIntervals(testCase)
		res := mergeIntervalsV2(testCase)
		fmt.Println(testCase)
		fmt.Println(res)
	}
}
