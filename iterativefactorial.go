package piscine

func IterativeFactorial(nb int) int {
	result := 0
	nbtemp := nb - 1
	for index := 1; index <= nbtemp; index++ {
		nb *= index
		result = nb
	}
	return result
}
