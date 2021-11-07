package main

import "fmt"

// 递增子序列 (子序列至少2个元素,可相等)
// 不能排序去重

// 定义梳理递归树
// 从上往下遍历问题集合(递增子序列) set
// 从左至右遍历剩余集合 nums
// 从剩余集合中选择元素，判断是否可选取(剪枝):
// 1.同一父节点同层，不能选已经选取过的元素
// 2.选取的元素必须大于递增子序列问题集合最后一个元素
// 叶子节点剩余节点为空为止
// 回溯递归
func findSubsequences(nums []int) (ans [][]int) {
	n := len(nums)
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if len(set) > 1 { //递增子序列至少2个元素
			tmp := make([]int, len(set))
			copy(tmp, set)
			ans = append(ans, tmp)
		}

		mapUsedSet := map[int]struct{}{}
		for i := cur; i < n; i++ {
			// 选取的元素必须大于递增子序列问题集合最后一个元素
			if len(set) > 0 && nums[i] < set[len(set)-1] {
				continue
			}
			// 同一父节点同层，不能选已经选取过的元素
			if _, ok := mapUsedSet[nums[i]]; ok {
				continue
			}
			mapUsedSet[nums[i]] = struct{}{}
			set = append(set, nums[i])
			dfs(i + 1)
			set = set[:len(set)-1] //回溯
		}
	}
	dfs(0)

	return
}

func main() {
	testCases := [][]int{
		{4, 6, 7, 7},
		{4, 7, 4, 6},
		{4, 4, 3, 2, 1},
	}
	for _, testCase := range testCases {
		res := findSubsequences(testCase)
		fmt.Println(res)
	}
}
