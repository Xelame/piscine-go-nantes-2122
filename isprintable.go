package piscine

func IsPrintable(aString string) bool {
	isPrint := true
	for _, value := range aString {
		if (0 <= int(value) && int(value) <= 30) && isPrint {
			isPrint = !isPrint
		}
	}
	return isPrint
}
