package plusone

// Accepted

func plusOne(digits []int) []int {
	var leftovers int
	var res []int
	summing := digits[len(digits)-1] + 1

	if summing > 9 {
		leftovers = 1
		res = make([]int, len(digits)+1)
		copy(res[1:], digits[:len(digits)-1])
		res[len(res)-1] = summing - 10
	} else {
		leftovers = 0
		res = make([]int, len(digits))
		copy(res, digits)
		res[len(res)-1] = summing
	}

	for i := len(res) - 2; i >= 0 && leftovers == 1; i-- {
		summing = res[i] + leftovers

		if summing > 9 {
			leftovers = 1
			res[i] = summing - 10
			continue
		}

		leftovers = 0
		res[i] = summing
	}

	if res[0] == 0 {
		return res[1:]
	}

	return res
}
