package addbinary

// Accepted

import "strings"

func addBinary(a string, b string) string {
	if a == "0" && b == "0" {
		return "0"
	}
	var (
		res      = ""
		leftover = 0
	)

	if len(a) < len(b) {
		a, b = b, a
	}

	i := len(a) - 1

	for j := len(b) - 1; i >= 0 && j >= 0; {
		sum := int(a[i]-'0') + int(b[j]-'0') + leftover
		res, leftover = calc(sum, res, leftover)

		i--
		j--
	}

	for ; i >= 0; i-- {
		sum := int(a[i]-'0') + leftover
		res, leftover = calc(sum, res, leftover)
	}

	res, _ = calc(leftover, res, leftover)
	return strings.TrimLeft(res, "0")
}

func calc(sum int, res string, leftover int) (string, int) {
	switch sum {
	case 2:
		return "0" + res, 1
	case 3:
		return "1" + res, 1
	case 0:
		return "0" + res, 0
	case 1:
		return "1" + res, 0
	}
	return res, leftover
}
