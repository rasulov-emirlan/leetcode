package palindromenumber

import (
	"strconv"
)

func isPalindrome(x int) bool {
	var s string = strconv.Itoa(x)
	var l int = len(s) / 2
	for i, j := 0, len(s)-1; i < l; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
