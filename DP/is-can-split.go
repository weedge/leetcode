package main


// 请在脑壳清晰的时候推导
// 如果用暴力去解答，还是会遇到怎么组合的问题，这种属于直男思维，可以换个思维反过来递推，进行状态转移，就是dp思想了
// 第i次的集合 为 第i-1次集合 和 nums[i] 的组合, 复用思想
// nums = [1,2,3]
// dp[0] = [[]]
// dp[1] = [[],[1]]
// dp[2] = [[],[1],[2],[1,2]]
// dp[3] = [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// tips: 上一次(i-1)的集合需要另开一段空间copy 进行操作， 不然会覆盖以前的集合
func isCanSplit(nums []int) bool {
	n := len(nums)
	ans := [][]int{}
	ans = append(ans, []int{})
	for i := 0; i < n; i++ {
		m := len(ans)
		for j := 0; j < m; j++ {
			tmpArr := make([]int, len(ans[j]))
			copy(tmpArr, ans[j])
			ans = append(ans, append(tmpArr, nums[i]))
		}
	}
	//以上是求子集, 然后折叠求和比较

	sum := func(nums []int) int {
		res := 0
		for _, num := range nums {
			res += num
		}
		return res
	}

	l := len(ans)
	for k := 0; k < l/2; k++ {
		if sum(ans[k]) == sum(ans[l-k-1]) {
			return true
		}
	}

	return false
}


func main() {
	
}
