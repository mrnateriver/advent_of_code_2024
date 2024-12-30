package day19

func CountWaysToCreateDesigns() int {
	grouped, maxPrefixLen, designs := readInput()

	dp := make(map[string]int)

	count := 0
	for design := range designs {
		count += countPossibleDesigns(grouped, maxPrefixLen, design, dp)
	}

	return count
}

func countPossibleDesigns(grouped map[string]struct{}, maxPrefixLen int, design string, dp map[string]int) (res int) {
	if design == "" {
		return 1
	}

	if res, ok := dp[design]; ok {
		return res
	}

	for l := range prefixes(grouped, maxPrefixLen, design) {
		res += countPossibleDesigns(grouped, maxPrefixLen, design[l:], dp)
	}

	dp[design] = res
	return
}
