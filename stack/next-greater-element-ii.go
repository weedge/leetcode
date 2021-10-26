package main

// double nums
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	stack := []int{}
	result := make([]int, n)
	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}

		result[i%n] = -1
		if len(stack) > 0 {
			result[i%n] = stack[len(stack)-1]
		}

		stack = append(stack, nums[i%n])
	}

	return result
}
