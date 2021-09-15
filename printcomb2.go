package piscine

import "github.com/01-edu/z01"

/*
Function to print all unique combinations of three different digit in ascending order
Output : Suit of answers separate by a comma
*/
func PrintComb2() {
	// Generator for our 2 differents two-digits
	for tens2 := '0'; tens2 <= '9'; tens2++ {
		for units2 := '0'; units2 <= '9'; units2++ {
			for tens1 := '0'; tens1 <= '9'; tens1++ {
				for units1 := '0'; units1 <= '9'; units1++ {
					// Testing if they're in ascending order
					if units2 < units1 && tens2 <= tens1 {
						z01.PrintRune(tens2)
						z01.PrintRune(units2)
						z01.PrintRune(' ')
						z01.PrintRune(tens1)
						z01.PrintRune(units1)
						if units1 == '9' && tens1 == '9' && units2 == '8' && tens2 == '9' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}