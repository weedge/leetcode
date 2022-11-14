package main

func peakIndexInMountainArray(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	l, h := 0, len(arr)-1
	if arr[l] > arr[l+1] {
		return l
	}
	if arr[h] > arr[h-1] {
		return h
	}

	for l <= h {
		m := l + (h-l)>>1
		if m == 0 && arr[m] < arr[m+1] {
			return m + 1
		}
		if m == len(arr)-1 && arr[m-1] > arr[m] {
			return m - 1
		}
		if arr[m] > arr[m-1] && arr[m] > arr[m+1] {
			return m
		} else if arr[m] > arr[m-1] {
			l = m + 1
		} else if arr[m] < arr[m-1] {
			h = m - 1
		}
	}
	return l
}
