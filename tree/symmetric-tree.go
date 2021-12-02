package main

import (
	. "github.com/weedge/leetcode/define"
)

//判断是否是镜像对称
//深度优先 递归 竖着比较
//分成两棵树比较,按照需求条件进行比较是否满足镜像对称条件
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var dfs func(node1, node2 *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}

		return dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
	}

	return dfs(root, root)
}

//判断是否是镜像对称
//广度搜索，队列 横着比较
//这里不是单个入队列方式，可以认为是两棵树的比较
//队列以两个比较的节点为一组入队列，出队列比较
func isSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return false
	}
	q := []*TreeNode{root, root}
	for len(q) > 0 {
		//2个比较节点为一组出队列比较
		n1, n2 := q[0], q[1]
		q = q[2:]

		// 判断
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}

		//队列以两个比较的节点为一组入队列，出队列比较
		q = append(q, n1.Left)
		q = append(q, n2.Right)

		q = append(q, n1.Right)
		q = append(q, n2.Left)
	}

	return true
}

func main() {

}
