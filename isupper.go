package piscine

func IsUpper(aString string) bool {
	isUpper := true
	for _, value := range aString {
		if !(65 <= int(value) && int(value) <= 90) {
			isUpper = !isUpper
		}
	}
	return isUpper
}
