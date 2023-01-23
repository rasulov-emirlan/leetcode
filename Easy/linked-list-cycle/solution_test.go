package linkedlistcycle

import "testing"

func TestHasCycle(t *testing.T) {
	example1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
		},
	}
	example11 := ListNode{
		Val:  -4,
		Next: &example1,
	}
	example1.Next.Next = &example11
	testCases := []struct {
		desc string
		head *ListNode
		want bool
	}{
		{
			desc: "Example 1",
			head: &ListNode{
				Val:  3,
				Next: &example1,
			},
			want: true,
		},
		{
			desc: "Example 2",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			want: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := hasCycle(tC.head)
			if got != tC.want {
				t.Errorf("hasCycle(%v) = %v, want %v", tC.head, got, tC.want)
			}
		})
	}
}
