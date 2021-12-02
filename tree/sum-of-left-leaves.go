package main

import (
	. "github.com/weedge/leetcode/define"
)

//左叶子和
//广度搜索 队列
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLeafVals := []int{}
	rightLeafVals := []int{}

	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				leftLeafVals = append(leftLeafVals, node.Left.Val)
			}
			q = append(q, node.Left)
		}
		if node.Right != nil {
			if node.Right.Left == nil && node.Right.Right == nil {
				rightLeafVals = append(rightLeafVals, node.Right.Val)
			}
			q = append(q, node.Right)
		}
	}

	res := 0
	//fmt.Println(leftLeafVals)
	//fmt.Println(rightLeafVals)
	for _, num := range leftLeafVals {
		res += num
	}

	return res
}

//左叶子和
//深度搜索 递归
func sumOfLeftLeavesV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLeafVals := []int{}
	rightLeafVals := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				leftLeafVals = append(leftLeafVals, node.Left.Val)
			} else {
				dfs(node.Left)
			}
		}

		if node.Right != nil {
			if node.Right.Left == nil && node.Right.Right == nil {
				rightLeafVals = append(rightLeafVals, node.Right.Val)
			} else {
				dfs(node.Right)
			}
		}
	}
	dfs(root)
	res := 0
	//fmt.Println(leftLeafVals)
	//fmt.Println(rightLeafVals)
	for _, num := range leftLeafVals {
		res += num
	}
	return res
}

func main() {

}
