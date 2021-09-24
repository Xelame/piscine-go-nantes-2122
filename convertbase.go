package piscine

func ConvertBase(number, baseFrom, baseTo string) string {
	result := []rune{}
	nbr := AtoiBase(number, baseFrom)
	if nbr == -9223372036854775808 {
		nbr = 223372036854775808
		result = append(result, '-')
		result = append(result, '9')
	} else if nbr < 0 {
		nbr *= -1
		result = append(result, '-')
	}
	var ints []int
	for nbr > 0 {
		ints = append(ints, nbr%len(baseTo))
		nbr = int(nbr / len(baseTo))
	}
	for i := len(ints) - 1; i > -1; i-- {
		result = append(result, rune(baseTo[ints[i]]))
	}
	return string(result)
}
