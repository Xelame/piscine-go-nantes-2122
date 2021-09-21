package piscine

func IterativeFactorial(nb int) int {
	result := nb
	for index := nb - 1; index > 0; index-- {
		result *= index
	}
	if result < 0 {
		return 0
	}
	return result
}
