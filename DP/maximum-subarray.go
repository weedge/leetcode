package main

func maxSubArray(nums []int) (int, int) {
	max := nums[0]
	s, e := 0, 0
	for i := 1; i < len(nums); i++ {
		a := nums[i]
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			if a == nums[i] {
				s = i
			}
			e = i
			max = nums[i]
		}
	}
	return max, e - s + 1
}

// DP[i] = max(DP[i-1] + nums[i],nums[i])
// res = max(DP[])
// maxLen sub array max len
func maxSubArrayV2(nums []int) (res, maxLen int) {
	if len(nums) == 0 {
		return
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res = nums[0]
	s, e := 0, 0
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			if dp[i] == nums[i] {
				s = i
			}
			res = dp[i]
			e = i
		}
	}
	maxLen = e - s + 1

	return
}
func max(a, b int) (max int) {
	max = b
	if a > b {
		max = a
	}
	return
}

func main() {
	max, maxLen := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	max2, maxLen2 := maxSubArrayV2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	println(max, maxLen, max2, maxLen2)
}
