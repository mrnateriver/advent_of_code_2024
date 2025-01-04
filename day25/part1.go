package day25

import "mrnateriver.io/advent_of_code_2024/shared"

func CountLockKeyCombinations() int {
	locks, keys := make([]row, 0, 4), make([]row, 0, 4)

	r, lock, heights := 0, false, row{}
	for line := range shared.ReadInputLines("day25/input") {
		if line == "" {
			if lock {
				locks = append(locks, heights)
			} else {
				keys = append(keys, heights)
			}

			r = 0
			lock = false
			heights = row{}
			continue
		}

		if r == 0 {
			lock = baseline(line)
		} else if r < 6 {
			for i := 0; i < 5; i++ {
				if line[i] == '#' {
					heights[i]++
				}
			}
		}
		r++
	}
	if lock {
		locks = append(locks, heights)
	} else {
		keys = append(keys, heights)
	}

	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			if keyFitsLock(key, lock) {
				count++
			}
		}
	}

	return count
}

func keyFitsLock(key, lock row) bool {
	for i := range key {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}

func baseline(line string) bool {
	for i := range line {
		if line[i] != '#' {
			return false
		}
	}
	return true
}

type row [5]byte
