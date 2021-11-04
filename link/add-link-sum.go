package main

import (
	"fmt"
	. "github.com/weedge/leetcode/define"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addLinkSum(n, m int, link1, link2 *ListNode) *ListNode {
	p1, p2 := link1, link2
	a1, a2 := []int{}, []int{}
	a3 := []int{}

	for p1 != nil {
		a1 = append(a1, p1.Val)
		p1 = p1.Next
	}
	for p2 != nil {
		a2 = append(a2, p2.Val)
		p2 = p2.Next
	}
	c := 0
	for i, j := n-1, m-1; i >= 0 || j >= 0 || c > 0; i, j = i-1, j-1 {
		d := c
		if i >= 0 {
			d += a1[i]
		}
		if j >= 0 {
			d += a2[j]
		}

		c = d / 10
		v := d % 10
		a3 = append(a3, v)
	}
	head := &ListNode{}
	tail := head
	for i := len(a3) - 1; i >= 0; i-- {
		tail.Next = &ListNode{Val: a3[i], Next: nil}
		tail = tail.Next
	}

	return head.Next
}

func main() {
	testCases := [][2][]int{
		{{9, 6, 3}, {3, 7}},
	}

	for _, items := range testCases {
		h1 := CreateLink(items[0])
		fmt.Println(h1)
		h2 := CreateLink(items[1])
		fmt.Println(h2)
		h3 := addLinkSum(len(items[0]), len(items[1]), h1, h2)
		fmt.Println(h3)
	}
}
