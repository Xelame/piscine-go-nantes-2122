package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args
	for i, strValue := range fichier {
		if i == len(fichier)-1 {
			for _, runeValue := range strValue {
				if rune(runeValue) != '/' {
					z01.PrintRune(rune(runeValue))
				}
			}
		}
	}
}
