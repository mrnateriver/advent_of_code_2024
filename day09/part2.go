package day09

import (
	"log"
	"slices"
)

func CalcChecksumAfterIntegralDefragmentation() int {
	defrag, files := readInput()

	slots := make([][]slot, 10) // sizes -> offsets

	n := len(defrag)
	start, curLen := 0, 0
	for i := 0; i < n; i++ {
		id := defrag[i]
		if id >= 0 {
			if curLen > 0 {
				ml := min(curLen, 9)
				if slots[ml] == nil {
					slots[ml] = make([]slot, 0, 2)
				}
				slots[ml] = append(slots[ml], slot{start, curLen})
			}
			start = i + 1
			curLen = 0
		} else {
			curLen++
		}
	}
	if curLen > 0 {
		ml := min(curLen, 9)
		if slots[ml] == nil {
			slots[ml] = make([]slot, 0, 2)
		}
		slots[ml] = append(slots[ml], slot{start, curLen})
	}

	for i := len(files) - 1; i >= 0; i-- {
		moveFile(files[i], slots, defrag)
	}

	empty := findLastOccupiedIndex(defrag)
	log.Printf("Last occupied index: %d\n", empty)

	for k, v := range slots {
		log.Printf("Size %d slots:\n", k)
		for _, s := range v {
			log.Printf("   - %v\n", s)
		}
	}

	return checksum(defrag)
}

func moveFile(f file, slots [][]slot, defrag []int) {
	// Naive O(n^2) solution

	//n := len(defrag)
	//fsize := int(f.size)
	//fbound := f.offset + fsize
	//
	//var i int
	//start, curLen := 0, 0
	//for i = 0; i < n && i < fbound; i++ {
	//	id := defrag[i]
	//	if fits := curLen >= fsize; id >= 0 || fits {
	//		if fits {
	//			for fs := 0; fs < fsize; fs++ {
	//				defrag[f.offset+fs] = -1
	//				defrag[start+fs] = f.id
	//			}
	//			return
	//		}
	//		start = i + 1
	//		curLen = 0
	//	} else {
	//		curLen++
	//	}
	//}

	// FIXME: produces invalid answer
	for s := f.size; s < 10; s++ {
		sls := slots[s]
		if sls == nil || len(sls) == 0 {
			continue
		}

		sl := sls[0]
		if sl.offset > f.offset {
			continue
		}

		fsize := int(f.size)
		for fs := 0; fs < fsize; fs++ {
			defrag[f.offset+fs] = -1
			defrag[sl.offset+fs] = f.id
		}

		slots[s] = sls[1:]
		if sl.size > fsize {
			sl.size = sl.size - fsize
			sl.offset = sl.offset + fsize

			ml := min(sl.size, 9)
			slots[ml] = append(slots[ml], sl)
			slices.SortFunc(slots[ml], func(a, b slot) int {
				return a.offset - b.offset
			})
		}

		return
	}
}

func findLastOccupiedIndex(defrag []int) int {
	c := 0
	for i := len(defrag) - 1; i >= 0; i-- {
		if defrag[i] >= 0 {
			c = i
			break
		}
	}
	return c
}

type slot struct {
	offset int
	size   int
}
