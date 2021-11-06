package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w *string) {
	s := strings.Split(*w, "")
	sort.Strings(s)
	*w = strings.Join(s, "")
	return
}

func permutation(s string) []string {
	n := len(s)
	if n == 0 {
		return []string{}
	}
	ans := []string{}
	set := []byte{}
	used := make([]bool, n)

	var dfs func()
	dfs = func() {
		if len(set) == n {
			//tmp:=make([]byte,len(set))
			//copy(tmp,set)
			//ans = append(ans,string(tmp))
			ans = append(ans, string(set))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && !used[i-1] && s[i] == s[i-1] {
				continue
			}

			used[i] = true
			set = append(set, s[i])
			dfs()
			set = set[:len(set)-1]
			used[i] = false
		}
	}
	SortString(&s)
	dfs()

	return ans
}

func main() {
	testCases := []string{
		"abb",
		"baa",
		"aa",
	}

	for _, testCase := range testCases {
		res := permutation(testCase)
		fmt.Println(res)
	}

}
