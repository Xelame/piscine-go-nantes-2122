package piscine

func IsAlpha(aString string) bool {
	isAlpha := true
	for _, value := range aString {
		if ((65 <= int(value) && int(value) <= 90) || (97 <= int(value) && int(value) <= 122)) && isAlpha {
			isAlpha = !isAlpha
		}
	}
	return isAlpha
}
