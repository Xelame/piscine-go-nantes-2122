package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power <= 1 {
		return 1
	} else {
		return nb * IterativePower(nb, power-1)
	}
}
