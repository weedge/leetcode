package main

func findLengthOfLCISV1(nums []int) int {
	max := 0
	s := 0
	for i, v := range nums {
		// un continue s = i
		if i > 0 && v <= nums[i-1] {
			s = i
		}
		if max < (i - s + 1) {
			max = i - s + 1
		}
	}

	return max
}

// dp[i+1] = dp[i] + 1 (nums[i+1]>nums[i])
// dp[0] = 0
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	max := 1
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 0; i < n-1; i++ {
		if nums[i+1] > nums[i] {
			dp[i+1] = dp[i] + 1
		}
		if max < dp[i+1] {
			max = dp[i+1]
		}
	}

	return max
}

func main() {
	nums := []int{1}
	max := findLengthOfLCIS(nums)
	println(max)
}
