package removeduplicatesfromsortedlist

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	cases := []struct {
		desc  string
		input *ListNode
		want  *ListNode
	}{
		{
			desc:  "nil input",
			input: nil,
			want:  nil,
		},
		{
			desc:  "single node",
			input: &ListNode{Val: 1},
			want:  &ListNode{Val: 1},
		},
		{
			desc:  "two nodes",
			input: &ListNode{Val: 1, Next: &ListNode{Val: 1}},
			want:  &ListNode{Val: 1},
		},
		{
			desc:  "three nodes",
			input: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			want:  &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
		{
			desc:  "four nodes",
			input: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
			want:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			desc:  "a lot of same numbers",
			input: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}},
			want:  &ListNode{Val: 1},
		},
		{
			desc:  "1,1,2,2",
			input: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}},
			want:  &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			got := deleteDuplicates(s.input)
			if !reflect.DeepEqual(got, s.want) {
				t.Errorf("wanted: %v, got: %v", s.want, got)
			}
		})
	}
}
