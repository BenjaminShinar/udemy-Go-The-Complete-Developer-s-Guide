package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (n int, err error) {
	//return 1, nil // legal garbage

	fmt.Println(string(bs))
	fmt.Println("wrote this many bytes:", len(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error getting", err)
		os.Exit(1)
	}
	//fmt.Println(resp)
	//fmt.Printf("%+v\n", resp)

	// bs := make([]byte, 99999) // create a slice with this size
	// n, err := resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// if err != nil {
	// 	fmt.Println("Error reading", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(n, "elements read!")
	lw := logWriter{}
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}
