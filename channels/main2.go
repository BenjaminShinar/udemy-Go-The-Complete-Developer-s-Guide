package main

import "fmt"

func wait(c chan string) {
	fmt.Println(<-c, "received")
	c <- "done"
}
func main() {
	c := make(chan string)
	go wait(c)
	c <- "Hi there!"
	fmt.Println("channel waiting!")
	fmt.Println(<-c)
}
