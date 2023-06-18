package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	res := []int{0, 0}

	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			res[0] = i + 1
			res[1] = j + 1
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return res
}
