package containerwithmostwater

func maxArea(height []int) int {
	biggest := 0

	l := 0
	r := len(height) - 1

	for l < r {
		if height[l] < height[r] {
			if area := height[l] * (r - l); area > biggest {
				biggest = area
			}
			l++
			continue
		}
		if area := height[r] * (r - l); area > biggest {
			biggest = area
		}
		r--
	}

	return biggest
}
