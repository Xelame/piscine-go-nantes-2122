package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args[1:]
	revfichier := []string{}
	for i := len(fichier) - 1; i >= 0; i-- {
		revfichier = append(revfichier, fichier[i])
	}
	for _, arg := range revfichier {
		for _, value := range arg {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
