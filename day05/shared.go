package day05

import (
	"fmt"
	"strconv"
	"strings"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func parseUpdates(parser func(upd update, rules [][]precedenceRule) int) (sum int) {
	input := readInput()

	// Page numbers are 0 < num < 100, so we can use a super-duper-fast stack-allocated array instead of a hashmap
	beforeRules := [100][]precedenceRule{}

	for entry := range input {
		switch val := entry.(type) {
		case precedenceRule:
			if beforeRules[val.page] == nil {
				beforeRules[val.page] = make([]precedenceRule, 0, 1)
			}
			beforeRules[val.page] = append(beforeRules[val.page], val)

		case update:
			sum += parser(val, beforeRules[:])
		}
	}

	return sum
}

func readInput() <-chan any {
	parsedRules := false
	return shared.ReadInput("day05/input", func(line string) (any, error) {
		if line == "" {
			parsedRules = true
			return nil, nil
		}

		if !parsedRules {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid rule: %s", line)
			}

			page, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("failed to parse %s as int, line %s; %w", parts[0], line, err)
			}

			beforePage, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("failed to parse %s as int, line %s; %w", parts[1], line, err)
			}

			return precedenceRule{page, beforePage}, nil

		} else {
			parts := strings.Split(line, ",")
			if len(parts) < 2 {
				return nil, fmt.Errorf("invalid update: %s", line)
			}

			pages := make(update, 0, len(parts))
			for _, part := range parts {
				page, err := strconv.Atoi(part)
				if err != nil {
					return nil, fmt.Errorf("failed to parse %s as int, line %s; %w", part, line, err)
				}
				pages = append(pages, page)
			}

			return pages, nil
		}
	})
}

type precedenceRule struct {
	page       int
	beforePage int
}

type update []int

func getMiddlePage(upd update) int {
	mid := len(upd) / 2
	return upd[mid]
}
