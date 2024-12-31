package day21

import (
	"fmt"
	"math"
	"mrnateriver.io/advent_of_code_2024/shared"
	"strconv"
	"strings"
)

func readInput() []string {
	codes := make([]string, 0, 4)
	for line := range shared.ReadInputLines("day21/input") {
		codes = append(codes, line)
	}
	return codes
}

func sumComplexities(codes []string, depth int) uint64 {
	var sum uint64 = 0
	for _, code := range codes {
		trimmed := strings.TrimPrefix(strings.TrimSuffix(code, "A"), "0")
		num, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(fmt.Errorf("failed to convert %s to int: %w", trimmed, err))
		}

		cost := findCost(code, depth)

		sum += uint64(num) * cost
	}
	return sum
}

func findCost(seq string, depth int) uint64 {
	return findCostRec(seq, depth, numpadPaths, make(map[solution]uint64))
}

func findCostRec(seq string, depth int, paths map[path][]string, dp map[solution]uint64) uint64 {
	if res, ok := dp[solution{seq, depth}]; ok {
		return res
	}

	var res uint64 = 0
	for i := 0; i < len(seq); i++ {
		var a, b byte
		if i == 0 {
			a, b = 'A', seq[i]
		} else {
			a, b = seq[i-1], seq[i]
		}
		pp := paths[path{a, b}]
		if depth == 0 {
			res += uint64(len(shortestString(pp)))
		} else {
			mm := uint64(math.MaxUint64)
			for _, p := range pp {
				v := findCostRec(p, depth-1, dirpadPaths, dp)
				if v < mm {
					mm = v
				}
			}
			res += mm
		}
	}

	dp[solution{seq, depth}] = res

	return res
}

type path struct {
	from, to byte
}

type solution struct {
	seq   string
	depth int
}

var numpad = [][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"#", "0", "A"},
}
var numpadEnd = map[byte]shared.Point2d{
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'0': {1, 3},
	'A': {2, 3},
}
var numpadPaths map[path][]string

var dirpad = [][]string{
	{"#", "^", "A"},
	{"<", "v", ">"},
}
var dirpadEnd = map[byte]shared.Point2d{
	'^': {1, 0},
	'v': {1, 1},
	'<': {0, 1},
	'>': {2, 1},
	'A': {2, 0},
}
var dirpadPaths map[path][]string

func shortestString(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	mm := strs[0]
	for _, s := range strs {
		if len(s) < len(mm) {
			mm = s
		}
	}
	return mm
}

func init() {
	numpadPaths = make(map[path][]string)
	for start := range numpadEnd {
		for end := range numpadEnd {
			s := numpadEnd[start]
			e := numpadEnd[end]
			numpadPaths[path{start, end}] = shortestPath(numpad, s, e)
		}
	}

	dirpadPaths = make(map[path][]string)
	for start := range dirpadEnd {
		for end := range dirpadEnd {
			s := dirpadEnd[start]
			e := dirpadEnd[end]
			dirpadPaths[path{start, end}] = shortestPath(dirpad, s, e)
		}
	}
}

func shortestPath(grid [][]string, s, e shared.Point2d) []string {
	paths := shared.FindShortestPaths(grid, s, e, "#")
	res := make([]string, 0, len(paths))
	for _, p := range paths {
		var path strings.Builder
		for i := 1; i < len(p.Points); i++ {
			a, b := p.Points[i-1], p.Points[i]
			dir := shared.GetDirection(a, b)
			dirChar := shared.DirChar(dir)

			path.WriteByte(dirChar)
		}
		path.WriteByte('A')

		res = append(res, path.String())
	}

	return res
}
