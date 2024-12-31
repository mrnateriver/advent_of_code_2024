package day21

func CalcCodesComplexity() uint64 {
	codes := readInput()
	return sumComplexities(codes, 2)
}
