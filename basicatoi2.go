package piscine

func BasicAtoi2(aString string) int {
	powerOfTen := 1
	result := 0
	for index := len(aString) - 1; index > 0; index-- {
		if 48 <= aString[index] && aString[index] <= 57 {
			result += powerOfTen * int(aString[index]%48)
			powerOfTen *= 10
		} else {
			return 0
		}
	}
	return result
}
