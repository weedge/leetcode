package main

import "fmt"

//无重复的数，给出全排列
func permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}
	ans := [][]int{}
	set := []int{}
	used := make([]bool, n) // 元素是否选择

	var dfs func()
	dfs = func() {// 无需 开始位置，需要从0开始
		if len(set) == n {// 集合大小等于长度返回
			tmp := make([]int, len(set))
			copy(tmp, set)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			set = append(set, nums[i])
			dfs()
			//回溯 状态
			set = set[:len(set)-1]
			used[i] = false
		}
	}
	dfs()

	return ans
}

func main() {
	testCases := [][]int{
		{1, 2, 3},
	}

	for _, testCase := range testCases {
		res := permute(testCase)
		fmt.Println(res)
	}
}
