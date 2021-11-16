package main

import (
	"fmt"
	. "github.com/weedge/leetcode/define"
)

//根节点开始，左中右
func inorderTraversal(root *TreeNode) []int {
	t := root
	if t == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	res := []int{}

	for t != nil || len(stack) > 0 {
		for t != nil {
			stack = append(stack, t)
			t = t.Left // 左
		}
		t = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, t.Val) // 中
		t = t.Right              // 右
	}

	return res
}

//根节点开始，左中右
//如果先把根节点入栈，在按栈是否空遍历操作,需要如下操作：
//栈中加入是否选标记NULL,出栈时判断top是否时NULL,是NULL时选择,不是NULL,按反着顺序(右中左)入栈
func inorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if item != nil {
			if item.Right != nil {
				stack = append(stack, item.Right) // 右
			}

			stack = append(stack, item) //中
			stack = append(stack, nil)  //设置未选中

			if item.Left != nil {
				stack = append(stack, item.Left) //左
			}
		} else {
			item = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res = append(res, item.Val)
		}

	}

	return res
}

func main() {
	treeNode2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	treeNode1 := &TreeNode{Val: 1, Left: treeNode2, Right: nil}
	treeNode := &TreeNode{Val: 0, Left: nil, Right: treeNode1}

	res := inorderTraversalV2(treeNode)
	fmt.Println(res)
}
