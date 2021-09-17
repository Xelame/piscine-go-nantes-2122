package piscine

func Atoi(aString string) int {
	powerOfTen := 1
	result := 0
	for index := range aString {
		if 48 <= aString[len(aString)-1-index] && aString[len(aString)-1-index] <= 57 {
			result += powerOfTen * int(aString[len(aString)-1-index]%48)
			powerOfTen *= 10
		} else if (aString[len(aString)-1-index] == 45 || aString[len(aString)-1-index] == 43) && len(aString)-1-index == 0 {
			powerOfTen *= 10
		} else {
			return 0
		}
	}
	if aString[0] == 45 {
		result = -result
	}
	return result
}
