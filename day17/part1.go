package day17

func RunProgram() string {
	cpu := readInput()
	debug(&cpu)

	return run(&cpu, true)
}

