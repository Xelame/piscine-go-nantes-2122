package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	fichier := os.Args[1:]
	arg := []rune(Join(fichier, " "))
	voyela, voyelb := 0, 0
	if len(arg) > 0 {
		for indexb, indexe := 0, len(arg)-1; indexb <= indexe; indexb, indexe = indexb+1, indexe-1 {
			if arg[indexe] == 'a' || arg[indexe] == 'e' || arg[indexe] == 'o' || arg[indexe] == 'i' || arg[indexe] == 'A' || arg[indexe] == 'E' || arg[indexe] == 'O' || arg[indexe] == 'I' {
				voyela = indexe
			}
			if arg[indexb] == 'a' || arg[indexb] == 'e' || arg[indexb] == 'o' || arg[indexb] == 'i' || arg[indexb] == 'A' || arg[indexb] == 'E' || arg[indexb] == 'O' || arg[indexb] == 'I' {
				voyelb = indexb
			}
			if voyela != 0 && voyelb != 0 {
				Swap(&arg[voyela], &arg[voyelb])
				voyela, voyelb = 0, 0
			}
		}
		for index := range arg {
			z01.PrintRune(arg[index])
		}
	}
	z01.PrintRune('\n')
}

func Swap(a *rune, b *rune) {
	*a, *b = *b, *a
}

func Concat(s1 string, s2 string) string {
	return s1 + s2
}

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	} else {
		return Concat(strs[0], Concat(sep, Join(strs[1:], sep)))
	}
}
