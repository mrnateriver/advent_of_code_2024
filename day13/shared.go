package day13

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"regexp"
	"strconv"
)

func readInput() []machine {
	buttonRegex := regexp.MustCompile("Button (A|B): X\\+(\\d+), Y\\+(\\d+)")
	prizeRegex := regexp.MustCompile("Prize: X=(\\d+), Y=(\\d+)")

	machines := make([]machine, 0, 4)

	m := machine{}
	lineIdx := 0
	for line := range shared.ReadInputLines("day13/input") {
		if lineIdx < 2 {
			matches := buttonRegex.FindStringSubmatch(line)
			if matches == nil {
				continue
			}

			x, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(fmt.Errorf("failed to parse button x on lineIdx %v: %w", lineIdx, err))
			}

			y, err := strconv.Atoi(matches[3])
			if err != nil {
				panic(fmt.Errorf("failed to parse button y on lineIdx %v: %w", lineIdx, err))
			}

			if matches[1] == "A" {
				m.ax, m.ay = x, y
			} else {
				m.bx, m.by = x, y
			}
		} else if lineIdx == 2 {
			matches := prizeRegex.FindStringSubmatch(line)
			if matches == nil {
				continue
			}

			x, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(fmt.Errorf("failed to parse prize x on lineIdx %v: %w", lineIdx, err))
			}

			y, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(fmt.Errorf("failed to parse prize y on lineIdx %v: %w", lineIdx, err))
			}

			m.tx, m.ty = x, y
		} else {
			machines = append(machines, m)
			m = machine{}
		}

		lineIdx = (lineIdx + 1) % 4
	}
	machines = append(machines, m)

	return machines
}

type machine struct {
	ax, ay, bx, by, tx, ty int
}

func solve(m machine) (a, b int) {
	b = calcB(m)
	a = calcA(m, b)

	return a, b
}

func calcA(m machine, b int) int {
	i := m.tx - m.bx*b
	if i%m.ax != 0 {
		return 0
	}

	return i / m.ax
}

func calcB(m machine) int {
	p1 := m.ty*m.ax - m.ay*m.tx
	p2 := m.by*m.ax - m.ay*m.bx
	if p1%p2 != 0 {
		return 0
	}

	return p1 / p2
}
