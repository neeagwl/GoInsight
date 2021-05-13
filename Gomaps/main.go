package main

import "fmt"

func main(){

	colors := map[string]string{
		"red" : "#ff000",
		"green" : "#f3470f",
	}

	colors["white"] ="#fffff"

	delete(colors, "green")

	printMap(colors)
	
}


func printMap(c map[string]string) {

	for color, hex := range c{
		fmt.Println("Hex color for ", color, " is ", hex )
	}
	
}