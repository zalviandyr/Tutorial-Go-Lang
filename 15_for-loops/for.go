package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Printf("Counter: %d\n", counter)

		counter++
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("I ke: %d\n", i)
	}

	fruits := []string{"Apel", "Anggur", "Kurma"}
	for i, fruit := range fruits {
		fmt.Printf("Index ke %d terdapat buah %s\n", i, fruit)
	}
}
