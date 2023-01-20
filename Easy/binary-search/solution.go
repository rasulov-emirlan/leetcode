package binarysearch

// Accepted

func search(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
		mid   = 0
	)

	for left <= right {
		mid = (right + left) / 2
		if target == nums[mid] {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
