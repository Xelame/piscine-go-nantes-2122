package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args
	revfichier := []string{}
	for i := len(fichier) - 1; i > 0; i-- {
		revfichier = append(revfichier, fichier[i])
	}
	for _, arg := range revfichier {
		begin := 0
		isFirst := true
		for i := len(arg) - 1; i >= 0; i-- {
			if arg[i] == '/' && isFirst {
				begin = i - 1
				isFirst = !isFirst
			}
		}
		for i, value := range arg {
			if i >= begin {
				z01.PrintRune(value)
			}
		}
		z01.PrintRune('\n')
	}
}
