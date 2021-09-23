package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args[0]
	begin := 0
	for i := len(fichier) - 1; i >= 0; i-- {
		if fichier[i] == '/' {
			begin = i + 1
		}
	}
	for begin < len(fichier) {
		z01.PrintRune(rune(fichier[begin]))
		begin++
	}
}
