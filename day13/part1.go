package day13

func CalcTokensForAllPrizes() int {
	machines := readInput()

	tokens := 0
	for _, m := range machines {
		a, b := solve(m)
		if a == 0 || b == 0 {
			continue
		}
		tokens += a*3 + b
	}

	return tokens
}
