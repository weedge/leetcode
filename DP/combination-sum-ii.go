package main

import (
	"fmt"
	"sort"
)

// 不重复
// subsets dp u know
// 全集ans -> filter =target -> res
func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)

	mapNode := map[string]struct{}{}
	res := [][]int{}
	ans := [][]int{{}}
	for i := 0; i < n; i++ {
		m := len(ans)
		for j := 0; j < m; j++ {
			tmp := make([]int, len(ans[j]))
			copy(tmp, ans[j])
			tmp = append(tmp, candidates[i])
			ans = append(ans, tmp)

			sum := 0
			key := ""
			for _, item := range tmp {
				sum += item
				key = fmt.Sprintf("%s%d", key, item)
			}
			if _, ok := mapNode[key]; !ok && sum == target {
				res = append(res, tmp)
			}
			mapNode[key] = struct{}{}
		}
	}

	return res
}

func main() {
	type Input struct {
		nums   []int
		target int
	}
	testCases := []Input{
		{nums: []int{10, 1, 2, 7, 6, 1, 5}, target: 8},
		{nums: []int{2, 5, 2, 1, 2}, target: 5},
		{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, target: 27},//超时
	}

	for _, testCase := range testCases {
		res := combinationSum2(testCase.nums, testCase.target)
		fmt.Println(res)
	}
}
