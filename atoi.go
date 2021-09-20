package piscine

func Atoi(aString string) int {
	if aString != "" {
		powerOfTen := 1
		result := 0
		for index := len(aString) - 1; index >= 0; index-- {
			if 48 <= aString[index] && aString[index] <= 57 {
				result += powerOfTen * int(aString[index]%48)
				powerOfTen *= 10
			} else if (aString[index] == 45 || aString[index] == 43) && index == 0 {
				powerOfTen *= 10
			} else {
				return 0
			}
		}
		if aString[0] == 45 {
			result = -result
		}
		return result
	} else {
		return 0
	}
}
