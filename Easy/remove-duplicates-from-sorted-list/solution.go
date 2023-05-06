package removeduplicatesfromsortedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var s string
	i := l
	for i != nil {
		s += fmt.Sprintf("%d, ", i.Val)
		i = i.Next
	}
	return s
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	it := head.Next
	for it != nil {
		if it.Val == prev.Val {
			prev.Next = it.Next
			it = it.Next
			continue
		}
		prev = it
		it = it.Next
	}
	return head
}
