package piscine

import "github.com/01-edu/z01"

func PrintNbr(number int) {
	if number == -9223372036854775808 {
		number = 223372036854775808
		z01.PrintRune('-')
		z01.PrintRune('9')
	}
	if number < 0 {
		z01.PrintRune('-')
		number = -number
	}
	var reversedList []int
	for number >= 10 {
		reversedList = append(reversedList, number%10)
		number /= 10
	}
	reversedList = append(reversedList, number%10)
	for end := len(reversedList) - 1; end >= 0; end-- {
		z01.PrintRune(rune(48 + reversedList[end]))
	}
}
