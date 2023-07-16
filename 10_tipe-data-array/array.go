package main

import "fmt"

func main() {
	var animals [2]string
	animals[0] = "Jerapah"
	animals[1] = "Gajah"

	fmt.Println(animals)

	var fruits = [3]string{
		"Melon",
		"Pisang",
		"Apel",
	}
	fmt.Println(fruits)

	fmt.Printf("Length of animals %d", len(animals))
}
