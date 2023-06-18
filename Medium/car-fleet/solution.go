package carfleet

import "sort"

// i dont get this problem at all :(

func carFleet(target int, position, speed []int) int {
	var (
		res int
		m   = make(map[int]float64)
	)

	for i, pos := range position {
		m[pos] = float64(target-pos) / float64(speed[i])
	}

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var max float64
	for i := len(keys) - 1; i >= 0; i-- {
		if m[keys[i]] > max {
			max = m[keys[i]]
			res++
		}
	}

	return res
}
