package main

import (
	"sort"
)

/*
Tips:

	just for one peak of multi peak;
	un support to get max peak from multi peak;
*/

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

func peakIndexInMountainArrayV1(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	if len(arr) == 2 && arr[0] > arr[1] {
		return 0
	}
	if len(arr) == 2 && arr[1] > arr[0] {
		return 1
	}

	pos := 0
	l, h := 1, len(arr)-2
	for l <= h {
		m := l + (h-l)>>1
		if arr[m] > arr[m-1] && arr[m] > arr[m+1] {
			return m
		} else if arr[m] < arr[m-1] && arr[m] > arr[m+1] {
			h = m - 1
			pos = h
		} else if arr[m] > arr[m-1] && arr[m] < arr[m+1] {
			l = m + 1
			pos = l
		}
	}
	return pos
}

func peakIndexInMountainArrayV2(arr []int) int {
	l, h := 0, len(arr)-1
	for l < h {
		m := l + (h-l)>>1
		if arr[m] > arr[m+1] {
			h = m
		} else {
			l = m + 1
		}
	}

	return l
}

// see search_test.go and doc
func peakIndexInMountainArrayV3(arr []int) int {
	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
}
