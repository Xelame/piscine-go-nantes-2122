package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) == 1 {
		if !(len(os.Args[1:]) > 1) {
			file, err := os.Open(os.Args[1])
			if err != nil {
				log.Fatal(err)
			}
			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}
			arr := make([]byte, fi.Size())
			file.Read(arr)
			fmt.Println(string(arr))
		} else {
			fmt.Println("Too many arguments")
		}
	} else {
		fmt.Println("File name missing")
	}
}
