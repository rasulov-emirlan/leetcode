package mergesortedarray

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m, 0; i <= len(nums1)-1; {
		nums1[i] = nums2[j]
		i++
		j++
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
}
