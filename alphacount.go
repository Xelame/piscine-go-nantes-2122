package piscine

func AlphaCount(aString string) int {
	countAlpha := 0
	for _, value := range aString {
		if (65 <= int(value) && int(value) <= 90) || (97 <= int(value) && int(value) <= 122) {
			countAlpha++
		}
	}
	return countAlpha
}
