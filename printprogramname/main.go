package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args[0]
	begin := 0
	isFirst := true
	for i := len(fichier) - 1; i >= 0; i-- {
		if fichier[i] == '/' && isFirst {
			begin = i
			isFirst = !isFirst
		}
	}
	for i, value := range fichier {
		if i > begin {
			z01.PrintRune(value)
		}
	}
	z01.PrintRune('\n')
}
