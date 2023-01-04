package romantointeger

// Accepted

var romanNums = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var sum int

	for i := 0; i < len(s); i++ {
		curr := romanNums[s[i]]
		switch curr {
		case 1: // 1 is I
			c, ok := peekNext(s, i)
			if !ok {
				break
			}
			if c == 'V' || c == 'X' {
				sum--
				continue
			}
		case 10: // 10 is X
			c, ok := peekNext(s, i)
			if !ok {
				break
			}

			if c == 'L' || c == 'C' {
				sum -= 10
				continue
			}
		case 100: // 100 is C
			c, ok := peekNext(s, i)
			if !ok {
				break
			}

			if c == 'D' || c == 'M' {
				sum -= 100
				continue
			}
		}
		sum += curr
	}

	return sum
}

func peekNext(s string, i int) (byte, bool) {
	if i+1 < len(s) {
		return s[i+1], true
	}
	return 0, false
}
