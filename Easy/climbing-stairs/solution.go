package climbingstairs

func climbStairs(n int) int {
	one, two := 1, 1

	for i := 2; i <= n; i++ {
		one, two = two, one+two
	}

	return two
}
