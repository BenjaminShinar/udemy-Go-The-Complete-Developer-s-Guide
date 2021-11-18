package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Holla!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	var eb englishBot
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
