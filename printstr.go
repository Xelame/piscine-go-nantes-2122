package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(aString string) {
	for _, letter := range aString {
		z01.PrintRune(letter)
	}
}
