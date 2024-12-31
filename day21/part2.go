package day21

func CalcCodesComplexityLarge() uint64 {
	codes := readInput()
	return sumComplexities(codes, 25)
}
