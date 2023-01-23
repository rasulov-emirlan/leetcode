package linkedlistcycle

// Accepted

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := make(map[*ListNode]int)

	iterator := head

	for iterator != nil {
		m[iterator]++
		if m[iterator] > 1 {
			return true
		}
		iterator = iterator.Next
	}

	return false
}
