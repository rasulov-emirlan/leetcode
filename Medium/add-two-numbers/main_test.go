package addtwonumbers

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	testCases := []struct {
		inputA *ListNode
		inputB *ListNode
		answer *ListNode
	}{
		{
			inputA: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			inputB: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			answer: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			inputA: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val:  9,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			inputB: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
			},
			answer: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val:  1,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			answer := addTwoNumbers(tC.inputA, tC.inputB)
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
