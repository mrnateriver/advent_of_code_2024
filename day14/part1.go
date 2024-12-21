package day14

func CountRobotsInQuadrantsAfter100Secs() int {
	const steps = 100
	const lenX = 101
	const lenY = 103

	input := readInput()
	for r := range input {
		input[r].move(lenX, lenY, steps)
	}

	return calcSafetyFactor(input, lenX, lenY)
}
