package day05

func SumMiddlePagesOfCorrectlyOrderedUpdates() (sum int) {
	seenPages := [100]int{}

	rules, updates := parseUpdates()

	i := 0
	for upd := range updates {
		i++
		if isUpdateValid(upd, rules, seenPages[:], i) {
			sum += getMiddlePage(upd)
		}
	}

	return
}
