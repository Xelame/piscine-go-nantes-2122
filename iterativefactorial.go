package piscine

func IterativeFactorial(nb int) int {
	result := 0
	nb2 := nb - 1
	for index := nb2; 9223372036854775807 > index && index > 0; index-- {
		nb *= index
		result = nb
	}
	return result
}
