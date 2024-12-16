package day04

import "iter"

func CountXmasOccurrences() int {
	return countPatterns(4, findXmasPatterns)
}

func findXmasPatterns(lines []string, y int) int {
	foundWords := 0
	lineLength := len(lines[0])

	for x := 0; x < lineLength; x++ {
		if lines[y][x] != 'X' {
			continue
		}

	vectors:
		for vector := range wordVectors(lines, x, y) {
			expected := byte('M')

			for char := range vector {
				if char == expected {
					expected = nextLetter(expected)
				} else {
					continue vectors
				}
			}

			if expected == 0 {
				foundWords++
			}
		}
	}

	return foundWords
}

func wordVectors(lines []string, x0, y0 int) iter.Seq[iter.Seq[byte]] {
	return func(yield func(iter.Seq[byte]) bool) {
		for diffX := -1; diffX < 2; diffX++ {
			for diffY := -1; diffY < 2; diffY++ {
				if diffX == 0 && diffY == 0 {
					continue
				}

				if !yield(wordVector(lines, x0, y0, diffX, diffY)) {
					return
				}
			}
		}
	}
}

func wordVector(lines []string, x0, y0, diffX, diffY int) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		lineLength := len(lines[0])
		totalLines := len(lines)

		for i := 1; i < 4; /*[X]MAS*/ i++ {
			x := x0 + diffX*i
			y := y0 + diffY*i

			if x < 0 || x >= lineLength || y < 0 || y >= totalLines {
				return
			}

			if !yield(lines[y][x]) {
				return
			}
		}
	}
}

func nextLetter(letter byte) byte {
	switch letter {
	case 'X':
		return 'M'

	case 'M':
		return 'A'

	case 'A':
		return 'S'

	default:
		return 0
	}
}
