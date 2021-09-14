package piscine

import "github.com/01-edu/z01"

/*
Function who returned T(rue) or F(alse) in function of parameter if he is positive or negative
Parameter : integrer - nb
Output : rune - T or F
*/

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('F')
		z01.PrintRune('\n')
	}
}
