package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("no argument")
		os.Exit(1)
	}
	fname := os.Args[1]
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("Error with file", fname, "Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
