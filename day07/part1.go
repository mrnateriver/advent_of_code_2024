package day07

func SumPossibleEquationsResults() uint64 {
	var sum uint64
	for equation := range readInput() {
		partsLen := len(equation.parts)
		ops := make([]string, partsLen-1)
		if testEquation(float64(equation.result), equation.parts, ops) {
			sum += equation.result
		}
	}

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
