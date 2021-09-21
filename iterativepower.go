package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power <= 1 {
		return nb
	} else {
		return nb * IterativePower(nb, power-1)
	}
}
