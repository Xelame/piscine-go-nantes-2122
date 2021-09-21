package piscine

func FindNextPrime(nb int) int {
	if nb < 0 {
		return 2
	}
	result := 0
	countDivideBy := 0
	var listOfPrimeNumber []int
	for number := 1; number < 100; number++ {
		countDivideBy = 0
		for diviser := 1; diviser <= number; diviser++ {
			if number%diviser == 0 {
				countDivideBy += 1
			}
		}
		if countDivideBy == 2 {
			listOfPrimeNumber = append(listOfPrimeNumber, number)
		}
	}
	for _, i := range listOfPrimeNumber {
		if i <= nb {
			result = i
		}
	}
	return result
}
