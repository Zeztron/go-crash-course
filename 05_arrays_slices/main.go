package main

import "fmt"

func main() {
	// Arrays
	var fruitArray [2]string
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	// OR
	fruitArr := [2]string{"Banana", "Strawberry"}

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArray)
	fmt.Println(fruitArr)
	fmt.Println(fruitSlice)
}
