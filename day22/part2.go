package day22

import (
	"iter"
)

func CalcMaxBananasFromSequences() uint64 {
	const depth = 2001

	dp := make(map[seq]uint64)
	for secretNumber := range readInput() {
		local := make(map[seq]uint64)
		for s, p := range changeSequences(secretNumber, depth) {
			if _, ok := local[s]; !ok {
				// We can only use the first occurrence of a sequence
				local[s] = p
			}
		}

		for s, p := range local {
			if e, ok := dp[s]; ok {
				dp[s] = e + p
			} else {
				dp[s] = p
			}
		}
	}

	m := uint64(0)
	for _, v := range dp {
		if v > m {
			m = v
		}
	}

	return m
}

func changeSequences(start uint64, depth int) iter.Seq2[seq, uint64] {
	return func(yield func(seq, uint64) bool) {
		s, si := [4]int8{}, 0
		price := buyerPrice(start)
		for d := 0; d < depth; d++ {
			secretNumber := calcNthSecretNumber(start, 1)
			curPrice := buyerPrice(secretNumber)
			if si < 4 {
				s[si] = int8(curPrice - price)
				si++
			} else {
				s[0] = s[1]
				s[1] = s[2]
				s[2] = s[3]
				s[3] = int8(curPrice - price)
			}

			if si >= 4 && !yield(s, curPrice) {
				return
			}

			price = curPrice
			start = secretNumber
		}
	}
}

func buyerPrice(num uint64) uint64 {
	return num % 10
}

type seq [4]int8
