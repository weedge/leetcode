package main

import "fmt"

func letterCasePermutation(s string) []string {
	n := len(s)
	if n == 0 {
		return []string{}
	}

	ans := []string{}
	set := make([]byte, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			//tmp := make([]byte, len(set))
			//copy(tmp, set)
			//ans = append(ans, string(tmp))
			ans = append(ans, string(set))
			return
		}
		if s[cur] >= '0' && s[cur] <= '9' {
			set[cur] = s[cur]
			dfs(cur + 1)
		} else if s[cur] >= 'a' && s[cur] <= 'z' { // 回溯 小写 变成 大写 再次搜索
			set[cur] = s[cur]
			dfs(cur + 1)
			set[cur] = s[cur] - 32
			dfs(cur + 1)
		} else if s[cur] >= 'A' && s[cur] <= 'Z' { // 回溯 大写 变成 小写 再次搜索
			set[cur] = s[cur]
			dfs(cur + 1)
			set[cur] = s[cur] + 32
			dfs(cur + 1)
		}
	}
	dfs(0)

	return ans
}

func main() {
	testCases := []string{
		"a1b2",
	}

	for _, testCase := range testCases {
		res := letterCasePermutation(testCase)
		fmt.Println(res)
	}

}
