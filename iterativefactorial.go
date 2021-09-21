package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for index := 1; 184467440737095516 > result && index <= nb; index++ {
		result *= index
	}
	if result == 1 || nb > 184467440737095516 {
		return 0
	}
	return result
}
