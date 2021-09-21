package piscine

func FindNextPrime(nb int) int {
	if nb <= 0 {
		return 2
	}
	countDivideBy := 0
	for diviser := 1; diviser <= nb && countDivideBy <= 2; diviser++ {
		if nb%diviser == 0 {
			countDivideBy += 1
		}
	}
	if countDivideBy == 2 {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}
