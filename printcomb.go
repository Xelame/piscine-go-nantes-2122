package piscine

import "github.com/01-edu/z01"

/*
Function to print all unique combinations of three different digit in ascending order
Output : Suit of answers separate by a comma
*/
func PrintComb() {
	// Generator for our three different digits
	for hundreds := '0'; hundreds <= '9'; hundreds++ {
		for tens := '0'; tens <= '9'; tens++ {
			for units := '0'; units <= '9'; units++ {
				// Testing if they're in ascending order
				if units > tens && tens > hundreds {
					z01.PrintRune(hundreds)
					z01.PrintRune(tens)
					z01.PrintRune(units)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
