package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args[1:]
	for _, arg := range fichier {
		for _, value := range arg {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
