package main

import (
	"fmt"
	. "github.com/weedge/leetcode/define"
)

func deleteDuplication(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}

	p := pHead
	mNode := map[int]int{}
	for p != nil {
		mNode[p.Val]++
		p = p.Next
	}
	p = pHead

	head := &ListNode{}
	tail := head
	for p != nil {
		if mNode[p.Val] == 1 {
			tail.Next = &ListNode{Val: p.Val, Next: nil}
			tail = tail.Next
		}
		p = p.Next
	}

	return head.Next
}

func main() {
	testCase := []int{1, 2, 3, 3, 4, 4, 5}
	pHead := &ListNode{}
	tail := pHead
	for _, val := range testCase {
		tail.Next = &ListNode{Val: val, Next: nil}
		tail = tail.Next
	}

	h := deleteDuplication(pHead.Next)
	fmt.Println(h)
}
