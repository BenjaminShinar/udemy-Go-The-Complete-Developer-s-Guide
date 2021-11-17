package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)
	map2 := make(map[int]string)
	map2[1] = "one"
	map2[0] = "zero"
	fmt.Println(map2)
	delete(map2, 0)
	fmt.Println(map2)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("color", color, "hex code", hex)
	}
}
