package addtwonumbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	answerr := &ListNode{}

	iter := answerr
	iter1 := l1
	iter2 := l2

	check := false

	for iter1 != nil || iter2 != nil {
		if iter1 != nil {
			iter.Val += iter1.Val
			iter1 = iter1.Next
		}
		if iter2 != nil {
			iter.Val += iter2.Val
			iter2 = iter2.Next
		}

		if iter.Val > 9 {
			check = true
			iter.Val -= 10
		}
		if iter1 != nil || iter2 != nil || check {
			iter.Next = &ListNode{}
			iter = iter.Next
			if check {
				iter.Val = 1
				check = false
			}
		}
	}
	return answerr
}

// numberl1 := []byte("")
// numberl2 := []byte("")

// i := l1
// j := l2

// // ""
// // "" + "23"
// // "23" + "343434"
// // "23343434"
// for i != nil {
// 	a := []byte(strconv.Itoa(i.Val))
// 	numberl1 = append(numberl1, a...)

// 	i = i.Next
// }

// for j != nil {
// 	b := []byte(strconv.Itoa(j.Val))
// 	numberl2 = append(numberl2, b...)

// 	j = j.Next
// }

// a, _ := strconv.ParseInt(reverse(numberl1), 10, 64)
// b, _ := strconv.ParseInt(reverse(numberl2), 10, 64)

// result := a + b

// answer := &ListNode{}
// realAnswer := answer

// for result > 0 {
// 	answer.Val = result % 10
// 	log.Println(answer.Val)
// 	result /= 10
// 	if result > 0 {
// 		log.Println("checking")
// 		answer.Next = &ListNode{}
// 		answer = answer.Next
// 	}
// }

// return realAnswer
// }

func reverse(rns []byte) string {
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
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
