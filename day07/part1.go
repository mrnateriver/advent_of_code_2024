package day07

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func SumPossibleEquationsResults() uint64 {
	input := shared.ReadInputLines("day07/input")

	totalFound := 0
	totalInvalid := 0

	var sum uint64
	for line := range input {
		split := strings.Split(line, ":")
		if len(split) != 2 {
			panic(fmt.Errorf("invalid line %s", line))
		}

		result, err := strconv.ParseUint(split[0], 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to convert result %s to int", split[0]))
		}

		partsStrings := strings.Split(strings.Trim(split[1], " "), " ")
		if len(partsStrings) < 2 {
			panic(fmt.Errorf("invalid parts substring %s", split[1]))
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

		ops := make([]string, partsLen-1)
		if testEquation(float64(result), parts, ops) {
			totalFound++

			str := stringifyEquation(parts, ops)
			eval := evaluateEquation(str)
			if eval != result {
				totalInvalid++
				log.Println("Invalid equation!")
				log.Printf("Valid: %d = %s", result, str)
				log.Printf("Evaluated: %d == %d", result, eval)
			}

			sum += result
		}
	}

	log.Printf("Total found: %d", totalFound)
	log.Printf("Total invalid: %d", totalInvalid)

	return sum
}

func testEquation(result float64, b []float64, ops []string) bool {
	if len(b) == 0 {
		return false
	}
	if len(b) == 1 {
		return result == b[0]
	}

	got := testEquation(result-b[0], b[1:], ops[1:])
	if got {
		ops[0] = "+"
		return true
	}

	if b[0] > 0 {
		ops[0] = "*"
		return testEquation(result/b[0], b[1:], ops[1:])
	}

	return false
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
		if r == ' ' {
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

			if result == 0 {
				result = conv
			} else if op == '+' {
				result += conv
			} else if op == '*' {
				result *= conv
			}

			if r == '+' || r == '*' {
				op = r
			}
		}
	}

	return
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
