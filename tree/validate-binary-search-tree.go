package main

import (
	. "github.com/weedge/leetcode/define"
)

//中序遍历：递归，左中右， 和前一个节点的值比较，前一个节点的值大于当前节点的值 返回false
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	res := true
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != nil && node.Val <= pre.Val {
			res = false
			return
		}
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	return res
}

//中序遍历：栈，左中右， 和前一个节点的值比较，前一个节点的值大于当前节点的值 返回false
func isValidBSTv2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	s := []*TreeNode{}
	var pre *TreeNode
	for root != nil || len(s) > 0 {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		node := s[len(s)-1]
		s = s[:len(s)-1]
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		root = node.Right
	}

	return true
}

func isValidBSTV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := true

	var dfs func(node *TreeNode, low, high int)
	dfs = func(node *TreeNode, low, high int) {
		if node == nil {
			return
		}
		if node.Val <= low || node.Val >= high {
			res = false
			return
		}
		dfs(node.Left, low, node.Val)
		dfs(node.Right, node.Val, high)
	}
	dfs(root, -(1<<32 - 1), 1<<32-1)

	return res
}

func main() {

}
