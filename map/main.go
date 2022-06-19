package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#FF0000",
		"green": "#00FF00",
		"white": "#000000",
	}

	printColors(colors)
}

func printColors(c map[string]string) {
	for color, hex := range c{
		fmt.Println("Hex color for", color, "is", hex)
	}
}
