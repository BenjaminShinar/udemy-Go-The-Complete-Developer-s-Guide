package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	// for l := range c {
	// 	go checkLink(l, c)
	// }
	for l := range c {
		//fmt.Println("l is ", l)
		go func(li string) {
			time.Sleep(5 * time.Second)
			checkLink(li, c)
		}(l) //invoke
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is responding")
	c <- link
}
