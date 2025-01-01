package day22

func CalcSecretNumbersSum() uint64 {
	const depth = 2000

	sum := uint64(0)
	for secretNumber := range readInput() {
		nth := calcNthSecretNumber(uint64(secretNumber), depth)
		sum += nth
	}

	return sum
}
