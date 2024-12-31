package day20

func CountLongerTopCheatsOnRacetrack() int {
	grid, s, e := readInput()
	track := trace(grid, s, e)
	ch := cheats(grid, s, e, 20, track)
	best := bestCheats(ch, 100)

	cheatStats(best)

	return len(best)
}
