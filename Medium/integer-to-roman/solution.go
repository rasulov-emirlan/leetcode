package integertoroman

import (
	"strings"
)

type Pair struct {
	Roman   string
	Integer int
}

var romanNums = []Pair{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// solved with a lot of help from neetcode and chatgpt

func intToRoman(num int) string {
	var res strings.Builder

	for i := 0; num > 0; i++ {
		for num >= romanNums[i].Integer {
			res.WriteString(romanNums[i].Roman)
			num -= romanNums[i].Integer
		}
	}

	return res.String()
}
