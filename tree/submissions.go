package main

import (
	"fmt"
	. "github.com/weedge/leetcode/define"
)

//获取全路径集合
//广度优先,2个同步队列，用于层次遍历节点，和对应节点路径；
//到了叶子节点，将路径写入集合
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	paths := []string{}
	nQ := []*TreeNode{root}
	pQ := []string{""}
	for len(nQ) > 0 {
		node, path := nQ[0], pQ[0]
		nQ = nQ[1:]
		pQ = pQ[1:]

		path += fmt.Sprintf("%d", node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
		}

		if node.Left != nil {
			nQ = append(nQ, node.Left)
			pQ = append(pQ, fmt.Sprintf("%s->", path))
		}
		if node.Right != nil {
			nQ = append(nQ, node.Right)
			pQ = append(pQ, fmt.Sprintf("%s->", path))
		}
	}

	return paths
}

//获取全路径集合
//深度优先递归, 路径需要一路带上，到了叶子节点将路径写入集合
//类似trace功能
func binaryTreePathsV2(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	paths := []string{}
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		ps := path
		ps += fmt.Sprintf("%d", node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, ps)
		} else {
			ps += "->"
			dfs(node.Left, ps)
			dfs(node.Right, ps)
		}
		return
	}
	dfs(root, "")

	return paths
}

func main() {

}
