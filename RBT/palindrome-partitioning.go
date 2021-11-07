package main

import "fmt"

func partition(s string) [][]string {
	n := len(s)
	if n == 0 {
		return [][]string{}
	}
	ans := [][]string{}
	set := []string{}

	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			tmp := make([]string, len(set))
			copy(tmp, set)
			ans = append(ans, tmp)
		}
		for i := cur; i < n; i++ {
			if !checkPalindrome(s[cur : i+1]) {
				continue
			}
			set = append(set, s[cur:i+1])
			//dfs(cur + 1)
			dfs(i + 1)
			set = set[:len(set)-1]
		}

	}
	dfs(0)

	return ans
}

func checkPalindrome(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	l, r := 0, n-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func main() {
	testCases := []string{
		"aab",
		"abb",
		"baa",
		"aa",
	}

	for _, testCase := range testCases {
		res := partition(testCase)
		fmt.Println(res)
		return
	}
}
