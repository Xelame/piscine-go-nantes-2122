package main

import (
	"fmt"
	"os"
)

func main() {
	if !(len(os.Args[1:]) < 1) {
		if !(len(os.Args[1:]) > 1) {
			file, _ := os.Open(os.Args[1])
			fi, _ := file.Stat()
			arr := make([]byte, fi.Size())
			file.Read(arr)
			fmt.Print(string(arr))
		} else {
			fmt.Println("Too many arguments")
		}
	} else {
		fmt.Println("File name missing")
	}
}
