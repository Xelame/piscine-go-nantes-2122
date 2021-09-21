package piscine

func IterativePower(nb int, power int) int {
	nbmul := nb
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	for index := 1; index <= power-1; index++ {
		nb *= nbmul
	}
	return nb
}
