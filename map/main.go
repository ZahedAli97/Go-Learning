package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	colors["white"] = "#ffffff"

	//delete(colors, "red")
	printMap(colors)

	//fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, ":", hex)
	}
}
