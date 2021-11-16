package main

import (
	"fmt"
	. "github.com/weedge/leetcode/define"
)

//从根节点开始遍历，所以 中右左 -反转-> 左右中
func postorderTraversal(root *TreeNode) []int {
	t := root
	if t == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{}
	for t != nil || len(stack) > 0 {
		for t != nil {
			res = append(res, t.Val) //中
			stack = append(stack, t)
			t = t.Right //右
		}
		t = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		t = t.Left //左
	}
	//反转
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}

//从根节点开始遍历，所以 中右左 -反转-> 左右中
//根节点开始，中右左
// 先把根节点入栈，出栈，选择，
// 出栈节点的左节点不空，入栈
// 出栈节点的右节点不空，入栈
// 重复，直到栈为空
// 反转
func postorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		res = append(res, item.Val)
		stack = stack[:len(stack)-1]

		if item.Left != nil {
			stack = append(stack, item.Left)
		}
		if item.Right != nil {
			stack = append(stack, item.Right)
		}
	}
	//反转
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}

//左右中
//如果先把根节点入栈，在按栈是否空遍历操作,需要如下操作：
//栈中加入是否选标记NULL,出栈时判断top是否时NULL,是NULL时选择,不是NULL,按反着顺序(中右左)入栈
func postorderTraversalV3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if item != nil {
			stack = append(stack, item) //中
			stack = append(stack, nil)  //设置未选中

			if item.Right != nil {
				stack = append(stack, item.Right) // 右
			}
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

	res := postorderTraversalV2(treeNode)
	fmt.Println(res)
}
