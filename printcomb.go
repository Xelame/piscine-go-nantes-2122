package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for units := '0'; units <= '9'; units++ {
		for tens := '0'; tens <= '9'; tens++ {
			for hundreds := '0'; hundreds <= '9'; hundreds++ {
				if units < tens && tens < hundreds {
					z01.PrintRune(units)
					z01.PrintRune(tens)
					z01.PrintRune(hundreds)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}