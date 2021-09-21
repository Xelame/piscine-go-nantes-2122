package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 20 || nb < 0 {
		return 0
	}
	for index := 1; index <= nb; index++ {
		result *= index
		result = nb
	}
	return result
}
