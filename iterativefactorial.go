package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for index := 1; 9223372036854775806 > result && index <= nb; index++ {
		result *= index
	}
	if result == 1 || nb > 9223372036854775806 {
		return 0
	}
	return result
}
