package main

import (
	"fmt"
	"sort"
)

// 去重需要先对集合排序
func subsetsWithDupFor(nums []int) (ans [][]int) {
	n := len(nums)

	set := []int{}
	used := make([]bool, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		tmp := make([]int, len(set))
		copy(tmp, set)
		ans = append(ans, tmp)

		for i := cur; i < n; i++ {
			//!used[i-1] && nums[i] == nums[i-1]说明 同一层有相同的，不可重复选, 已经处理过了
			//used[i-1] && nums[i] == nums[i-1]说明 同一树枝上就可以重复取，因为同一树枝上元素的集合才是唯一子集！
			if i > 0 && !used[i-1] && nums[i] == nums[i-1] {
				continue
			}
			used[i] = true
			set = append(set, nums[i])
			dfs(i + 1)
			set = set[:len(set)-1] //回溯
			used[i] = false
		}
	}
	sort.Ints(nums)
	dfs(0)

	return
}

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}
	ans := [][]int{}
	set := []int{}
	mapNode := map[string]struct{}{}
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {

			//选择
			tmp := make([]int, len(set))
			copy(tmp, set)
			ans = append(ans, tmp)
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1] //回溯
		dfs(cur + 1)
	}
	dfs(0)

	res := [][]int{}
	for _, item := range ans {
		sort.Ints(item)
		key := ""
		for _, k := range item {
			key = fmt.Sprintf("%s%d", key, k)
		}
		if _, ok := mapNode[key]; ok {
			continue
		}
		mapNode[key] = struct{}{}
		res = append(res, item)
	}

	return res
}

func main() {
	//nums := []int{1, 2, 2}
	testCases := [][]int{
		{0},
		{1, 2, 2},
		{4, 1, 0},
		{4, 4, 4, 1, 4},//test sort for filter
	}
	for _, testCase := range testCases {
		//res := subsetsWithDup(testCase)
		res := subsetsWithDupFor(testCase)
		fmt.Println(res)
	}
}
