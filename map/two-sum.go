package main

/*
map:
1. hash, un order map: O(1)  bad case all the same key -> link/array: O(n)
2. search Tree, order map, multi order map, eg: RBT(red black tree): O(logn)

//think more~
//reduce:
//multi merge sort (Parallel and Concurrent)
*/

func TwoSumWithIndex(nums []int, target int) (res [][2]int) {
	hashMapArr := map[int][]int{}

	for i, v := range nums {
		if indexes, ok := hashMapArr[target-v]; ok {
			for _, index := range indexes {
				res = append(res, [2]int{index, i})
			}
		}
		hashMapArr[v] = append(hashMapArr[v], i)
	}

	return res
}
