package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	} else if nb == 4 {
		return 5
	}
	isPrime := true
	for diviser := nb - Sqrt(nb) - 1; diviser > 1 && isPrime; diviser-- {
		if nb%diviser == 0 {
			isPrime = false
		}
	}
	if isPrime {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}
