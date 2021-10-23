package main

import (
	"fmt"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	m := len(matrix[0])
	nums := make([]int, m)
	max := -(1<<32 - 1)

	for i := 0; i < n; i++ {
		for a := 0; a < m; a++ {
			nums[a] = 0
		}
		for j := i; j < n; j++ {
			for k := 0; k < m; k++ {
				nums[k] += matrix[j][k]
			}
			// 每扫一行，累计至一维nums中计算最大子序列和，区间[s,e] 对应 二维空间的列;
			subMax, _, _ := MaxSubArrayV0ForK(nums, k)
			if subMax > max && subMax <= k {
				max = subMax
			}
		}
	}

	return max
}

// 该暴力还是要暴力, 理解万岁。。。
func MaxSubArrayV0ForK(tmp []int, k int) (int, int, int) {
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

func main() {
	matrix := [][]int{{2, 2, -1}}
	k := 3
	max := maxSumSubmatrix(matrix, k)
	println(max)

	nums := []int{2, 2, -1}
	l := 0
	res, s, e := MaxSubArrayV0ForK(nums, l)
	println(res, s, e)
	fmt.Printf("%v\n", nums[s:e+1])
}
