package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, arg := range a {
		for _, value := range arg {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
