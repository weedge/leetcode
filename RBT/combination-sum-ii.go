package main

import (
	"fmt"
	"sort"
)

// 不重复
// 梳理递归分支
// dfs(i+1, sum) //i+1 不重复,当前位置可重复使用计算, 组合问题，不讲究顺序（即 [2, 2, 3] 与 [2, 3, 2] 视为同列表时）
// 排序后，加速剪枝，无需用map过滤
func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return [][]int{}
	}

	ans := [][]int{}
	set := []int{}
	//mapSet := map[string]struct{}{}
	var dfs func(cur, sum int)
	dfs = func(cur, sum int) {
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
			if i > cur && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			set = append(set, candidates[i])
			dfs(i+1, sum) //和组合总和的区别，这里是i+1，每个数字在每个组合中只能使用一次
			// 回溯
			sum -= candidates[i]
			set = set[:len(set)-1]
		}
	}
	//排序，相同元素在一起，方便后面剪枝 i>cur
	sort.Ints(candidates)
	dfs(0, 0)

	return ans
}

// 不重复
// subsets u know
// 排序问题 -> filter -> 组合问题
func combinationSum2ByAll(candidates []int, target int) [][]int {
	n := len(candidates)
	if n == 0 {
		return [][]int{}
	}

	ans := [][]int{}
	set := []int{}
	var dfs func(cur, sum int)
	dfs = func(cur, sum int) {
		if sum == target {
			tmp := make([]int, len(set))
			copy(tmp, set)
			ans = append(ans, tmp)
			return
		}
		if cur == n {
			return
		}
		sum += candidates[cur]
		set = append(set, candidates[cur])
		dfs(cur+1, sum)
		sum -= candidates[cur]
		set = set[:len(set)-1]
		dfs(cur+1, sum)
	}
	sort.Ints(candidates)
	dfs(0, 0)

	res := [][]int{}
	// filter
	mapNode := map[string]struct{}{}
	for _, item := range ans {
		key := ""
		for _, k := range item {
			key = fmt.Sprintf("%s%d", key, k)
		}
		if _, ok := mapNode[key]; !ok {
			res = append(res, item)
			mapNode[key] = struct{}{}
		}
	}

	return res
}

func testCost() {
	type Input struct {
		nums   []int
		target int
	}
	testCases := []Input{
		{nums: []int{-2}, target: -2},
		{nums: []int{0, -2}, target: -2},
		{nums: []int{0, -1, 2, -3}, target: -2},
		{nums: []int{0, -1, -2, -3}, target: -2},
		{nums: []int{10, 1, 2, 7, 6, 1, 5}, target: 8},
		{nums: []int{2, 5, 2, 1, 2}, target: 5},
		{nums: []int{1, 1}, target: 2},
		{nums: []int{2}, target: 2},
		{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, target: 27}, //测试超时
		{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, target: 30}, //测试超时
	}

	for _, testCase := range testCases {
		//println(len(testCase.nums))
		res := combinationSum2(testCase.nums, testCase.target)
		fmt.Println(res)
	}
}

func main() {
	testCost()
}
