package main

import (
	. "github.com/weedge/leetcode/define"
)

//中序遍历 计算
func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	var maxCount, cn, base int
	update := func(val int) {
		if val == base {
			cn++
		} else {
			base, cn = val, 1
		}

		if cn == maxCount {
			res = append(res, base)
		} else if cn > maxCount {
			maxCount = cn
			res = []int{base}
		}
	}

	s := []*TreeNode{}
	for len(s) > 0 || root != nil {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		node := s[len(s)-1]
		s = s[:len(s)-1]

		update(node.Val)

		root = node.Right
	}

	return res
}

//递归中序遍历 计算
func findModeV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	var maxCount, cn, base int
	update := func(val int) {
		if val == base {
			cn++
		} else {
			base, cn = val, 1
		}

		if cn == maxCount {
			res = append(res, base)
		} else if cn > maxCount {
			maxCount = cn
			res = []int{base}
		}
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)

	return res
}

func main() {

}
