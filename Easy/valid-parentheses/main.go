package validparentheses

func isValid(r string) bool {
	s := []byte(r)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ')':
			if i-1 >= 0 {
				if s[i-1] != '(' {
					return false
				}
				s = append(s[:i-1], s[i+1:]...)
				i = 0
			}
		case '}':
			if i-1 >= 0 {
				if s[i-1] != '{' {
					return false
				}
				s = append(s[:i-1], s[i+1:]...)
				i = 0
			}
		case ']':
			if i-1 >= 0 {
				if s[i-1] != '[' {
					return false
				}
				s = append(s[:i-1], s[i+1:]...)
				i = 0
			}
		}
	}
	return len(s) == 0
}

func find(c byte, s string, from int) bool {
	for i := from; i < len(s); i++ {

	}
	return true
}