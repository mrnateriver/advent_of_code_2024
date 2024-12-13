package day05

func SumMiddlePagesOfReorderedUpdates() (sum int) {
	seenPages := [100]seenIndex{}

	rules, updates := parseUpdates()

	i := 1
	for upd := range updates {
		i++
		if reordered := reorderInvalidUpdate(upd, rules, seenPages[:], &i); reordered != nil {
			sum += getMiddlePage(reordered)
		}
	}

	return
}

func reorderInvalidUpdate(upd update, rules [][]precedenceRule, seenPages []seenIndex, i *int) (reordered update) {
	// TODO: create a sorting map and sort updates according to it using Topological Sorting & Kahn's Algorithm
	// But since there are no performance requirements for the task, and the input is quite small, we can get away with O(n^2) complexity

repeat:
	for {
		for pageIdx, page := range upd {
			pageRules := rules[page]
			if pageRules != nil {
				for _, rule := range pageRules {
					if seen := seenPages[rule.beforePage]; seen.i == *i { // We're comparing to i because we don't want to reset seenPages after every search
						reordered = upd

						reordered[pageIdx] = rule.beforePage
						reordered[seen.idx] = page

						*i++
						continue repeat
					}
				}

			}
			seenPages[page] = seenIndex{*i, pageIdx}
		}
		break
	}

	return
}

type seenIndex struct {
	i   int
	idx int
}
