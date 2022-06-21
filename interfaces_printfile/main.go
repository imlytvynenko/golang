package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("Err:", "pls pass the file name in args")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	io.Copy(os.Stdout, file)
}
