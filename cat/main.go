package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) <= 2 {
		file1, _ := os.Open(os.Args[1])
		fi1, _ := file1.Stat()
		arr1 := make([]byte, fi1.Size())
		file1.Read(arr1)
		for _, value := range arr1 {
			z01.PrintRune(rune(value))
		}
	}
	if len(os.Args[1:]) == 2 {
		file2, _ := os.Open(os.Args[2])
		fi2, _ := file2.Stat()
		arr2 := make([]byte, fi2.Size())
		file2.Read(arr2)
		for _, value := range arr2 {
			z01.PrintRune(rune(value))
		}
	}
}
