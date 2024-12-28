package day17

import (
	"fmt"
	"strconv"
)

func FindSelfProducingProgramRegister() int {
	c := readInput()

	candidates := []int{0}
	for i := len(c.program) - 1; i >= 0; i-- {
		opcode := c.program[i]

		nextCandidates := make([]int, 0)
		for _, candidate := range candidates {
			// Multiply by 8 because OUT opcode does modulo 8 of its operand, so there
			// can be at most 8 ints that result in the necessary value
			shifted := candidate << 3
			for attempt := shifted; attempt <= shifted+8; attempt++ {
				c := cpu{
					regA:    attempt,
					regB:    0,
					regC:    0,
					program: c.program,
					ptr:     0,
				}

				firstOutput := runUntilFirstOutput(&c, false)
				if firstOutput == opcode {
					nextCandidates = append(nextCandidates, attempt)
				}
			}
		}

		candidates = nextCandidates
		if len(candidates) == 0 {
			break
		}
	}

	for _, cand := range candidates {
		if cand > 0 {
			return cand
		}
	}

	return -1
}

func runUntilFirstOutput(c *cpu, d bool) int {
	pl := len(c.program)
	for c.ptr+1 < pl {
		output := runStep(c, d)
		if output != "" {
			outputInt, err := strconv.Atoi(output)
			if err != nil {
				panic(fmt.Errorf("failed to convert output %v to int: %w", output, err))
			}
			return outputInt
		}
	}
	panic(fmt.Errorf("program did not produce any output"))
}
