package main

import ("github.com/01-edu/z01"
		"piscine"
)

func main() {
}

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}