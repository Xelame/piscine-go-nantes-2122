package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for index := 1; index <= nb; index++ {
		result *= index
	}
	if result == 1 {
		return 0
	}
	return result
}
