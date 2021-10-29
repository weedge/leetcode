package main

import (
	"fmt"
)

func getIslandCnV2(grid [][]int) int {
	num := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < len(grid) && i >= 0 && j < len(grid[0]) && j >= 0 && grid[i][j] == 1 {
			grid[i][j] = 0
			dfs(i, j+1)
			dfs(i, j-1)
			dfs(i-1, j)
			dfs(i+1, j)
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 {
				num++
				dfs(j, i)
			}
		}
	}
	return num
}

func getIslandCnV1(nums [][]int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	num := 0
	m := len(nums[0])
	mapTags := make(map[string]string, n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums[i][j] == 0 {
				continue
			}
			ikey := fmt.Sprintf("%d-%d", i, j)
			tag, ok := mapTags[ikey]
			if !ok {
				mapTags[ikey] = ikey
				tag = ikey
			}
			if i-1 > 0 && nums[i-1][j] == 1 {
				key := fmt.Sprintf("%d-%d", i-1, j)
				mapTags[key] = tag
			}
			if i+1 < n && nums[i+1][j] == 1 {
				key := fmt.Sprintf("%d-%d", i+1, j)
				mapTags[key] = tag
			}
			if j+1 < n && nums[i][j+1] == 1 {
				key := fmt.Sprintf("%d-%d", i, j+1)
				mapTags[key] = tag
			}
			if j-1 > 0 && nums[i][j-1] == 1 {
				key := fmt.Sprintf("%d-%d", i, j-1)
				mapTags[key] = tag
			}
		}
	}
	//fmt.Println(mapTags, len(mapTags))

	mapCn := make(map[string]int, n)
	for _, tag := range mapTags {
		mapCn[tag]++
	}
	num = len(mapCn)

	return num
}

func main() {
	testCases := [][][]int{
		{
			{0, 1},
			{1, 0},
		},
		{
			{1, 1, 1, 0},
			{1, 1, 0, 0},
			{1, 0, 0, 0},
		},
		{
			{1, 1, 1, 0},
			{1, 1, 0, 0},
			{1, 0, 1, 0},
			{1, 0, 0, 1},
		},
	}

	for _, testCase := range testCases {
		n1 := getIslandCnV1(testCase)
		println(n1)
	}
	for _, testCase := range testCases {
		n2 := getIslandCnV2(testCase)
		println(n2)
	}
}
