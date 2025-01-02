package day23

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"slices"
	"strings"
)

func readInput() (unique []addr, graph network) {
	unique = make([]addr, 0, 8)
	graph = make(network)
	for line := range shared.ReadInputLines("day23/input") {
		split := strings.Split(line, "-")
		if len(split) != 2 {
			panic(fmt.Errorf("invalid input line: %v", line))
		}

		a := addr{split[0][0], split[0][1]}
		b := addr{split[1][0], split[1][1]}
		if _, exists := graph[a]; !exists {
			graph[a] = make(map[addr]struct{})
			unique = append(unique, a)
		}
		if _, exists := graph[b]; !exists {
			graph[b] = make(map[addr]struct{})
			unique = append(unique, b)
		}
		graph[a][b] = struct{}{}
		graph[b][a] = struct{}{}
	}
	return
}

func clusters(graph network, start addr, nodesLimit int, dp map[cluster][][]addr) [][]addr {
	if nodesLimit == 0 {
		return [][]addr{}
	}

	if res, ok := dp[cluster{start, nodesLimit}]; ok {
		return res
	}

	connections, exists := graph[start]
	if !exists || nodesLimit == 1 {
		res := [][]addr{{start}}
		dp[cluster{start, nodesLimit}] = res
		return res
	}

	res := make([][]addr, 0, len(connections))

	if nodesLimit == 2 {
		for c := range connections {
			res = append(res, []addr{start, c})
		}
	} else {
		for c := range connections {
			smallerCluster := clusters(graph, c, nodesLimit-1, dp)
		cluster:
			for _, sc := range smallerCluster {
				for _, node := range sc {
					if !connected(graph, start, node) {
						continue cluster
					}
				}
				res = append(res, slices.Concat([]addr{start}, sc))
			}
		}
	}

	dp[cluster{start, nodesLimit}] = res
	return res
}

func connected(graph network, a, b addr) bool {
	_, exists1 := graph[a][b]
	_, exists2 := graph[b][a]
	return exists1 && exists2
}

func dedupe(clusters [][]addr) [][]addr {
	res := make([][]addr, 0, len(clusters))
	dp := make(map[string]struct{})
	for _, cl := range clusters {
		if len(cl) == 0 {
			continue
		}

		inted := make([]uint16, 0, len(cl))
		for _, a := range cl {
			inted = append(inted, addrToUint(a))
		}
		slices.Sort(inted)

		var sb strings.Builder
		for _, i := range inted {
			sb.WriteString(fmt.Sprintf("%d", i))
		}

		key := sb.String()
		if _, exists := dp[key]; !exists {
			res = append(res, cl)
			dp[key] = struct{}{}
		}
	}
	return res
}

func addrToUint(a addr) uint16 {
	return uint16(a[0])<<8 + uint16(a[1])
}

type network map[addr]map[addr]struct{}

type addr [2]byte

type cluster struct {
	addr addr
	size int
}
