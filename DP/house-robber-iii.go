package main

import (
	"github.com/weedge/leetcode/define"
	"github.com/weedge/lib/mathx"
)

func robIII(root *define.TreeNode) int {
	selected, unSelected := dfs(root)
	return mathx.Max(selected, unSelected)
}

// 树形🌲 DP
// 深度优先搜索，后续遍历(递归退栈从下往上算) + 当前节点位置是否选择抢的情况
// 选择当前节点：
// 左右孩子都不能被选中，故当前节点被选中情况下子树上被选中点的最大权值和为左孩子节点和右孩子节点不被选中的最大权值和加当前节点值，
// selected = node.Val + lUnSelected + rUnSelected
// 不选择当前节点：
// 左右孩子可以被选中，也可以不被选中；是孩子节点被选中和不被选中情况下权值较大值之和。
// unSelected = mathx.Max(rSelected, rUnSelected) + mathx.Max(lSelected, lUnSelected)
func dfs(node *define.TreeNode) (selected, unSelected int) {
	if node == nil {
		return
	}

	lSelected, lUnSelected := dfs(node.Left)
	rSelected, rUnSelected := dfs(node.Right)

	selected = node.Val + lUnSelected + rUnSelected
	unSelected = mathx.Max(rSelected, rUnSelected) + mathx.Max(lSelected, lUnSelected)

	return
}
