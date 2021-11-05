package main

import "fmt"

func subsetsDP(nums []int) (ans [][]int) {
	n := len(nums)
	ans = [][]int{}
	ans = append(ans, []int{})
	for i := 0; i < n; i++ {
		m := len(ans)
		for j := 0; j < m; j++ {
			tmpArr := make([]int, len(ans[j]))
			copy(tmpArr, ans[j])
			ans = append(ans, append(tmpArr, nums[i]))
		}
	}
	return
}

func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}

func getSubArrs(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	mapNums := map[int]struct{}{}
	for _, num := range nums {
		mapNums[num] = struct{}{}
	}

	n := len(mapNums)
	tNums := []int{}
	for item, _ := range mapNums {
		tNums = append(tNums, item)
	}

	res := [][]int{{}}
	for i := 1; i <= n; i++ { //长度
		res = append(res, getSubForK(tNums, i)...)
		//fmt.Println(res)
	}

	return res
}

func getSubForK(nums []int, k int) [][]int {
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n; i++ {
		item := []int{nums[i]}
		if k == 1 {
			res = append(res, item)
			continue
		}
		for j, num := range nums {
			if j < i {
				continue
			}
			if num == nums[i] {
				continue
			}
			item = append(item, num)
			if len(item) == k {
				res = append(res, item)
				item = []int{nums[i]}
			}
		}
	}

	return res
}

func main() {
	//nums1 := []int{4, 3, 2, 1}
	//nums1 := []int{3, 2, 4, 1}
	//fmt.Println(getSubForK(nums1, 3))
	//return
	nums := []int{9, 0, 3, 5, 7}
	res := subsets(nums)
	fmt.Println(res)

	nums2 := []int{9, 0, 3, 5, 7}
	res2 := subsetsDP(nums2)
	fmt.Println(res2)
	//nums := [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}
}
