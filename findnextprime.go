package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	isPrime := true
	for diviser := 2; diviser*diviser <= nb && isPrime; diviser++ {
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
