package piscine

func IsNumeric(aString string) bool {
	isNum := true
	for _, value := range aString {
		if !(48 <= int(value) && int(value) <= 57) && isNum {
			isNum = !isNum
		}
	}
	return isNum
}
