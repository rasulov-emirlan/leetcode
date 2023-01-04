package lengthoflastword

// Accepted

func lengthOfLastWord(s string) int {
	var sum int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if sum > 0 {
				return sum
			}
			continue
		}
		sum++
	}
	return sum
}
