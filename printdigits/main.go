package main

import "github.com/01-edu/z01"

func main() {
	// Write numbers to 0 at 9 on a same line with a z01's function
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}

	// Make a new line
	z01.PrintRune('\n')
}
