package main

import (
	. "github.com/weedge/leetcode/define"
)

//找最底层的最左边的叶子节点值 从右至左，找左边的深度最大的
//广度优先 队列,每层按从左到右入队列，每层第一个最左边，最底层为最后一层第一个
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, node := range tmp {
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return res
}

//找最底层的最左边的叶子节点值 从右至左，找左边的深度最大的
//深度优先 递归
func findBottomLeftValueV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	dep := -1
	var dfs func(node *TreeNode, d int)
	dfs = func(node *TreeNode, d int) {
		if node == nil {
			return
		}
		if d >= dep {
			dep = d
			res = node.Val
		}

		/*
			d++
			dfs(node.Right, d)
			dfs(node.Left, d)
			d--//回溯
		*/
		dfs(node.Right, d+1) //隐藏回溯
		dfs(node.Left, d+1)  //隐藏回溯
		return
	}
	dfs(root, 0)

	return res
}

func main() {

}
