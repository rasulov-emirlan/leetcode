package mergetwosortedlists

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		inputA *ListNode
		inputB *ListNode
		answer *ListNode
	}{
		{
			inputA: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			inputB: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			answer: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
		{
			inputA: &ListNode{
				Val:  5,
				Next: nil,
			},
			inputB: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			answer: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			answer := mergeTwoLists(tC.inputA, tC.inputB)
			i := answer
			j := tC.answer

			for i != nil && j != nil {
				if i.Val != j.Val {
					t.Errorf("what wanted: %s, what we got: %s", tC.answer, answer)
				}
				i = i.Next
				j = j.Next
			}
			fmt.Println(answer)
		})
	}
}
