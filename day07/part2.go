package day07

import (
	"fmt"
	"strconv"
	"strings"
)

func SumPossibleEquationsResultsWithConcatenation() uint64 {
	var sum uint64
	for equation := range readInput() {
		partsLen := len(equation.parts)
		ops := make([]string, partsLen-1)
		if testEquationWithConcatenation(float64(equation.result), equation.parts, ops) {
			sum += equation.result
		}
	}

	return sum
}

func testEquationWithConcatenation(result float64, b []float64, ops []string) bool {
	if len(b) == 0 {
		return false
	}
	if len(b) == 1 {
		return result == b[0]
	}

	got := testEquationWithConcatenation(result-b[0], b[1:], ops[1:])
	if got {
		ops[0] = "+"
		return true
	}

	if b[0] > 0 {
		got = testEquationWithConcatenation(result/b[0], b[1:], ops[1:])
		if got {
			ops[0] = "*"
			return true
		}
	}

	// Short-circuit if result is fractional, since there's no way to get it via concatenation
	if result != float64(uint64(result)) {
		return false
	}

	resultStr := fmt.Sprintf("%d", uint64(result))
	b0Str := fmt.Sprintf("%d", uint64(b[0]))
	resultCut, found := strings.CutSuffix(resultStr, b0Str)

	// If the next number is exactly the same as the result, there's also no way to achieve result with concatenation
	if !found || resultCut == "" {
		return false
	}

	resultConv, err := strconv.ParseFloat(resultCut, 64)
	if err != nil {
		panic(fmt.Errorf("invalid number %s", resultCut))
	}

	got = testEquationWithConcatenation(resultConv, b[1:], ops[1:])
	if got {
		ops[0] = "||"
		return true
	}

	return false
}
