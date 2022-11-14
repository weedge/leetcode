package main

func search(arr []int, target int) (pos int) {
	n := len(arr)
	l, h := 0, n-1
	for l <= h {
		m := l + (h-l)>>1
		if arr[m] == target {
			return m
		}
		if arr[0] <= arr[m] { //左边有序
			if target >= arr[0] && target < arr[m] {
				h = m - 1
			} else {
				l = m + 1
			}
		}
		if arr[n-1] >= arr[m] { //右边有序
			if target > arr[m] && target <= arr[n-1] {
				l = m + 1
			} else {
				h = m - 1
			}
		}
	}

	return -1
}
