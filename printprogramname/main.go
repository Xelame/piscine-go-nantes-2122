package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args[0]
	for _, runeValue := range fichier {
		if rune(runeValue) != '/' && rune(runeValue) != '.' {
			z01.PrintRune(rune(runeValue))
		}
	}
}
