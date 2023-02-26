package firstbadversion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 1 2 3 4 5 6 7 8 9 10 11 12 13
// l         m                 r
func firstBadVersion(n int) int {
	left := 0
	right := n
	mid := 0

	for left < right {
		mid = right / 2
		if isBadVersion(mid) {
			right -= mid
		}
	}

	return mid
}

func isBadVersion(v int) bool {
	return false
}
