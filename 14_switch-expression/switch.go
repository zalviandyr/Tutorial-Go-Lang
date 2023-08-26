package main

import "fmt"

func main() {
	name := "Alvixn"

	switch name {
	case "Alvin":
		fmt.Println("Hai Alvin")
	default:
		fmt.Println("Orang asing")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama lebih dari 5")
	case false:
		fmt.Println("Nama kurang dari 5")
	}

	length2 := len(name)
	switch {
	case length2 > 2:
		fmt.Println("Nama lebih dari 2")
	}
}
