package day19

func CountPossibleDesigns() int {
	patterns, maxPrefixLen, designs := readInput()

	count := 0
	for design := range designs {
		if patternPossible(patterns, maxPrefixLen, design) {
			count++
		}
	}

	return count
}

func patternPossible(patterns map[string]struct{}, maxPrefixLen int, design string) bool {
	if design == "" {
		return true
	}

	for l := range prefixes(patterns, maxPrefixLen, design) {
		res := patternPossible(patterns, maxPrefixLen, design[l:])
		if res {
			return true
		}
	}

	return false
}
