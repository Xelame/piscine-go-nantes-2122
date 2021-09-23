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
