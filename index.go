package piscine

func Index(aString string, toFind string) int {
	for i := 0; i <= len(aString)-len(toFind); i++ {
		if toFind == aString[i:i+len(toFind)] {
			return i
		}
	}
	return -1
}
