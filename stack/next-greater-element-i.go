package main

// 暴力，O(n*m)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	if n == 0 {
		return []int{}
	}
	m := len(nums2)

	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = -1
		for j, val := range nums2 {
			if val == nums1[i] {
				for k := j; k < m; k++ {
					if nums2[k] > val {
						result[i] = nums2[k]
						break
					}
				}
				break
			}
		}
	}

	return result
}

// 单调栈 + map，O(M+N)
func nextGreaterElementV1(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	if n == 0 {
		return []int{}
	}

	m := len(nums2)
	stack := []int{}
	mapNext := make(map[int]int, m)
	for j := m - 1; j >= 0; j-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums2[j] {
			stack = stack[:len(stack)-1]
		}
		mapNext[nums2[j]] = -1
		if len(stack) > 0 {
			mapNext[nums2[j]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[j])
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
		if val, ok := mapNext[nums1[i]]; ok {
			result[i] = val
		}
	}

	return result
}
