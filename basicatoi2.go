package piscine

func BasicAtoi2(aString string) int {
	powerOfTen := 1
	result := 0
	for index := range aString {
		if 48 <= aString[len(aString)-1-index] && aString[len(aString)-1-index] <= 57 {
			result += powerOfTen * int(aString[len(aString)-1-index]%48)
			powerOfTen *= 10
		} else {
			return 0
		}
	}
	return result
}
