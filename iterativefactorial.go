package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for index := 1; index <= nb && nb <= 20; index++ {
		result *= index
		result = nb
	}
	if result == 1 {
		return 0
	}
	return result
}
