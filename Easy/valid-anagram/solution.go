package validanagram

// Accepted

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if m[t[i]] <= 0 {
			return false
		}
		m[t[i]]--
	}

	return true
}
