package piscine

func Capitalize(aString string) string {
	aString = ToLower(aString)
	listOfCharacter := []rune(aString)
	if IsLower(string(listOfCharacter[0])) {
		listOfCharacter[0] -= 32
	}
	for i := 1; i < len(listOfCharacter); i++ {
		if (64 >= listOfCharacter[i-1] && listOfCharacter[i-1] >= 58) || (47 >= listOfCharacter[i-1] && listOfCharacter[i-1] >= 32) || (96 >= listOfCharacter[i-1] && listOfCharacter[i-1] >= 91) || (177 >= listOfCharacter[i-1] && listOfCharacter[i-1] >= 123) {
			if IsLower(string(listOfCharacter[i])) {
				listOfCharacter[i] -= 32
			}
		}
	}
	return string(listOfCharacter)
}
