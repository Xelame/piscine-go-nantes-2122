package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	countDivideBy := 0
	for diviser := 2; diviser < nb && countDivideBy <= 2; diviser++ {
		if nb%diviser == 0 {
			countDivideBy += 1
		}
	}
	if countDivideBy == 0 {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}
