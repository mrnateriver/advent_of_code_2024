package day07

import (
	"fmt"
	"strconv"
	"strings"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func readInput() <-chan equation {
	return shared.ReadInput("day07/input", func(line string) (equation, error) {
		split := strings.Split(line, ":")
		if len(split) != 2 {
			return equation{}, fmt.Errorf("invalid line %s", line)
		}

		result, err := strconv.ParseUint(split[0], 10, 64)
		if err != nil {
			return equation{}, fmt.Errorf("failed to convert result %s to int", split[0])
		}

		partsStrings := strings.Split(strings.Trim(split[1], " "), " ")
		if len(partsStrings) < 2 {
			return equation{}, fmt.Errorf("invalid parts substring %s", split[1])
		}

		partsLen := len(partsStrings)
		parts := make([]float64, partsLen)
		for i, v := range partsStrings {
			part, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				panic(fmt.Errorf("invalid part %s", v))
			}
			parts[partsLen-i-1] = float64(part) // Reverse order to support left to right evaluation in tesEquation
		}

		return equation{result, parts}, nil
	})
}

type equation struct {
	result uint64
	parts  []float64
}

func stringifyEquation(parts []float64, ops []string) string {
	partsLen := len(parts)

	var builder strings.Builder
	for i := partsLen - 1; i >= 0; i-- {
		builder.WriteString(fmt.Sprintf("%d", uint64(parts[i])))
		if i > 0 {
			builder.WriteString(fmt.Sprintf(" %s ", ops[i-1]))
		}
	}

	return builder.String()
}

func evaluateEquation(eq string) (result uint64) {
	var conv uint64
	var num string
	var err error
	var op rune

	for _, r := range fmt.Sprintf("%s+", eq) {
		if r == ' ' || (r == '|' && op == '|') {
			continue
		} else if isDigit(r) {
			num += string(r)
		} else {
			if len(num) > 0 {
				conv, err = strconv.ParseUint(num, 10, 64)
				if err != nil {
					panic(fmt.Errorf("invalid number %s", num))
				}
				num = ""
			}

			if result == 0 || op == '|' {
				result = conv
			} else if op == '+' {
				result += conv
			} else if op == '*' {
				result *= conv
			}

			if r == '+' || r == '*' || r == '|' {
				op = r
				if r == '|' {
					num = fmt.Sprintf("%d", result)
				}
			}
		}
	}

	return
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
