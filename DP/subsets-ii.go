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
	ans = append(ans, []int{})
	for i := 0; i < n; i++ {
		m := len(ans)
		for j := 0; j < m; j++ {
			tmp := make([]int, len(ans[j]))
			copy(tmp, ans[j])
			ans = append(ans, append(tmp, nums[i]))
		}
	}
	mapNode := map[string]struct{}{}
	res := [][]int{}
	for _, item := range ans {
		key := ""
		sort.Ints(item)
		for _, i := range item {
			key = fmt.Sprintf("%s%d", key, i)
		}
		if _, ok := mapNode[key]; !ok {
			res = append(res, item)
		}
		mapNode[key] = struct{}{}
	}

	return res
}

func main() {
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}
