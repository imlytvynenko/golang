package main

import "fmt"

func main() {
	// First way to declare map
	colors := map[string]string{
		"red": "#FF0000",
		"green": "#00FF00",
	}

	fmt.Println(colors)

	delete(colors, "green")

	fmt.Println(colors)

	// Second way to declare map
	darkColors := make(map[string]string)
	darkColors["black"] = "#000000"

	fmt.Println(darkColors)

	// Third way to declare map
	var rainbowColors map[string]string

	fmt.Println(rainbowColors)
}
