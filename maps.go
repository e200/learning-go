package main

import "fmt"

func main() {
	colors := make(map[string]string)

	colors["white"] = "#fff"
	colors["red"] = "#f00"
	colors["black"] = "#000"
	colors["blue"] = "#00f"

	printMap(colors)

	delete(colors, "white")

	printMap(colors)
}

func printMap(colors map[string]string) {
	for key, color := range colors {
		fmt.Println(key + ": " + color)
	}
}
