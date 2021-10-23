package main

import "fmt"

func maxSubArrayV0(tmp []int) (int, int, int) {
	n := len(tmp)
	if n == 0 {
		return 0, 0, 0
	}
	max := -(1<<32 - 1)
	s, e := 0, 0

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += tmp[j]
			if sum > max {
				max = sum
				s = i
				e = j
			}
		}
	}

	return max, s, e
}

// 该暴力还是要暴力, 理解万岁。。。
func maxSubArrayV0ForK(tmp []int, k int) (int, int, int) {
	n := len(tmp)
	if n == 0 {
		return 0, 0, 0
	}
	max := -(1<<32 - 1)
	s, e := 0, 0

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += tmp[j]
			if sum > max && sum <= k {
				max = sum
				s = i
				e = j
			}
		}
	}

	return max, s, e
}

func maxSubArrayV1(tmp []int) (int, int, int) {
	nums := make([]int, len(tmp))
	copy(nums, tmp)
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
	return max, s, e
}

// DP[i] = max(DP[i-1] + nums[i],nums[i])
// res = max(DP[])
// maxLen sub array max len
func maxSubArrayForLong(nums []int) (res, s, e int) {
	if len(nums) == 0 {
		return
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res = nums[0]
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

	return
}
func max(a, b int) (max int) {
	max = b
	if a > b {
		max = a
	}
	return
}

func maxSubArrayForShort(nums []int) (res, s, e int) {
	start := 0
	n := len(nums)
	sum := -(1<<32 - 1)
	res = -(1<<32 - 1)
	for i := 0; i < n; i++ {
		if sum < 0 {
			start = i
			sum = nums[i]
		} else {
			sum = sum + nums[i]
		}
		if res < sum {
			res = sum
			s = start
			e = i
		}
	}

	return
}

func main() {
	/*
		nums := []int{-2, 1, 6, -3, 4, -1, 2, 1, -5, 4, 7}
		fmt.Printf("%v\n", nums)
		max, s, e := maxSubArrayV1(nums)
		println(max, s, e)
		fmt.Printf("%v\n", nums[s:e+1])
	*/

	nums := []int{450, 309, 460, -155, 337, -335, 45, -220, -209, 352, -348, -448, 412, 37, 155, -283, 493, -500, -116, -260, -443, 40, -235, 400, 468, 400, 446, 186, -84, 196, -363, -252, -193, -110, -26, -124, 116, -210, -374, -69, 142, 137, 316, 136, -73, 393, 388, -188, 37, 41, 286, -93, -107, 80, -125, -486, 119, 197, -333, 409}

	max0, s0, e0 := maxSubArrayV0(nums)
	println(max0, s0, e0)
	fmt.Printf("%v\n", nums[s0:e0+1])

	max1, s1, e1 := maxSubArrayV1(nums)
	println(max1, s1, e1)
	fmt.Printf("%v\n", nums[s1:e1+1])

	max2, s2, e2 := maxSubArrayForLong(nums)
	println(max2, s2, e2)
	fmt.Printf("%v\n", nums[s2:e2+1])

	max3, s3, e3 := maxSubArrayForShort(nums)
	println(max3, s3, e3)
	fmt.Printf("%v\n", nums[s3:e3+1])

}
