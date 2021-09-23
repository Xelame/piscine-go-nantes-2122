package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(number int, base string) {
	if TestBase(base) {
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
		for number > 0 {
			reversedList = append(reversedList, number%len(base))
			number /= len(base)
		}
		for end := len(reversedList) - 1; end >= 0; end-- {
			z01.PrintRune(rune(base[reversedList[end]]))
		}
	}
}

func TestBase(base string) bool {
	isOkay := true
	for i := 0; i < len(base)-1 && isOkay; i++ {
		for j := i + 1; j < len(base) && isOkay; j++ {
			if base[i] == base[j] || base[i] == 45 {
				z01.PrintRune('N')
				z01.PrintRune('V')
				isOkay = false
			}
		}
	}
	return isOkay
}
