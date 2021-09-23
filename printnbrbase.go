/*package piscine

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
	for i := len(base) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if (base[i] == base[j]) && base[j] != 45 && base[j] != 43 && isOkay {
				z01.PrintRune('N')
				z01.PrintRune('V')
				isOkay = false
			}
		}
	}
	return isOkay
}
*/

package piscine

import (
	"github.com/01-edu/z01"
)

func TestBase(s string) bool {
	runes := []rune(s)
	if runes[len(runes)-1] == '-' {
		return true
	}
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] || runes[i] == '-' {
				return true
			}
		}
	}
	return false
}

func PrintNbrBase(nbr int, base string) {
	if len(base) <= 1 || (TestBase(base)) {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		if nbr == -9223372036854775808 {
			nbr = 223372036854775808
			z01.PrintRune('-')
		} else if nbr < 0 {
			nbr *= -1
			z01.PrintRune('-')
		}
		var ints []int
		for nbr > 0 {
			ints = append(ints, nbr%len(base))
			nbr = int(nbr / len(base))
		}
		for i := len(ints) - 1; i > -1; i-- {
			z01.PrintRune(rune(base[ints[i]]))
		}
	}
}
