package mergetwosortedlists

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	tempNode := &ListNode{}
	answer := tempNode

	i := list1
	j := list2

	for i != nil && j != nil {
		if i.Val < j.Val {
			tempNode.Val = i.Val
			i = i.Next
			tempNode.Next = &ListNode{}
			tempNode = tempNode.Next
			continue
		}
		tempNode.Val = j.Val
		j = j.Next
		tempNode.Next = &ListNode{}
		tempNode = tempNode.Next
	}
	if j != nil {
		tempNode.Val = j.Val
		tempNode.Next = j.Next
	}
	if i != nil {
		tempNode.Val = i.Val
		tempNode.Next = i.Next
	}
	return answer
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
