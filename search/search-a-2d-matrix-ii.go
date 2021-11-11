package main

import "fmt"

// 因为行，列都是有序的
// 从右上角开始查找, O(M+N)
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	x, y := 0, m-1
	for x < n && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else if matrix[x][y] < target {
			x++
		}
	}
	return false
}

func main() {
	type TestCase struct {
		nums   [][]int
		target int
	}
	testCases := []TestCase{
		{
			nums:   [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target: 5,
		},
	}

	for _, testCase := range testCases {
		res := searchMatrix(testCase.nums, testCase.target)
		fmt.Println(res)
	}
}
