package main

import "fmt"

var mapSet map[string]struct{}

// 有重复的数，给出不重复全排列
// dp permutations + mapSet filter
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}
	mapSet = map[string]struct{}{}

	ans := [][]int{{nums[0]}}
	for i := 1; i < n; i++ {
		m := len(ans)
		res := [][]int{}
		for j := 0; j < m; j++ {
			tmpArr := insertUnique(ans[j], nums[i])
			res = append(res, tmpArr...)
		}
		ans = res
	}

	return ans
}

// insert [1,2],3 => [[3,1,2],[1,3,2],[1,2]]
// insert [1,1],1 => [[1,1,1]]
// insert [1,2,2],1 => [[1,1,2,2],[1,2,1,2],[1,2,2,1]]
func insertUnique(arr []int, val int) [][]int {
	res := [][]int{}
	for pos := 0; pos < len(arr)+1; pos++ {
		tmp := []int{}
		for i, item := range arr {
			if i == pos {
				tmp = append(tmp, val)
			}
			tmp = append(tmp, item)
		}
		if pos == len(arr) {
			tmp = append(tmp, val)
		}

		key := ""
		for _, k := range tmp {
			key = fmt.Sprintf("%s%d", key, k)
		}
		if _, ok := mapSet[key]; ok {
			continue
		}
		mapSet[key] = struct{}{}

		res = append(res, tmp)
	}

	return res
}

func main() {
	//fmt.Println(insertUnique([]int{1, 2}, 3))
	//fmt.Println(insertUnique([]int{1, 1}, 1))
	//fmt.Println(insertUnique([]int{1, 2, 2}, 1))
	//return
	testCases := [][]int{
		{2, 2, 1, 1},
		{1},
		{1, 1},
		{1, 1, 2},
		{1, 2, 3},
		{3, 2, 1},
		{3, 2, 2},
		{3, 2, 3},
	}

	for _, testCase := range testCases {
		res := permuteUnique(testCase)
		fmt.Println(res)
	}
}
