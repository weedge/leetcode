package main

import (
	"fmt"
)

// 可重复使用，所以不会出现有0，负数和正数混合的情况,不然无解了
// 梳理递归分支
// dfs(cur, sum) //当前位置可重复使用计算, 排列问题，讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为不同列表时）
// dfs(i, sum) //当前位置可重复使用计算, 组合问题，不讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为同列表时）
// tips: 如果通过排序问题->组合问题 需要copy后去重, 防止对原来的集合路径排序后影响回溯
func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return [][]int{}
	}

	ans := [][]int{}
	set := []int{}
	//mapSet := map[string]struct{}{}
	var dfs func(cur int, sum int)
	dfs = func(cur int, sum int) {
		if target < sum {
			return
		}
		if sum == target {
			tmp := make([]int, len(set))
			copy(tmp, set)

			/*
				//sort tmp
				sort.Ints(tmp)
				key := ""
				for _, k := range tmp {
					key = fmt.Sprintf("%s%d", key, k)
				}
				if _, ok := mapSet[key]; ok {
					return
				}
				mapSet[key] = struct{}{}
			*/

			ans = append(ans, tmp)
			return
		}
		for i := cur; i < n; i++ {
			sum += candidates[i]
			set = append(set, candidates[i])
			//dfs(cur, sum) //当前位置可重复使用计算, 排列问题，讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为不同列表时）
			dfs(i, sum) //当前位置可重复使用计算, 组合问题，不讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为同列表时）
			// 回溯
			sum -= candidates[i]
			set = set[:len(set)-1]
		}
	}
	dfs(0, 0)

	return ans
}

func main() {
	type Input struct {
		nums   []int
		target int
	}
	testCases := []Input{
		{nums: []int{2, 3, 6, 7}, target: 7},
		{nums: []int{2, 3, 5}, target: 8},
		{nums: []int{2}, target: 1},
	}

	for _, testCase := range testCases {
		res := combinationSum(testCase.nums, testCase.target)
		fmt.Println(res)
	}
}
