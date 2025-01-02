package day24

import (
	"fmt"
	"mrnateriver.io/advent_of_code_2024/shared"
	"regexp"
	"strconv"
	"strings"
)

func RunWiring() uint64 {
	initRegex := regexp.MustCompile("(\\w+): (\\d+)")
	gateRegex := regexp.MustCompile("(\\w+) (AND|OR|XOR) (\\w+) -> (\\w+)")

	values := make(map[wire]int)
	groupedVertices := make(map[wire][]*gateDesc)

	for line := range shared.ReadInputLines("day24/input") {
		if line == "" {
			continue
		}

		matches := initRegex.FindStringSubmatch(line)
		if matches != nil {
			w := matches[1]
			val, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(fmt.Errorf("invalid value %v for wire %v: %w", matches[2], w, err))
			}

			values[wire(w)] = val
			continue
		}

		matches = gateRegex.FindStringSubmatch(line)
		if matches != nil {
			a := matches[1]
			g := parseGate(matches[2])
			b := matches[3]
			o := matches[4]

			vertex := gateDesc{wire(a), wire(b), wire(o), -1, -1, nil, nil, g}
			storeVertex(groupedVertices, &vertex)
			continue
		}
	}

	for _, gates := range groupedVertices {
		for _, g := range gates {
			if strings.HasPrefix(string(g.o), "z") {
				continue
			}
			g.next = groupedVertices[g.o]
		}
	}

	var output uint64
	for w, v := range values {
		for _, g := range groupedVertices[w] {
			output |= propagateSignal(g, w, v)
		}
	}

	return output
}

func propagateSignal(gate *gateDesc, w wire, val int) (outputMask uint64) {
	if gate.vo != nil {
		return *gate.vo
	}

	updateGate(gate, w, val)
	if gateComplete(gate) {
		nextVal := gate.g(gate.va, gate.vb)
		if gate.next == nil || len(gate.next) == 0 {
			outputMask = outputBit(gate.o, uint64(nextVal))
		} else {
			for _, next := range gate.next {
				outputMask |= propagateSignal(next, gate.o, nextVal)
			}
		}

		gate.vo = &outputMask
	}

	return
}

func gateComplete(gate *gateDesc) bool {
	return gate.va >= 0 && gate.vb >= 0
}

func updateGate(gate *gateDesc, w wire, val int) {
	if gate.a == w {
		gate.va = val
	} else {
		gate.vb = val
	}
}

func outputBit(w wire, v uint64) uint64 {
	if !strings.HasPrefix(string(w), "z") {
		panic(fmt.Errorf("invalid output wire %v", w))
	}

	bit, err := strconv.Atoi(strings.TrimPrefix(string(w[1:]), "0"))
	if err != nil {
		panic(fmt.Errorf("failed to parse output bit %v: %w", w, err))
	}

	return v << bit
}

func storeVertex(vertices map[wire][]*gateDesc, v *gateDesc) {
	if _, exists := vertices[v.a]; !exists {
		vertices[v.a] = make([]*gateDesc, 0, 2)
	}
	vertices[v.a] = append(vertices[v.a], v)

	if _, exists := vertices[v.b]; !exists {
		vertices[v.b] = make([]*gateDesc, 0, 2)
	}
	vertices[v.b] = append(vertices[v.b], v)
}

func parseGate(g string) gate {
	switch g {
	case "AND":
		return andGate
	case "OR":
		return orGate
	case "XOR":
		return xorGate
	}
	panic(fmt.Errorf("unknown gate: %v", g))
}

func andGate(a, b int) int {
	return a & b
}

func orGate(a, b int) int {
	return a | b
}

func xorGate(a, b int) int {
	return a ^ b
}

type gate func(int, int) int

type wire string

type gateDesc struct {
	a, b, o wire
	va, vb  int
	vo      *uint64
	next    []*gateDesc
	g       gate
}
