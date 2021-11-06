package main

import "fmt"

//无重复的数，给出全排列
// nums = [1,2,3]
// dp[0] = [[1]]
// dp[1]需要将2插入到[1]中，即dp[1] = [2,1] [1,2];
// dp[2]需要将3插入到dp[1]中，即dp[2] = [3,2,1] [2,3,1] [2,1,3] [3,1,2] [1,3,2] [1,2,3]
// (前三个是将3插入到[2,1]中，后三个是将3插入到[1,2]中)
// dp[k] = 将 nums[k] 分别插入dp[k-1]集合不同排序数组中的不同位置
func permute(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}

	ans := [][]int{{nums[0]}}
	for i := 1; i < n; i++ {
		m := len(ans)
		res := [][]int{}
		for j := 0; j < m; j++ {
			tmpArr := insert(ans[j], nums[i])
			res = append(res, tmpArr...)
		}
		ans = res
	}

	return ans
}

// insert [1,2],3 => [[3,1,2],[1,3,2],[1,2]]
// insert [1,1],1 => [[1,1,1],[1,1,1],[1,1,1]]
func insert(arr []int, val int) [][]int {
	res := [][]int{}
	for pos := 0; pos < len(arr)+1; pos++ {
		tmp := []int{}
		for i, item := range arr {
			if i == pos {
				tmp = append(tmp, val)
			}
			tmp = append(tmp, item)
		}
		if pos == len(arr) {
			tmp = append(tmp, val)
		}
		res = append(res, tmp)
	}

	return res
}

func main() {
	//fmt.Println(insert([]int{1, 2}, 3))
	//fmt.Println(insert([]int{1, 1}, 1))
	//return
	testCases := [][]int{
		{1, 2, 3},
	}

	for _, testCase := range testCases {
		res := permute(testCase)
		fmt.Println(res)
	}
}
