package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for index := nb; index > 0; index-- {
		result *= index
	}
	return result
}
