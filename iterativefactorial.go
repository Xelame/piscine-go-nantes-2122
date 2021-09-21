package piscine

func IterativeFactorial(nb int) int {
	result := 0
	for index := nb - 1; index > 0; index-- {
		nb *= index
		result = nb
	}
	return result
}
