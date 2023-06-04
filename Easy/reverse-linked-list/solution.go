package reverselinkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	it := l
	str := ""

	for it != nil {
		str += " -> " + fmt.Sprint(it.Val)
		it = it.Next
	}

	return str
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := new(ListNode)
	it := head

	for it != nil {
		temp := it.Next
		it.Next = prev
		prev = it
		it = temp
	}

	head.Next = nil

	return prev
}
