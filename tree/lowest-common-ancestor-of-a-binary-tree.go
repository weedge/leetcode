package main

import (
	. "github.com/weedge/leetcode/define"
)

//公共父节点
//后序遍历 回溯
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	//一路出栈返回（理解）
	if right == nil {
		return left
	}

	return right
}

func main() {

}
