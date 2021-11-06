package main

import (
	"fmt"
	"sort"
)

//有重复的数，给出不重复全排列
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}

	ans := [][]int{}
	set := []int{}
	used := make([]bool, n)
	//mapNode := map[string]struct{}{}

	var dfs func()
	dfs = func() {
		if len(set) == n {
			tmp := make([]int, len(set))
			copy(tmp, set)

			/*
				key := ""
				for _, k := range tmp {
					key = fmt.Sprintf("%s%d", key, k)
				}
				if _, ok := mapNode[key]; ok {
					return
				}
				mapNode[key] = struct{}{}
			*/
			ans = append(ans, tmp)

			return
		}

		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}

			if i > 0 && !used[i-1] && nums[i] == nums[i-1] {
				continue
			}

			used[i] = true
			set = append(set, nums[i])
			dfs()
			set = set[:len(set)-1]
			used[i] = false
		}

		return
	}
	// 需要排序，才能剪掉相近相同的
	sort.Ints(nums)
	dfs()

	return ans
}

func main() {
	testCases := [][]int{
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
