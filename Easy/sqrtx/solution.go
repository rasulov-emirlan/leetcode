package sqrtx

// Accepted. Solved with help from Github Copilot. Thanks, Copilot!

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	low, high := 1, x/2
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}
