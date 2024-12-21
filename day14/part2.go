package day14

import "mrnateriver.io/advent_of_code_2024/shared"

func MeasureTimeForRobotsToFormChristmasTree() int {
	const lenX = 101
	const lenY = 103

	grid := shared.CreateDotGrid(lenX, lenY)

	input := readInput()

	steps := 0
	for steps < 10000 {
		for r := range input {
			input[r].move(lenX, lenY, 1)
		}
		steps++

		if !overlappingRobots(input) {
			break
		}
	}

	printGrid(input, grid)

	return steps
}
