package main

//	   1
//	1     1
// 1 1   * *
type Node struct {
	Left, Right *Node
}

// n/2 O(n)
func insertQ(root, node *Node) {
	if root == nil {
		root = node
		return
	}
	q := []*Node{root}
	for len(q) > 0 {
		tmpQ := q
		q = nil
		for _, item := range tmpQ {
			if item.Left != nil {
				q = append(q, item.Left)
			} else {
				item.Left = node
				return
			}

			if item.Right != nil {
				q = append(q, item.Right)
			} else {
				item.Right = node
				return
			}
		}
	}
}
//between(log(n),n) O(logN)
func insertByDFS(root, node *Node) {
	if root == nil {
		root = node
		return
	}
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node.Left == nil {
			node.Left = node
			return
		}
		if node.Right == nil {
			node.Right = node
			return
		}
		cur := node
		lLen, rLen := 0, 0
		for cur != nil {
			lLen++
			cur = cur.Left
		}
		cur = node
		for cur != nil {
			rLen++
			cur = cur.Right
		}

		if lLen == rLen {
			dfs(node.Left)
		}
		if lLen>rLen{
			dfs(node.Right)
		}
		if lLen<rLen{
			panic("no")
		}
	}
	dfs(root)
}

func main() {

}
