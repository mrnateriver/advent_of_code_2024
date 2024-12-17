package day09

func CalcChecksumAfterIntegralDefragmentation() int {
	defrag, files := readInput()

	for i := len(files) - 1; i >= 0; i-- {
		moveFile(files[i], defrag)
	}

	return checksum(defrag)
}

func moveFile(f file, defrag []int) {
	n := len(defrag)
	fsize := int(f.size)
	fbound := f.offset + fsize

	var i int
	start, curLen := 0, 0
	for i = 0; i < n && i < fbound; i++ {
		id := defrag[i]
		if fits := curLen >= fsize; id >= 0 || fits {
			if fits {
				for fs := 0; fs < fsize; fs++ {
					defrag[f.offset+fs] = -1
					defrag[start+fs] = f.id
				}
				return
			}
			start = i + 1
			curLen = 0
		} else {
			curLen++
		}
	}
}
