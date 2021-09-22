package piscine

func ToUpper(aString string) string {
	listOfCharacter := []rune(aString)
	for index, value := range listOfCharacter {
		if 97 <= value && value <= 122 {
			listOfCharacter[index] -= 32
		}
	}
	return string(listOfCharacter)
}
