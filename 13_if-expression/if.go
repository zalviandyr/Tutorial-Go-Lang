package main

import "fmt"

func main() {
	name := "alvin"

	if name == "alvin" {
		fmt.Printf("Hai %s", name)
	} else {
		fmt.Println("Orang asing")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama Lebih dari 5")
	}
}
