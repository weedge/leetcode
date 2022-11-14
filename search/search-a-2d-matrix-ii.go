package main

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
