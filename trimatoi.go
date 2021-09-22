package piscine

func TrimAtoi(aString string) int {
	result := 0
	if aString != "" {
		limits := 0
		powerOfTen := 1
		for index := len(aString) - 1; index >= 0; index-- {
			if 48 <= aString[index] && aString[index] <= 57 {
				limits = index
				result += powerOfTen * int(aString[index]-48)
				powerOfTen *= 10
			}
		}
		for search := limits; search >= 0; search-- {
			if aString[search] == 45 {
				return -result
			}
		}
	}
	return result
}
