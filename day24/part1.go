package day24

func RunWiring() uint64 {
	gates, values := readInput()

	return performAddition(gates, values)
}
