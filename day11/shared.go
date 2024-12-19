package day11

import (
	"fmt"
	"strconv"
	"strings"
)

func readInput() []string {
	input := "572556 22 0 528 4679021 1 10725 2790"
	return strings.Split(input, " ")
}

func countStonesAfterBlinking(stones []string, times int) (count int) {
	dp := make(map[pair]int)

	for i := 0; i < len(stones); i++ {
		count += countStone(stones[i], times, dp)
	}

	return
}

func countStone(stone string, times int, dp map[pair]int) int {
	if times == 0 {
		return 1
	}

	if c, ok := dp[pair{stone, times}]; ok {
		return c
	}

	c := 0
	a, b := splitStone(stone)
	if b == "" {
		c = countStone(a, times-1, dp)
	} else {
		c = countStone(a, times-1, dp) + countStone(b, times-1, dp)
	}
	dp[pair{stone, times}] = c
	return c
}

func splitStone(s string) (string, string) {
	if s == "0" {
		return "1", ""
	}

	if len(s)%2 == 0 {
		b := s[len(s)/2:]
		b = strings.TrimLeft(b, "0")
		if b == "" {
			b = "0"
		}
		return s[:len(s)/2], b
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("invalid number %s", s))
	}

	return strconv.Itoa(num * 2024), ""
}

type pair struct {
	stone string
	times int
}
