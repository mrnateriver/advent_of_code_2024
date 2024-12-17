package day09

func CalcChecksumAfterDefragmentation() int {
	defrag, files := readInput()

	var j int
	n := len(defrag)
	x := len(files) - 1
	for j = 0; j < n; j++ {
		if defrag[j] < 0 {
			break
		}
	}
	for x >= 0 && j < n {
		f := files[x]

		if j > f.offset {
			break
		}

		j = moveFileBlocks(f, j, n, defrag)
		x--
	}

	return checksum(defrag)
}

func moveFileBlocks(f file, j, n int, defrag []int) int {
	for s := byte(0); s < f.size && j < n; s++ {
		defrag[f.offset+int(s)] = -1
		defrag[j] = f.id

		for j < n && defrag[j] >= 0 {
			j++
		}
	}
	return j
}
