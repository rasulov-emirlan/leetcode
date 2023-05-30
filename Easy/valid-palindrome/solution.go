package validpalindrome

import (
	"strings"
	"unicode"
)

// a faster solution of O(n) is possible if we check palindrome inside of the first function
// and it is also a O(1) in terms of memory. but i am too lazy for that

func isPalindrome(s string) bool {
	tmp := strings.Builder{}
	for _, char := range s {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			tmp.WriteRune(unicode.ToLower(char))
		}
	}
	s = tmp.String()

	for l, r := 0, len(s)-1; l <= len(s)/2 && r >= len(s)/2; {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
