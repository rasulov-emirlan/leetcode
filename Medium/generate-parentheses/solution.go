package generateparentheses

func backtrack(stack *[]byte, res *[]string, openN, closedN, N int) {
	if len(*stack) == N*2 {
		*res = append(*res, string(*stack))
		return
	}

	if openN < N {
		*stack = append(*stack, '(')
		backtrack(stack, res, openN+1, closedN, N)
		*stack = (*stack)[:len(*stack)-1]
	}

	if closedN < openN {
		*stack = append(*stack, ')')
		backtrack(stack, res, openN, closedN+1, N)
		*stack = (*stack)[:len(*stack)-1]
	}
}

func generateParenthesis(n int) []string {
	stack := []byte{}
	res := []string{}

	backtrack(&stack, &res, 0, 0, n)

	return res
}
