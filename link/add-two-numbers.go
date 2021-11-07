package main

import (
	"fmt"
	. "github.com/weedge/leetcode/define"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tmp *ListNode
	p, q := l1, l2
	a, b, c := 0, 0, 0
	for q != nil || p != nil || c > 0 {
		if q == nil {
			a = 0
		} else {
			a = q.Val
			q = q.Next
		}
		if p == nil {
			b = 0
		} else {
			b = p.Val
			p = p.Next
		}

		val := int((a + b + c) % 10)
		c = int((a + b + c) / 10)
		if head == nil {
			head = &ListNode{Val: val}
			tmp = head
		} else {
			tmp.Next = &ListNode{Val: val}
			tmp = tmp.Next
		}
	}

	return head
}

func main() {
	testCases := [][2][]int{
		{{2, 4, 3}, {5, 6, 4}},
		{{9, 6, 3}, {3, 7}},
	}

	for _, items := range testCases {
		h1 := CreateLink(items[0])
		fmt.Println(h1)
		h2 := CreateLink(items[1])
		fmt.Println(h2)
		h3 := addTwoNumbers(h1, h2)
		fmt.Println(h3)
	}
}
