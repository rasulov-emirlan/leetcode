package implementstrstr

import "strings"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	return strings.Index(haystack, needle)
}
