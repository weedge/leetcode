package main

import "fmt"

// subsets u know
func combine(n int, k int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if k == 0 {
		return [][]int{}
	}
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
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

func main() {
	res := combine(3, 2)
	fmt.Println(res)
}
