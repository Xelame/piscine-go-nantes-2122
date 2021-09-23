package piscine

func AtoiBase(aString string, base string) int {
	if !TestBase(base) {
		powerOfBase := 1
		result := 0
		for index := len(aString) - 1; index >= 0; index-- {
			for digit, value := range base {
				if int(aString[index]) == int(value) {
					result += powerOfBase * digit
					powerOfBase *= len(base)
				}
			}
		}
		return result
	} else {
		return 0
	}
}
