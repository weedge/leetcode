package main

import (
	"fmt"
	"sort"
)

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
func getSumCombination(nums []int, target int) (res [][]int) {
	sort.Ints(nums)

	dfs(nums, target, 0, []int{}, &res)

	return
}

func getSumCombinationV2(nums []int, target int) (res [][]int) {
	sort.Ints(nums)

	count(nums, target, []int{}, &res)

	return
}

func count(nums []int, sum int, arr []int, res *[][]int) {
	if len(nums) == 0 {
		return
	}
	if nums[0] > sum {
		return
	}

	resArr := make([]int, len(arr))
	copy(resArr, arr)
	/*
		resArr := arr //slice not deep copy, append slice size new obj
	*/

	find := false
	for _, num := range nums {
		if num == sum {
			find = true
			break
		}
	}

	if find {
		resArr = append(resArr, sum)
		*res = append(*res, resArr)
		//fmt.Println(resArr)
		resArr = resArr[0 : len(resArr)-1]
		//fmt.Println(resArr)
		return
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < sum {
			resArr = append(resArr, nums[i])
			count(nums[0:i], sum-nums[i], resArr, res)
			resArr = resArr[0 : len(resArr)-1]
		}
	}

	return
}

func dfs(nums []int, target, begin int, _arr []int, res *[][]int) {
	var arr []int
	for i := range _arr {
		arr = append(arr, _arr[i])
	}

	if target == 0 {
		*res = append(*res, arr)
		return
	}

	for i := begin; i < len(nums); i++ {
		if target-nums[i] < 0 {
			return
		}
		arr = append(arr, nums[i])
		dfs(nums, target-nums[i], i+1, arr, res)

		arr = arr[:len(arr)-1]

		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}

	}
}

func testGetSumCombination() {
	arrCases := map[int][][]int{
		//30: {{1, 2, 3, 4, 5, 6, 10, 9, 7, 11, 23}, {-1, 2, 3, 4, 5, 6, 10, 9, 7, 11, 23}},
		//-20: {{-1, -2, -7, -10}},
		//0:   {{1, 2, 3, 4, 5, 6, 10, 9, 7, 11, 23}, {-1, 2, 3, 4, 5, 6, 10, 9, 7, 11, 23}},
		//-1:  {{1, 2, 3, 4, 5, 6, 10, 9, 7, 11, 23}, {-1, 2, 3, 4, 5, 6, 10, 9, 7, 11, 23}},
		//100: {{1, 2, 3, 4, 5, 6, 10, 9, 7, 11, 23}, {-1, 2, 3, 4, 5, 6, 10, 9, 7, 11, 23}},
		//1:   {{}, {-1}},
		//1:  {{1}, {0, -1, 2}},
		-2: {{-2}, {-2, -3, 3}, {-1, 1, -2}, {0, -1, 2, -3}},
	}

	for target, testCases := range arrCases {
		for _, testCase := range testCases {
			fmt.Printf("target: %d \t testCase:%v\n", target, testCase)
			res := getSumCombination(testCase, target)
			//res := getSumCombinationV2(testCase, target)
			fmt.Printf("target: %d \t res:%v\n", target, res)
			for _, item := range res {
				sum := 0
				for _, v := range item {
					sum += v
				}
				if sum == target {
					fmt.Printf("item:%v\ttarget:%d \t pass\n", item, target)
				} else {
					fmt.Printf("item:%v\ttarget:%d \t unpass!!\n", item, target)
				}
			}
		}
	}

}

func main() {
	testGetSumCombination()
}
