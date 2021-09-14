package main

import "github.com/01-edu/z01"

func main() {
	// Write the reversed ABC on a same line with a z01's function
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}

	// Make a new line
	z01.PrintRune('\n')
}
