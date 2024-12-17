package day09

import "mrnateriver.io/advent_of_code_2024/shared"

func readInput() (defrag []int, files []file) {
	input := shared.ReadInputLines("day09/input")
	line := <-input

	defrag = make([]int, 0, 2048)
	files = make([]file, 0, 1024)

	id, j := 0, 0
	expectFile := true
	for i := range line {
		size := parseDigit(line[i])

		sectorId := -1
		if expectFile {
			sectorId = id
			files = append(files, file{id, j, size})
			id++
		}

		for x := byte(0); x < size; x++ {
			defrag = append(defrag, sectorId)
			j++
		}

		expectFile = !expectFile
	}

	return
}

func checksum(defrag []int) int {
	checksum := 0
	for i, v := range defrag {
		if v >= 0 {
			checksum += v * i
		}
	}
	return checksum
}

func parseDigit(c byte) byte {
	return c - '0'
}

type file struct {
	id     int
	offset int
	size   byte
}
