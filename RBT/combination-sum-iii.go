package main

import "fmt"

// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字
func combinationSum3(k int, n int) [][]int {
	nums := [9]int{}
	for i := 0; i < 9; i++ {
		nums[i] = i + 1
	}
	ans := [][]int{}
	set := []int{}
	var dfs func(cur, sum int)
	dfs = func(cur, sum int) {
		if len(set) == k && sum == n {
			tmp := make([]int, len(set))
			copy(tmp, set)
			ans = append(ans, tmp)
		}
		for i := cur; i < len(nums); i++ {
			set = append(set, nums[i])
			sum += nums[i]
			dfs(i+1, sum)
			set = set[:len(set)-1]
			sum -= nums[i]
		}
	}
	dfs(0, 0)
	return ans
}

// once pass nice~
func main() {
	testCases := [][2]int{
		{3, 7},
		{3, 9},
	}
	for _, testCase := range testCases {
		res := combinationSum3(testCase[0], testCase[1])
		fmt.Println(res)
	}
}
