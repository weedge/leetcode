package main

// 请在脑壳清晰的时候推导
// 推导:
// 如果用暴力去解答，还是会遇到怎么组合的问题，这种属于直男思维，可以换个思维反过来递推，进行状态转移，就是dp思想了
// 第i次的集合 为 第i-1次集合 和 nums[i] 的组合, 复用思想
// tips: 上一次(i-1)的集合需要另开一段空间copy 进行操作， 不然会覆盖以前的集合
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	ans = [][]int{}
	ans = append(ans, []int{})
	for i := 0; i < n; i++ {
		m := len(ans)
		for j := 0; j < m; j++ {
			tmpArr := make([]int, len(ans[j]))
			copy(tmpArr, ans[j])
			ans = append(ans, append(tmpArr, nums[i]))
		}
	}
	return
}

// 测试集不能太大, O(n*2^n)
func main() {

}
