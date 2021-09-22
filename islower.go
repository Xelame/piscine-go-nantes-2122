package piscine

func IsLower(aString string) bool {
	isLower := true
	for _, value := range aString {
		if !(97 <= int(value) && int(value) <= 122) && isLower {
			isLower = !isLower
		}
	}
	return isLower
}
