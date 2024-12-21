package day13

func CalcTokensForAllPrizesIncreased() int {
	machines := readInput()

	tokens := 0
	for _, m := range machines {
		m.tx += 10000000000000
		m.ty += 10000000000000

		a, b := solve(m)
		if a == 0 || b == 0 {
			continue
		}
		tokens += a*3 + b
	}

	return tokens
}
