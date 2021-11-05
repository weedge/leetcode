package main

import (
	"fmt"
	"sort"
)

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
	nums := []int{4, 1, 0}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}
