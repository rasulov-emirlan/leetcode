package reverselinkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	cases := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head:     &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			expected: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
		},
		{
			head:     &ListNode{1, nil},
			expected: &ListNode{1, nil},
		},
		{
			head:     nil,
			expected: nil,
		},
	}

	for _, c := range cases {
		actual := reverseList(c.head)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("expected: %v, actual: %v", c.expected, actual)
		}
	}
}
