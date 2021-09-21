package piscine

func IterativeFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	}
	result := 1
	for index := 1; index <= nb; index++ {
		result *= index
	}
	return result
}
