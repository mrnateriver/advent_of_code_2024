package day19

import (
	"iter"
	"mrnateriver.io/advent_of_code_2024/shared"
	"strings"
	"sync"
)

func readInput() (patterns map[string]struct{}, maxPrefixLen int, designs <-chan string) {
	patternsOut := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer close(patternsOut)

		readPatterns := false
		for line := range shared.ReadInputLines("day19/input") {
			if !readPatterns {
				patterns = groupedPatterns(strings.Split(line, ", "))
				readPatterns = true
				wg.Done()
				continue
			}
			if line == "" {
				continue
			}
			patternsOut <- line
		}
	}()

	wg.Wait()

	for pattern := range patterns {
		if len(pattern) > maxPrefixLen {
			maxPrefixLen = len(pattern)
		}
	}

	designs = patternsOut
	return
}

func prefixes(grouped map[string]struct{}, maxPrefixLen int, design string) iter.Seq[int] {
	return func(yield func(int) bool) {
		l := len(design)
		for i := 1; i <= maxPrefixLen && i <= l; i++ {
			if _, ok := grouped[design[:i]]; ok {
				if !yield(i) {
					return
				}
			}
		}
	}
}

func groupedPatterns(patterns []string) map[string]struct{} {
	grouped := make(map[string]struct{})
	for _, pattern := range patterns {
		grouped[pattern] = struct{}{}
	}
	return grouped
}
