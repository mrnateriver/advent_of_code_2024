package day24

import (
	"fmt"
	"strconv"
)

func RewireCircuit() string {
	gates, values := readInput()
	groupedGates := groupGates(gates)
	connectGates(groupedGates)

	maxI := 0
	for w := range values {
		if w[0] == 'x' {
			x, err := strconv.Atoi(string(w[1:]))
			if err != nil {
				panic(fmt.Errorf("failed to parse x wire %v: %w", w, err))
			}
			if x > maxI {
				maxI = x
			}
		}
	}

	//res := make([]wire, 0, 8)
	for i := 0; i <= maxI; i++ {
		if !validFullAdderExists(groupedGates, i) {
			fmt.Printf("failed to find full adder for %v\n", i)
		}
	}

	return ""
}

func validFullAdderExists(groupedGates map[wire][]*gateDesc, i int) bool {
	x := wire(fmt.Sprintf("x%02d", i))
	y := wire(fmt.Sprintf("y%02d", i))
	z := wire(fmt.Sprintf("z%02d", i))

	abx := findAbxGate(groupedGates, x, y)
	aba := findAbaGate(groupedGates, x, y)
	if abx == nil || aba == nil {
		panic(fmt.Errorf("inputs are not connected for bit %v", i))
	}

	if z == "z00" {
		return true
	}

	var c0, abxcx, c1 *gateDesc

	c0 = findC0Gate(groupedGates, abx.o, z)
	if c0 != nil {
		abxcx = findAbxcxGate(groupedGates, abx.o, c0.b)
	}
	if abxcx != nil {
		c1 = findC1Gate(groupedGates, abxcx.o, aba.o)
	}

	return c0 != nil && abxcx != nil && c1 != nil
}

func findAbxGate(groupedGates map[wire][]*gateDesc, x, y wire) *gateDesc {
	for _, g := range groupedGates[x] {
		if (g.a == y || g.b == y) && g.kind == gateKindXor {
			return g
		}
	}
	return nil
}

func findAbaGate(groupedGates map[wire][]*gateDesc, x, y wire) *gateDesc {
	for _, g := range groupedGates[x] {
		if (g.a == y || g.b == y) && g.kind == gateKindAnd {
			return g
		}
	}
	return nil
}

func findC0Gate(groupedGates map[wire][]*gateDesc, abx, z wire) *gateDesc {
	for _, g := range groupedGates[abx] {
		if g.o == z && g.kind == gateKindXor {
			return g
		}
	}
	return nil
}

func findAbxcxGate(groupedGates map[wire][]*gateDesc, abx, c0 wire) *gateDesc {
	for _, g := range groupedGates[abx] {
		if (g.a == c0 || g.b == c0) && g.kind == gateKindAnd {
			return g
		}
	}
	return nil
}

func findC1Gate(groupedGates map[wire][]*gateDesc, abxcx, aba wire) *gateDesc {
	for _, g := range groupedGates[abxcx] {
		if (g.a == aba || g.b == aba) && g.kind == gateKindOr {
			return g
		}
	}
	return nil
}

func swapOutputs(a, b *gateDesc) {
	a.o, b.o = b.o, a.o
	a.next, b.next = b.next, a.next
}

// https://en.wikipedia.org/wiki/Adder_(electronics)#Full_adder
func buildFullAdder(i int) (res, c0gates []*gateDesc, c1gate *gateDesc) {
	x0 := wire(fmt.Sprintf("x%v", i))
	y0 := wire(fmt.Sprintf("y%v", i))
	c0 := wire(fmt.Sprintf("c%v", i))
	z0 := wire(fmt.Sprintf("z%v", i))
	c1 := wire(fmt.Sprintf("c%v", i+1))
	abx := wire(fmt.Sprintf("abx%v", i))
	abxcx := wire(fmt.Sprintf("abxcx%v", i))
	aba := wire(fmt.Sprintf("aba%v", i))

	res = []*gateDesc{
		{gateKindXor, x0, y0, abx, -1, -1, nil, nil, xorGate},
		{gateKindXor, abx, c0, z0, -1, -1, nil, nil, xorGate},

		{gateKindAnd, abx, c0, abxcx, -1, -1, nil, nil, andGate},
		{gateKindAnd, x0, y0, aba, -1, -1, nil, nil, andGate},

		{gateKindOr, abxcx, aba, c1, -1, -1, nil, nil, orGate},
	}

	res[0].next = []*gateDesc{res[1], res[2]}
	res[2].next = []*gateDesc{res[4]}
	res[3].next = []*gateDesc{res[4]}

	c0gates = []*gateDesc{res[1], res[2]}
	c1gate = res[4]
	return
}

func buildRippleCarryAdder(bits int) (res []*gateDesc) {
	res = make([]*gateDesc, 0, (bits+1)*5)

	var gates, c0s []*gateDesc
	var c1, newC1 *gateDesc
	for i := 0; i < bits; i++ {
		gates, c0s, newC1 = buildFullAdder(i)
		if c1 != nil {
			c1.next = c0s
		} else {
			for _, c0 := range c0s {
				updateGate(c0, "c0", 0)
			}
		}
		c1 = newC1

		res = append(res, gates...)
	}

	c1.o = wire(fmt.Sprintf("z%v", bits))
	return
}
