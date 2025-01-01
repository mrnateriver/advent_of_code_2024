package day22

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"strconv"
)

func CalcSecretNumbersSum() uint64 {
	const depth = 2000

	sum := uint64(0)
	for line := range shared.ReadInputLines("day22/input") {
		secretNumber, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Errorf("failed to parse secret number %v: %w", line, err))
		}

		nth := calcNthSecretNumber(uint64(secretNumber), depth)
		sum += nth
	}

	return sum
}

func calcNthSecretNumber(start uint64, n int) uint64 {
	if n <= 0 {
		return start
	}

	s1 := secretStage1(start)
	s2 := secretStage2(s1)
	s3 := secretStage3(s2)

	return calcNthSecretNumber(s3, n-1)
}

func secretStage1(start uint64) uint64 {
	return (start ^ (start << 6)) % modMask
}

func secretStage2(start uint64) uint64 {
	return (start ^ (start >> 5)) % modMask
}

func secretStage3(start uint64) uint64 {
	return (start ^ (start << 11)) % modMask
}

const modMask = 1 << 24
