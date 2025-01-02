package day23

func CountSetsOfComputerTrios() int {
	unique, graph := readInput()

	res := make([][]addr, 0, 8)

	dp := make(map[cluster][][]addr)
	for _, a := range unique {
		c := clusters(graph, a, 3, dp)
		for _, cl := range c {
			for _, node := range cl {
				if node[0] == 't' {
					res = append(res, cl)
					break
				}
			}
		}
	}

	res = dedupe(res)

	return len(res)
}
