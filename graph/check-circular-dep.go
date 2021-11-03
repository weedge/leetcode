package main

import (
	"fmt"
)

// 循环依赖检查
// [["A","B"],["A1","B"],["B","C"],["B","D"],["B","E"],["C","A"]] => true
// [["A","B"],["B","C"],["C","D"]] => false
// 没有依赖的节点就是入度为0的节点，没有依赖的节点入队列;
// 出队列放入依赖顺序节点数组中;
// 被依赖这个节点的出度子节点入度数目-1, 入度为0的没有依赖的节点入队列，直到队列为空
// 如果依赖顺序节点数组长度和节点集合长度一样，返回true,没有循环依赖,否则为false
func checkCircularDep(depArr [][]string) bool {
	//节点出度子节点
	subNodes := map[string][]string{}
	//节点入度数目
	nodeInDeg := map[string]int{}
	//节点集合
	nodes := map[string]struct{}{}
	//节点依赖顺序
	depRes := []string{}

	for _, item := range depArr {
		if _, ok := subNodes[item[0]]; ok {
			subNodes[item[0]] = append(subNodes[item[0]], item[1])
		} else {
			subNodes[item[0]] = []string{item[1]}
		}

		for i := range item {
			if _, ok := nodes[item[i]]; !ok {
				nodes[item[i]] = struct{}{}
			}
		}

		nodeInDeg[item[1]]++
	}

	queue := []string{}
	for node := range nodes {
		if cn, ok := nodeInDeg[node]; !ok {
			if cn == 0 {
				queue = append(queue, node)
			}
		}
	}

	fmt.Printf("nods:%+v\tsubNodes:%+v\t nodeInDeg:%+v\t queue:%+v\n", nodes, subNodes, nodeInDeg, queue)

	for len(queue) > 0 {
		a := queue[0]
		queue = queue[1:]

		depRes = append(depRes, a)

		for _, node := range subNodes[a] {
			nodeInDeg[node]--
			if nodeInDeg[node] == 0 {
				queue = append(queue, node)
			}
		}
	}

	if len(depRes) == len(nodes) {
		fmt.Printf("depRunRes:%+v\n", depRes)
		return true
	}

	return false
}

func main() {
	testCases := [][][]string{
		{{"A", "B"}, {"A1", "B"}, {"B", "C"}, {"B", "D"}, {"B", "E"}, {"C", "A"}},
		{{"A", "B"}, {"A1", "B"}, {"B", "C"}, {"B", "D"}, {"B", "E"}},
		{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}},
	}
	for _, depStrs := range testCases {
		res := checkCircularDep(depStrs)
		fmt.Printf("depStrs:%+v\n res: %v\n", depStrs, res)
	}
}
