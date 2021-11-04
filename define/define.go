package define

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() (str string) {
	if head == nil {
		return
	}
	p := head
	for p != nil {
		if len(str) > 0 {
			str = fmt.Sprintf("%s->%d", str, p.Val)
		} else {
			str = fmt.Sprintf("%d", p.Val)
		}
		p = p.Next
	}
	return
}

func CreateLink(arr []int) (head *ListNode) {
	pHead := &ListNode{}
	tail := pHead
	for _, val := range arr {
		tail.Next = &ListNode{Val: val, Next: nil}
		tail = tail.Next
	}

	return pHead.Next
}
