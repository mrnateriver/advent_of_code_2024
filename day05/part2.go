package day05

func SumMiddlePagesOfReorderedUpdates() int {
	seenPages := [100]seenIndex{}
	i := 1

	return parseUpdates(func(upd update, rules [][]precedenceRule) int {
		i++

		if reordered := reorderInvalidUpdate(upd, rules, seenPages[:], &i); reordered != nil {
			return getMiddlePage(reordered)
		}

		return 0
	})
}

func reorderInvalidUpdate(upd update, rules [][]precedenceRule, seenPages []seenIndex, i *int) (reordered update) {
	// TODO: create a list of page numbers which is sorted according to the rules; then somehow move the numbers from an update according to positions of elements in that array -- perhaps filter the sorted array for update numbers and compact it down to the same length as the update

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
