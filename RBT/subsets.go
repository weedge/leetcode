package main

import (
	"fmt"
)

// 定义梳理递归树
// 从左至右遍历剩余集合 nums
// 从上往下遍历问题集合 set
// 从剩余集合中选择元素，判断是否可选取(剪枝)
// 叶子节点剩余节点为空为止
// 回溯递归
func subsetsBTFor(nums []int) (ans [][]int) {
	n := len(nums)
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		tmp := make([]int, len(set))
		copy(tmp, set)
		ans = append(ans, tmp)

		for i := cur; i < n; i++ {
			set = append(set, nums[i])
			dfs(i + 1)
			set = set[:len(set)-1] //回溯
		}
	}
	dfs(0)
	return
}

func subsetsBT(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
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
	return
}

func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

func getSubArrs(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	mapNums := map[int]struct{}{}
	for _, num := range nums {
		mapNums[num] = struct{}{}
	}

	n := len(mapNums)
	tNums := []int{}
	for item, _ := range mapNums {
		tNums = append(tNums, item)
	}

	res := [][]int{{}}
	for i := 1; i <= n; i++ { //长度
		res = append(res, getSubForK(tNums, i)...)
		//fmt.Println(res)
	}

	return res
}

// 暴力成回溯了。。。
// 回溯需要脑袋瓜子清晰的时候来写
func getSubForK(nums []int, k int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}
	ans := [][]int{}
	set := []int{}
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			if len(set) == k {
				tmp := make([]int, len(set))
				copy(tmp, set)
				ans = append(ans, tmp)
			}
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)

	return ans
}

// 测试集不能太大, O(n*2^n)
func main() {
	//nums1 := []int{4, 3, 2, 1}
	//nums1 := []int{3, 2, 4, 1}
	//fmt.Println(getSubForK(nums1, 3))
	//return
	testCases := [][]int{
		{1, 2, 3},
		{9, 0, 3, 5, 7},
	}
	for _, testCase := range testCases {
		//res := subsets(testCase)
		//res := subsetsBT(testCase)
		res := subsetsBTFor(testCase)
		fmt.Println(res)
	}
}
