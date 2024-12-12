package day04

func CountXmasCrosses() int {
	return countPatterns(3, findXmasCrossesPatterns)
}

func findXmasCrossesPatterns(lines *[]string, y int) int {
	if y < 1 || y > len(*lines)-2 {
		return 0
	}

	foundWords := 0
	lineLength := len((*lines)[0])

	for x := 1; x < lineLength-1; x++ {
		if (*lines)[y][x] != 'A' {
			continue
		}

		topLeft := (*lines)[y-1][x-1]
		topRight := (*lines)[y-1][x+1]
		bottomLeft := (*lines)[y+1][x-1]
		bottomRight := (*lines)[y+1][x+1]

		if ((topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')) && ((topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M')) {
			foundWords++
		}
	}

	return foundWords
}
