package evaluatereversepolishnotation

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+":
			plusRes := stack[len(stack)-1] + stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, plusRes)
		case "-":
			minusRes := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, minusRes)
		case "/":
			divRes := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, divRes)
		case "*":
			mulRes := stack[len(stack)-1] * stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, mulRes)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[len(stack)-1]
}
