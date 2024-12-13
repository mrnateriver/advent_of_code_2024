package day05

func SumMiddlePagesOfCorrectlyOrderedUpdates() int {
	seenPages := [100]int{}
	i := 0

	return parseUpdates(func(upd update, rules [][]precedenceRule) int {
		i++

		if isUpdateValid(upd, rules, seenPages[:], i) {
			return getMiddlePage(upd)
		}

		return 0
	})
}

func isUpdateValid(upd update, rules [][]precedenceRule, seenPages []int, i int) bool {
	for _, v := range upd {
		pageRules := rules[v]
		if pageRules != nil {
			for _, rule := range pageRules {
				if seenPages[rule.beforePage] == i { // We're comparing to i because we don't want to reset seenPages after every search
					return false
				}
			}

		}
		seenPages[v] = i
	}

	return true
}
