package day24

import "fmt"

func RewireCircuit() string {
	// gates, values := readInput()

	testGates := buildRippleCarryAdder(4)

	result := performAddition(testGates, map[wire]int{
		"x0": 1,
		"x1": 1,
		"x2": 0,
		"x3": 1,

		"y0": 1,
		"y1": 0,
		"y2": 1,
		"y3": 1,
	})

	fmt.Printf("result: %v\n", result)

	return ""
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
		{x0, y0, abx, -1, -1, nil, nil, xorGate},
		{abx, c0, z0, -1, -1, nil, nil, xorGate},

		{abx, c0, abxcx, -1, -1, nil, nil, andGate},
		{x0, y0, aba, -1, -1, nil, nil, andGate},

		{abxcx, aba, c1, -1, -1, nil, nil, orGate},
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
