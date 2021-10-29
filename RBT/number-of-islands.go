package main

import (
	"fmt"
)

// 回溯
// 深度优先
// 以当等于1的点，开始递归回溯上下左右为1的点，变为0, 直到数组边界和中途遇到0的点
func getIslandCnV2(grid [][]int) int {
	num := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < len(grid) && i >= 0 && j < len(grid[0]) && j >= 0 && grid[i][j] == 1 {
			grid[i][j] = 0
			dfs(i-1, j)
			dfs(i+1, j)
			dfs(i, j-1)
			dfs(i, j+1)
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

// key在map中有对应的值，oldVal
// 把map中的值等于val的key,对应值全部替换成oldVal
// 近朱者赤近墨者黑
func changeSameKeyToSameVal(mapTags map[string]string, key, val string) {
	if oldVal, ok := mapTags[key]; ok {
		for k, v := range mapTags {
			if v == val {
				mapTags[k] = oldVal
			}
		}
	}
}

func TestChangeSameKeyToSameVal() {
	mapTags := map[string]string{"0-0": "0-0", "0-2": "0-2", "0-3": "0-2", "0-4": "0-2", "1-0": "0-0", "1-2": "0-2", "1-4": "0-2", "2-0": "0-0",
		"2-1": "0-2", "2-2": "0-2", "2-4": "0-2"}
	fmt.Println(mapTags)
	changeSameKeyToSameVal(mapTags, "2-1", "0-0")
	fmt.Println(mapTags)
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
			if j-1 > 0 && nums[i][j-1] == 1 {
				key := fmt.Sprintf("%d-%d", i, j-1)
				mapTags[key] = tag
			}
			if i+1 < n && nums[i+1][j] == 1 {
				key := fmt.Sprintf("%d-%d", i+1, j)
				if itag, ok := mapTags[key]; ok {
					mapTags[ikey] = itag
				} else {
					mapTags[key] = tag
				}
			}
			if j+1 < m && nums[i][j+1] == 1 {
				key := fmt.Sprintf("%d-%d", i, j+1)
				if itag, ok := mapTags[key]; ok {
					mapTags[ikey] = itag
				} else {
					mapTags[key] = tag
				}
			}
			changeSameKeyToSameVal(mapTags, ikey, tag)
			//fmt.Println(i, j, mapTags)
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
	//TestChangeSameKeyToSameVal()
	//return
	testCases := [][][]int{
		{
			{1, 0, 1, 1, 1},
			{1, 0, 1, 0, 1},
			{1, 1, 1, 0, 1},
		},
		{
			{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1},
		},
		{
			{1, 1},
		},
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

	println("---")

	for _, testCase := range testCases {
		n2 := getIslandCnV2(testCase)
		println(n2)
	}
}
