package piscine

func IterativeFactorial(nb int) int {
	if nb > 1000000000000000000 {
		return 0
	}
	result := 1
	for index := 1; 1000000000000000000 > result && index <= nb; index++ {
		result *= index
	}
	if result == 1 {
		return 0
	}
	return result
}
