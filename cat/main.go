package main

import (
	"fmt"
	"os"
)

func main() {
	if !(len(os.Args[1:]) < 2) {
		file1, _ := os.Open(os.Args[1])
		fi1, _ := file1.Stat()
		file2, _ := os.Open(os.Args[2])
		fi2, _ := file2.Stat()
		arr1 := make([]byte, fi1.Size())
		arr2 := make([]byte, fi2.Size())
		file1.Read(arr1)
		file2.Read(arr2)
		fmt.Print(string(arr1))
		fmt.Print(string(arr2))
	} else {
		fmt.Println("File name missing")
	}
}
