package main

import "fmt"

// subsets u know
// 递推 需要复用前面的状态
// 得到结果需要过滤下
func combine(n int, k int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if k == 0 {
		return [][]int{}
	}
	if k > n {
		return [][]int{}
	}
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	ans := [][]int{}
	ans = append(ans, []int{})
	for i := 0; i < n; i++ {
		m := len(ans)
		for j := 0; j < m; j++ {
			if len(ans[j]) > k-1 {
				continue
			}
			tmp := make([]int, len(ans[j]))
			copy(tmp, ans[j])
			ans = append(ans, append(tmp, nums[i]))
		}
	}

	//filter
	res := [][]int{}
	for _, item := range ans {
		if len(item) != k {
			continue
		}
		res = append(res, item)
	}

	return res
}

func main() {
	res := combine(3, 2)
	fmt.Println(res)
}
