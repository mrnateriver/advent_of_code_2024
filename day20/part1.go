package day20

func CountTopCheatsOnRacetrack() int {
	grid, s, e := readInput()
	track := trace(grid, s, e)
	ch := cheats(grid, s, e, 2, track)
	best := bestCheats(ch, 100)

	cheatStats(best)

	return len(best)
}
