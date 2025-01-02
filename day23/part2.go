package day23

import (
	"mrnateriver.io/advent_of_code_2024/shared"
	"slices"
	"strings"
)

func FindLargestCluster() string {
	unique, graph := readInput()

	p := make(map[addr]struct{})
	for _, u := range unique {
		p[u] = struct{}{}
	}

	clique := largestClique(graph, p, map[addr]struct{}{}, map[addr]struct{}{})
	if clique == nil {
		return ""
	}

	res := shared.Keys(clique)

	strRes := make([]string, 0, len(res))
	for _, r := range res {
		strRes = append(strRes, string([]byte{r[0], r[1]}))
	}

	slices.SortFunc(strRes, func(a, b string) int {
		return strings.Compare(a, b)
	})

	return strings.Join(strRes, ",")
}

// Bron-Kerbosch algorithm for finding the largest clique in a graph
// https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
func largestClique(graph network, p, r, x map[addr]struct{}) map[addr]struct{} {
	if len(p) == 0 && len(x) == 0 {
		return r
	}

	var withMostNeighbors *addr
	for k := range shared.Union(p, x) {
		if withMostNeighbors == nil || len(graph[*withMostNeighbors]) < len(graph[k]) {
			withMostNeighbors = &k
		}
	}
	if withMostNeighbors == nil {
		return make(map[addr]struct{})
	}

	var res map[addr]struct{}
	withMostNeighborsConnections := graph[*withMostNeighbors]
	for v := range p {
		if _, exists := withMostNeighborsConnections[v]; exists {
			continue
		}

		connections := graph[v]

		newP := shared.Intersect(p, connections)
		newR := shared.Union(r, map[addr]struct{}{v: {}})
		newX := shared.Intersect(x, connections)

		clique := largestClique(graph, newP, newR, newX)
		if res == nil || len(clique) > len(res) {
			res = clique
		}
	}

	return res
}
